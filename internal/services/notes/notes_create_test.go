package notes

import (
	"context"
	"errors"
	"testing"
)

func TestNotesService_CreateNote_Success(t *testing.T) {
	var gotAuthorId int
	var gotText string

	repo := &mockNotesRepo{
		createNoteFn: func(ctx context.Context, authorId int, text string) (int, error) {
			gotAuthorId = authorId
			gotText = text
			return 7, nil
		},
	}
	svc := NewNotesService(repo)

	id, err := svc.CreateNote(context.Background(), 42, "hello")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if id != 7 {
		t.Errorf("id = %d, want 7", id)
	}
	if gotAuthorId != 42 {
		t.Errorf("repo received authorId = %d, want 42", gotAuthorId)
	}
	if gotText != "hello" {
		t.Errorf("repo received text = %q, want %q", gotText, "hello")
	}
}

func TestNotesService_CreateNote_RepoError(t *testing.T) {
	wantErr := errors.New("db is down")
	repo := &mockNotesRepo{
		createNoteFn: func(ctx context.Context, authorId int, text string) (int, error) {
			return 0, wantErr
		},
	}
	svc := NewNotesService(repo)

	id, err := svc.CreateNote(context.Background(), 1, "hello")
	if !errors.Is(err, wantErr) {
		t.Errorf("err = %v, want %v", err, wantErr)
	}
	if id != 0 {
		t.Errorf("id = %d, want 0 on error", id)
	}
}
