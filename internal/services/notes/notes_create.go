package notes

import "context"

func (ns *NotesService) CreateNote(
	ctx context.Context,
	authorId int,
	text string,
) (int, error) {
	return ns.notesRepo.CreateNote(ctx, authorId, text)
}
