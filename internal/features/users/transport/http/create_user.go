package users_transport_http

import (
	"net/http"

	core_logger "github.com/darealeuslsmachine/go-todoapp/internal/core/logger"
	core_http_request "github.com/darealeuslsmachine/go-todoapp/internal/core/transport/http/request"
	core_http_response "github.com/darealeuslsmachine/go-todoapp/internal/core/transport/http/response"
)

type CreateUserRequest struct {
	FullName    string  `json:"full_name" validate:"required,min=3,max=100"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,required,min=10,max=15,startswith=+"`
}

type CreateUserResponse struct {
	ID          string  `json:"id"`
	Version     int     `json:"version"`
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

func (h *UserHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	log.Debug("invoke CreateUser handler")

	var requestDTO CreateUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &requestDTO); err != nil {
		responseHandler.ErrorResponse(err, "create user. failed to decode and validate HTTP request")
		return
	}

	rw.WriteHeader(http.StatusOK)
}
