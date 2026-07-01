package notes

import (
	"context"

	"yat-blog-notes/internal/repositories"
)

type mockNotesRepo struct {
	createNoteFn         func(ctx context.Context, authorId int, text string) (int, error)
	countNotesByAuthorFn func(ctx context.Context, authorId int) (int, error)
	getNoteFn            func(ctx context.Context, noteId int) (*repositories.Note, error)
	getNotesByAuthorFn   func(ctx context.Context, authorId, limit, offset int) (*repositories.GetNotesByAuthorIdDto, error)
}

func (m *mockNotesRepo) CreateNote(
	ctx context.Context,
	authorId int,
	text string,
) (int, error) {
	return m.createNoteFn(ctx, authorId, text)
}

func (m *mockNotesRepo) CountNotesByAuthorId(
	ctx context.Context,
	authorId int,
) (int, error) {
	return m.countNotesByAuthorFn(ctx, authorId)
}

func (m *mockNotesRepo) GetNote(
	ctx context.Context,
	noteId int,
) (*repositories.Note, error) {
	return m.getNoteFn(ctx, noteId)
}

func (m *mockNotesRepo) GetNotesByAuthorId(
	ctx context.Context,
	authorId, limit, offset int,
) (*repositories.GetNotesByAuthorIdDto, error) {
	return m.getNotesByAuthorFn(ctx, authorId, limit, offset)
}
