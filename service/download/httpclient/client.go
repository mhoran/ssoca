package httpclient

import (
	"errors"
	"fmt"
	"io"
	"net/url"

	"github.com/dpb587/ssoca/httpclient"
	"github.com/dpb587/ssoca/service/download/api"

	boshcrypto "github.com/cloudfoundry/bosh-utils/crypto"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

func New(baseclient httpclient.Client, service string) (Client, error) {
	if baseclient == nil {
		return nil, errors.New("client is nil")
	}

	return &client{
		client:  baseclient,
		service: service,
	}, nil
}

type client struct {
	client  httpclient.Client
	service string
}

func (c client) GetList() (api.ListResponse, error) {
	out := api.ListResponse{}
	path := fmt.Sprintf("/%s/list", c.service)

	err := c.client.APIGet(path, &out)
	if err != nil {
		return out, bosherr.WrapErrorf(err, "Getting %s", path)
	}

	return out, nil
}

func (c client) Download(name string, target io.ReadWriteSeeker) error {
	list, err := c.GetList()
	if err != nil {
		return bosherr.WrapError(err, "Listing artifacts")
	}

	for _, file := range list.Files {
		if file.Name != name {
			continue
		}

		return c.download(file, target)
	}

	return fmt.Errorf("File is not known: %s", name)
}

func (c client) download(file api.ListFileResponse, target io.ReadWriteSeeker) error {
	path := fmt.Sprintf("/%s/get?name=%s", c.service, url.QueryEscape(file.Name))

	res, err := c.client.Get(path)
	if err != nil {
		return bosherr.WrapError(err, "Getting file")
	}

	_, err = io.Copy(target, res.Body)
	if err != nil {
		return bosherr.WrapError(err, "Streaming to file")
	}

	var algo boshcrypto.Algorithm
	var hash string

	if file.Digest.SHA512 != "" {
		algo = boshcrypto.DigestAlgorithmSHA512
		hash = fmt.Sprintf("sha512:%s", file.Digest.SHA512)
	} else if file.Digest.SHA256 != "" {
		algo = boshcrypto.DigestAlgorithmSHA256
		hash = fmt.Sprintf("sha256:%s", file.Digest.SHA256)
	} else if file.Digest.SHA1 != "" {
		algo = boshcrypto.DigestAlgorithmSHA1
		hash = file.Digest.SHA1
	} else {
		return errors.New("No digest available to verify download")
	}

	digest, err := algo.CreateDigest(target)
	if err != nil {
		return bosherr.WrapError(err, "Creating digest")
	}

	if digest.String() != hash {
		return fmt.Errorf("Expected digest '%s' but got '%s'", hash, digest.String())
	}

	return nil
}