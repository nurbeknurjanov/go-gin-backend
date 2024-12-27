package grpc

import (
	"fmt"
	api "github.com/nurbeknurjanov/go-grpc/api"
	"google.golang.org/grpc"
	"net"
)

type Deps struct {
	UserHandler api.AuthServer
}
type Server struct {
	srv *grpc.Server
	Deps
}

func NewServer(deps Deps) *Server {
	return &Server{
		srv:  grpc.NewServer(),
		Deps: deps,
	}
}

func (s *Server) ListenAndServer(port int) error {
	addr := fmt.Sprintf(":%d", port)
	listenTCP, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	api.RegisterAuthServer(s.srv, s.UserHandler)

	if err := s.srv.Serve(listenTCP); err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop() {
	s.srv.GracefulStop()
}
