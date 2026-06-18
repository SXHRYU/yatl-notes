package http_server

import (
	"context"

	"yat-blog-notes/internal/configs"
)

type notesService interface {
	GetAuthorNotes(ctx context.Context, authorId, page, limit int) ([]string, error)
}

type NotesController struct {
	notesSrv notesService
	config   *configs.Config
}

func NewNotesController(notesSrv notesService, config *configs.Config) *NotesController {
	return &NotesController{notesSrv: notesSrv, config: config}
}
