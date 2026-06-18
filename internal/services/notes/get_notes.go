package notes

import "context"

func (nc *NotesService) GetAuthorNotes(
	ctx context.Context,
	authorId, page, limit int,
) ([]string, error) {
	return nc.notesRepo.GetNotesByAuthorId(ctx, authorId, limit, (page-1)*limit)
}
