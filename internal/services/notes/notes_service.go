package notes

import "context"

type notesRepository interface {
	GetNotesByAuthorId(ctx context.Context, authorId, limit, offset int) ([]string, error)
}

type NotesService struct {
	notesRepo notesRepository
}

func NewNotesService(notesRepo notesRepository) *NotesService {
	return &NotesService{notesRepo: notesRepo}
}
