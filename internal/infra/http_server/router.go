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
	api.HandleFunc("POST /create/", isAuthenticated(controller.CreateNote)).With(
		option.Summary("Create note"),
		option.Request(new(CreateNoteRequest)),
		option.Response(http.StatusCreated, new(handlers.CreateNoteResponse)),
		option.Response(http.StatusBadRequest, new(ErrorResponse)),
		option.Response(http.StatusRequestEntityTooLarge, new(ErrorResponse)),
	)
	api.HandleFunc("GET /", controller.GetAuthorNotes).With(
		option.Summary("Get author's notes"),
		option.Request(new(GetAuthorNotesRequest)),
		option.Response(http.StatusOK, new(GetAuthorNotesResponse)),
		option.Response(http.StatusBadRequest, new(ErrorResponse)),
	)
	api.HandleFunc("GET /{id}", controller.GetNote).With(
		option.Summary("Get note"),
		option.Request(new(GetNoteRequest)),
		option.Response(http.StatusOK, new(handlers.Note)),
		option.Response(http.StatusNotFound, new(ErrorResponse)),
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

// TODO: добавить слаг
type GetNoteRequest struct {
	Id int `path:"id" required:"true"`
}

type CreateNoteRequest struct {
	// TODO: убрать как появится JWT
	UserId int    `header:"User-Id" required:"true"`
	Text   string `required:"true" json:"text" minLength:"1" maxLength:"10<<10"`
}

type ErrorResponse struct {
	Error string
}
