package http_server

import (
	"net/http"
	"time"

	"yat-blog-notes/internal/configs"
)

type Server struct {
	srv    *http.Server
	config *configs.Config
}

func NewServer(addr string, router http.Handler, config *configs.Config) *Server {
	srv := &http.Server{
		Addr:              addr,
		Handler:           router,
		ReadHeaderTimeout: time.Duration(config.Http.ReadHeaderTimeout),
	}
	return &Server{
		srv:    srv,
		config: config,
	}
}

func (s *Server) ListenAndServe() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Close() error {
	return s.srv.Close()
}
