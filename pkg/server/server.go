package server

import (
	"context"
	"net/http"
)

type Server struct {
	http.Server
}

func New(addr string, handler http.Handler) *Server {
	s := &Server{}
	s.Addr = addr
	s.Handler = handler
	return s
}

func (s *Server) Start() error {
	return s.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.Shutdown(ctx)
}
