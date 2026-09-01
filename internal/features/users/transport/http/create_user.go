package users_transport_http

import (
	"encoding/json"
	"fmt"
	"net/http"

	core_logger "github.com/darealeuslsmachine/go-todoapp/internal/core/logger"
)

type CreateUserRequest struct {
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
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

	log.Debug("invoke CreateUser handler")

	var request CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		fmt.Println("error:", err)
	}

	rw.WriteHeader(http.StatusOK)
}
