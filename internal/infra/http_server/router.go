package http_server

import (
	"net/http"

	handlers "yat-blog-notes/internal/infra/http_server/handlers"
)

type Router struct {
	mux *http.ServeMux
}

func NewRouter(controller *handlers.NotesController) *Router {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /create/", isAuthenticated(controller.CreateNote))

	return &Router{
		mux: mux,
	}
}

func (r *Router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(rw, req)
}
