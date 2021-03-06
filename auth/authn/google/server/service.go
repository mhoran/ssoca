package server

import (
	"net/http"

	"github.com/dpb587/ssoca/auth"
	svc "github.com/dpb587/ssoca/auth/authn/google"
	svcconfig "github.com/dpb587/ssoca/auth/authn/google/config"
	oauth2support "github.com/dpb587/ssoca/auth/authn/support/oauth2"
	"github.com/dpb587/ssoca/server/service/req"
)

type Service struct {
	svc.Service

	name   string
	config svcconfig.Config
	oauth  oauth2support.Backend
}

func NewService(name string, config svcconfig.Config, oauth oauth2support.Backend) Service {
	return Service{
		name:   name,
		config: config,
		oauth:  oauth,
	}
}

func (s Service) Name() string {
	return s.name
}

func (s Service) Metadata() interface{} {
	return nil
}

func (s Service) GetRoutes() []req.RouteHandler {
	return s.oauth.GetRoutes(s.OAuthUserProfileLoader)
}

func (s Service) VerifyAuthorization(_ http.Request, _ *auth.Token) error {
	return nil
}
