package users_transport_http

import (
	"context"
	"net/http"

	"github.com/darealeuslsmachine/go-todoapp/internal/core/domain"
	core_http_server "github.com/darealeuslsmachine/go-todoapp/internal/core/transport/http/server"
)

type UserHTTPHandler struct {
	UsersService UsersService
}

type UsersService interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)
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
