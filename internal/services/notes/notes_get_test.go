package notes

import (
	"context"
	"errors"
	"testing"
	"time"

	"yat-blog-notes/internal/repositories"
)

func TestNotesService_GetNote_Success(t *testing.T) {
	createdAt := time.Date(2026, 9, 1, 0, 0, 0, 0, time.UTC)
	var gotNoteId int

	repo := &mockNotesRepo{
		getNoteFn: func(ctx context.Context, noteId int) (*repositories.Note, error) {
			gotNoteId = noteId
			return &repositories.Note{
				Id:        noteId,
				Title:     "Title",
				Text:      "Text",
				CreatedAt: createdAt,
			}, nil
		},
	}
	svc := NewNotesService(repo)

	note, err := svc.GetNote(context.Background(), 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotNoteId != 5 {
		t.Errorf("repo received noteId = %d, want 5", gotNoteId)
	}
	if note.Id != 5 || note.Title != "Title" || note.Text != "Text" {
		t.Errorf("unexpected note: %+v", note)
	}
	if !note.CreatedAt.Equal(createdAt) {
		t.Errorf("CreatedAt = %v, want %v", note.CreatedAt, createdAt)
	}
}

func TestNotesService_GetNote_NotFound(t *testing.T) {
	wantErr := errors.New("sql: no rows in result set")
	repo := &mockNotesRepo{
		getNoteFn: func(ctx context.Context, noteId int) (*repositories.Note, error) {
			return nil, wantErr
		},
	}
	svc := NewNotesService(repo)

	note, err := svc.GetNote(context.Background(), 999)
	if !errors.Is(err, wantErr) {
		t.Errorf("err = %v, want %v", err, wantErr)
	}
	if note != nil {
		t.Errorf("note = %+v, want nil on error", note)
	}
}
