package http_server

import "net/http"

// type notesService interface {}

type NotesController struct {
	// service notesService
}

func NewNotesController() *NotesController {
	return &NotesController{}
}

func (ac *NotesController) CreateNote(rw http.ResponseWriter, req *http.Request) {}
