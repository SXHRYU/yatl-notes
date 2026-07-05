package notes

import (
	"context"

	notes "yat-blog-notes/internal/services"
)

func (nc *NotesService) GetNote(ctx context.Context, noteId int) (*notes.Note, error) {
	response, err := nc.notesRepo.GetNote(ctx, noteId)
	if err != nil {
		return nil, err // todo: service errors
	}
	return notes.ToServiceNote(response), nil
}

func (nc *NotesService) GetNoteBySlug(
	ctx context.Context,
	slug string,
) (*notes.Note, error) {
	response, err := nc.notesRepo.GetNoteBySlug(ctx, slug)
	if err != nil {
		return nil, err // todo: service errors
	}
	return notes.ToServiceNote(response), nil
}
