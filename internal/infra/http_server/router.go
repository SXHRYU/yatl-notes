package http_server

import (
	"net/http"

	"github.com/oaswrap/spec/adapter/httpopenapi"
	"github.com/oaswrap/spec/option"
	handlers "yat-blog-notes/internal/infra/http_server/handlers"
)

type Router struct {
	mux *http.ServeMux
}

func NewRouter(controller *handlers.NotesController) *Router {
	mux := http.NewServeMux()
	r := httpopenapi.NewGenerator(
		mux,
		option.WithTitle("Notes API"),
		option.WithVersion("1.0.0"),
	)

	api := r.Group("/api")
	api.HandleFunc("POST /create/", isAuthenticated(controller.CreateNote))
	api.HandleFunc("GET /", controller.GetAuthorNotes).With(
		option.Summary("Get author's notes"),
		option.Request(new(GetAuthorNotesRequest)),
		option.Response(http.StatusOK, new(GetAuthorNotesResponse)),
		option.Response(http.StatusBadRequest, new(ErrorResponse)),
	)

	return &Router{
		mux: mux,
	}
}

func (r *Router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(rw, req)
}

type PageLimitPaginated struct {
	Page  int `query:"page"  default:"1"`
	Limit int `query:"limit" default:"20"`
}

type GetAuthorNotesRequest struct {
	PageLimitPaginated
	AuthorId string `query:"author_id" required:"true"`
}

type GetAuthorNotesResponse struct {
	Notes handlers.GetAuthorNotesResponse
}

type ErrorResponse struct {
	Error string
}
