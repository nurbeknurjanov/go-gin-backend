package go_backend

import (
	"context"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Start(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Handler: handler,
		Addr:    ":" + port,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
