package notes

import (
	"context"

	"yat-blog-notes/internal/repositories"
)

type notesRepository interface {
	CreateNote(ctx context.Context, authorId int, text string) (int, error)
	GetNote(ctx context.Context, noteId int) (*repositories.Note, error)
	GetNotesByAuthorId(
		ctx context.Context,
		authorId, limit, offset int,
	) (*repositories.GetNotesByAuthorIdDto, error)
}

type NotesService struct {
	notesRepo notesRepository
}

func NewNotesService(notesRepo notesRepository) *NotesService {
	return &NotesService{notesRepo: notesRepo}
}
