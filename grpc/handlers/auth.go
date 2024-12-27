package handlers

import (
	"fmt"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/services"
	"github.com/nurbeknurjanov/go-grpc/api"
)

type AuthHandler struct {
	api.UnimplementedAuthServer

	services *services.Services
}

func NewAuthHandler(services *services.Services) *AuthHandler {
	return &AuthHandler{services: services}
}

func (h *AuthHandler) Login(req *api.LoginRequest) (*api.LoginResponse, error) {

	fmt.Println("api login request", req.Email, req.Password)

	res := &api.LoginResponse{
		Token: "new generated token",
	}

	//h.service.Login()
	return res, nil
}
