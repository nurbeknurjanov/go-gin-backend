package grpc_handlers

import (
	"fmt"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/services"
	"github.com/nurbeknurjanov/go-grpc/api"
)

type AuthHandler struct {
	api.UnimplementedAuthServer

	authService services.Auth
}

func NewAuthHandler(authService services.Auth) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(req *api.LoginRequest) (*api.LoginResponse, error) {

	fmt.Println("api login request", req.Email, req.Password)

	res := &api.LoginResponse{
		Token: "new generated token",
	}

	//h.authService.Login()
	return res, nil
}
