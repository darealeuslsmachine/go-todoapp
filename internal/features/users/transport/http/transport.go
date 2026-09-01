package users_transport_http

import (
	"net/http"

	core_http_server "github.com/darealeuslsmachine/go-todoapp/internal/core/transport/http/server"
)

type UserHTTPHandler struct {
	UsersService UsersService
}

type UsersService interface {
}

func NewUserHTTPHandler(
	usersService UsersService,
) *UserHTTPHandler {
	return &UserHTTPHandler{
		UsersService: usersService,
	}
}

func (h *UserHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: h.CreateUser,
		},
	}
}
