package http_server

import "net/http"

type Server struct {
	srv *http.Server
}

func NewServer(addr string, router http.Handler) *Server {
	srv := &http.Server{Addr: addr, Handler: router}
	return &Server{srv}
}

func (s *Server) ListenAndServe() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Close() error {
	return s.srv.Close()
}
