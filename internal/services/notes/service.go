package notes

import (
	"context"

	"yat-blog-notes/internal/repositories"
)

type notesRepository interface {
	CreateNote(ctx context.Context, authorId int, text string) (int, error)
	CountNotesByAuthorId(ctx context.Context, authorId int) (int, error)
	GetNote(ctx context.Context, noteId int) (*repositories.Note, error)
	GetNoteBySlug(ctx context.Context, slug string) (*repositories.Note, error)
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
