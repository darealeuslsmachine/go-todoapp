package users_transport_http

import (
	"net/http"

	"github.com/darealeuslsmachine/go-todoapp/internal/core/domain"
	core_logger "github.com/darealeuslsmachine/go-todoapp/internal/core/logger"
	core_http_request "github.com/darealeuslsmachine/go-todoapp/internal/core/transport/http/request"
	core_http_response "github.com/darealeuslsmachine/go-todoapp/internal/core/transport/http/response"
)

type CreateUserRequest struct {
	FullName    string  `json:"full_name" validate:"required,min=3,max=100"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,required,min=10,max=15,startswith=+"`
}

type CreateUserResponse struct {
	ID          int     `json:"id"`
	Version     int     `json:"version"`
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

func (h *UserHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var requestDTO CreateUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &requestDTO); err != nil {
		responseHandler.ErrorResponse(err, "creating user. failed to decode and validate HTTP request")
		return
	}

	userDomain := domainFromDTO(requestDTO)

	userDomain, err := h.UsersService.CreateUser(ctx, userDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "creating user. failed to create user")

		return
	}

	response := dtoFromDomain(userDomain)

	responseHandler.JSONResponse(response, http.StatusCreated)
}

func dtoFromDomain(user domain.User) CreateUserResponse {
	return CreateUserResponse{
		ID:          user.ID,
		Version:     user.Version,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
	}
}

func domainFromDTO(dto CreateUserRequest) domain.User {
	return domain.NewUserUninitialized(
		dto.FullName,
		dto.PhoneNumber,
	)
}
