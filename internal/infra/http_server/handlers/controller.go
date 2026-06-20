package http_server

import (
	"context"

	"yat-blog-notes/internal/configs"
	notes "yat-blog-notes/internal/services"
)

type notesService interface {
	GetAuthorNotes(
		ctx context.Context,
		authorId, page, limit int,
	) (*notes.GetNotesByAuthorIdDto, error)
}

type NotesController struct {
	notesSrv notesService
	config   *configs.Config
}

func NewNotesController(notesSrv notesService, config *configs.Config) *NotesController {
	return &NotesController{notesSrv: notesSrv, config: config}
}
