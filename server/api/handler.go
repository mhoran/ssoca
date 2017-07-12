package api

import (
	"errors"
	"net/http"

	"github.com/Sirupsen/logrus"
	apierr "github.com/dpb587/ssoca/server/api/errors"
	"github.com/dpb587/ssoca/server/service"
	"github.com/dpb587/ssoca/server/service/req"

	uuid "github.com/nu7hatch/gouuid"
)

type apiHandler struct {
	authService service.AuthService
	apiService  service.Service
	handler     req.RouteHandler
	logger      logrus.FieldLogger
}

func CreateHandler(authService service.AuthService, apiService service.Service, handler req.RouteHandler, logger logrus.FieldLogger) (http.Handler, error) {
	return apiHandler{
		authService: authService,
		apiService:  apiService,
		handler:     handler,
		logger:      logger,
	}, nil
}

func (h apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	request := req.Request{
		RawRequest:  r,
		RawResponse: w,
	}

	requestUUID, err := uuid.NewV4()
	if err != nil {
		h.sendGenericErrorResponse(request, apierr.WrapError(err, "Generating request ID"))

		return
	}

	request.ID = requestUUID.String()
	request.LoggerContext = logrus.Fields{
		"server.request.id":          request.ID,
		"server.request.remote_addr": r.RemoteAddr,
		"service.name":               h.apiService.Name(),
		"service.type":               h.apiService.Type(),
	}

	token, err := h.authService.ParseRequestAuth(*r)
	if err != nil {
		h.sendGenericErrorResponse(request, apierr.WrapError(err, "Parsing authentication token"))

		return
	}

	request.AuthToken = token

	authz, err := h.apiService.IsAuthorized(*r, request.AuthToken)
	if err != nil {
		h.sendGenericErrorResponse(request, apierr.WrapError(err, "Checking service authorization"))

		return
	} else if !authz {
		h.sendErrorResponse(request, apierr.NewError(errors.New("Not authorized"), http.StatusForbidden, ""))

		return
	}

	if token != nil {
		request.LoggerContext["auth.user_id"] = token.ID
	}

	err = h.handler.Execute(request)
	if err != nil {
		h.sendGenericErrorResponse(request, apierr.WrapError(err, "Executing handler"))
	}
}

func (h apiHandler) sendGenericErrorResponse(request req.Request, err error) {
	h.sendErrorResponse(request, apierr.NewError(err, http.StatusInternalServerError, ""))
}

func (h apiHandler) sendErrorResponse(request req.Request, err apierr.Error) {
	request.RawResponse.WriteHeader(err.Status)

	var loggerFunc func(args ...interface{})
	logger := h.logger.WithFields(request.LoggerContext).WithFields(logrus.Fields{
		"server.request.method": request.RawRequest.Method,
		"server.request.path":   request.RawRequest.URL.Path,
	})

	if err.Status >= 500 {
		loggerFunc = logger.Error
	} else {
		loggerFunc = logger.Warn
	}

	loggerFunc(err.Error())

	request.WritePayload(map[string]interface{}{
		"error": map[string]interface{}{
			"status":  err.Status,
			"message": err.PublicError,
		},
	})
}