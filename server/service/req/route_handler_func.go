package req

import (
	"net/http"

	"github.com/dpb587/ssoca/auth"
	"github.com/dpb587/ssoca/auth/authz/filter"
)

type RouteHandlerFunc struct {
	Path        string
	Func        http.HandlerFunc
	Requirement filter.Requirement
}

var _ RouteHandler = RouteHandlerFunc{}

func (h RouteHandlerFunc) Route() string {
	return h.Path
}

func (h RouteHandlerFunc) Execute(request Request) error {
	h.Func(request.RawResponse, request.RawRequest)

	return nil
}

func (h RouteHandlerFunc) VerifyAuthorization(r *http.Request, token *auth.Token) error {
	if h.Requirement == nil {
		return nil
	}

	return h.Requirement.VerifyAuthorization(r, token)
}
