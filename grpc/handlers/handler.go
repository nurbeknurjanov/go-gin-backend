package grpc_handlers

import (
	"github.com/nurbeknurjanov/go-gin-backend/pkg/services"
)

type Deps struct {
	Auth services.Auth
}

type GrpcHandlers struct {
	AuthHandler *AuthHandler
}

func NewGrpcHandlers(deps Deps) *GrpcHandlers {
	return &GrpcHandlers{
		AuthHandler: NewAuthHandler(deps.Auth),
	}
}
