package client

import (
	"fmt"
	"net/http"

	boshuaa "github.com/cloudfoundry/bosh-cli/uaa"
	"github.com/pkg/errors"

	"github.com/dpb587/ssoca/auth/authn/uaa/api"
	"github.com/dpb587/ssoca/config"
	env_api "github.com/dpb587/ssoca/service/env/api"
)

func (s Service) AuthLogin(remoteService env_api.InfoServiceResponse) (interface{}, error) {
	metadata := api.Metadata{}

	err := config.RemarshalJSON(remoteService.Metadata, &metadata)
	if err != nil {
		return nil, errors.Wrap(err, "Parsing metadata")
	}

	client, err := s.uaaClientFactory.CreateClient(metadata.URL, metadata.ClientID, metadata.ClientSecret, metadata.CACertificate)
	if err != nil {
		return nil, errors.Wrap(err, "Creating UAA client")
	}

	prompts, err := client.Prompts()
	if err != nil {
		return nil, errors.Wrap(err, "Discovering UAA prompts")
	}

	ui := s.runtime.GetUI()
	var answers []boshuaa.PromptAnswer

	for _, prompt := range prompts {
		if len(metadata.Prompts) > 0 {
			var matchedPrompt bool

			for _, expectedKey := range metadata.Prompts {
				if prompt.Key == expectedKey {
					matchedPrompt = true

					break
				}
			}

			if !matchedPrompt {
				continue
			}
		}

		var askFunc func(string) (string, error)

		if prompt.IsPassword() {
			askFunc = ui.AskForPassword
		} else {
			askFunc = ui.AskForText
		}

		value, err1 := askFunc(prompt.Label)
		if err1 != nil {
			return nil, err1
		}

		if value != "" {
			answer := boshuaa.PromptAnswer{Key: prompt.Key, Value: value}
			answers = append(answers, answer)
		}
	}

	accessToken, err := client.OwnerPasswordCredentialsGrant(answers)
	if err != nil {
		return nil, errors.Wrap(err, "Fetching credentials grant")
	}

	auth := AuthConfig{
		URL:           metadata.URL,
		CACertificate: metadata.CACertificate,
		ClientID:      metadata.ClientID,
		ClientSecret:  metadata.ClientSecret,
		RefreshToken:  accessToken.RefreshToken().Value(),
	}

	return auth, nil
}

func (s Service) AuthLogout() error {
	return nil
}

func (s Service) AuthRequest(req *http.Request) error {
	env, err := s.runtime.GetEnvironment()
	if err != nil {
		return errors.Wrap(err, "Getting environment")
	}

	authConfig := AuthConfig{}
	err = env.Auth.UnmarshalOptions(&authConfig)
	if err != nil {
		return errors.Wrap(err, "Parsing authentication options")
	}

	client, err := s.uaaClientFactory.CreateClient(authConfig.URL, authConfig.ClientID, authConfig.ClientSecret, authConfig.CACertificate)
	if err != nil {
		return errors.Wrap(err, "Creating UAA client")
	}

	staleToken := client.NewStaleAccessToken(authConfig.RefreshToken)
	accessToken, err := staleToken.Refresh()
	if err != nil {
		return errors.Wrap(err, "Refreshing token")
	}

	req.Header.Add("Authorization", fmt.Sprintf("%s %s", accessToken.Type(), accessToken.Value()))

	return nil
}
