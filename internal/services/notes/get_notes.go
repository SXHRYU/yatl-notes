package notes

import (
	"context"

	notes "yat-blog-notes/internal/services"
)

func (nc *NotesService) GetAuthorNotes(
	ctx context.Context,
	authorId, page, limit int,
) (*notes.GetNotesByAuthorIdDto, error) {
	response, err := nc.notesRepo.GetNotesByAuthorId(ctx, authorId, limit, (page-1)*limit)
	if err != nil {
		return nil, err // todo: service errors
	}
	return notes.ToServiceAuthorNotes(response), nil
}
