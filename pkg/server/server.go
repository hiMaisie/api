package server

import (
	"errors"
	"log"
	"net"
	"net/http"
	"context"
)

var (
	errRouteExists = errors.New("route already exists in server")
	routes = map[string]http.HandlerFunc{
		"/api/v1/health": handleHealthCheck,
	}
)

type Server struct {
	mux *http.ServeMux
	log    *log.Logger
	srv *http.Server
}

func New(logger *log.Logger, addr string) *Server {
	s := &Server{
		mux: http.NewServeMux(),
		log: logger,
	}

	for route, handler := range routes {
		s.mux.HandleFunc(route, handler)
	}

	s.srv = &http.Server{
		Addr: addr,
		Handler: s.mux,

		ErrorLog: s.log,

		// TODO: Start OpenTelemetry instrumentation here.
		BaseContext: func(lis net.Listener) context.Context {
			return context.Background()
		},
	}

	return s
}

func (s *Server) Start() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	s.log.Println("Shutting down...")
	return s.srv.Shutdown(ctx)
}
