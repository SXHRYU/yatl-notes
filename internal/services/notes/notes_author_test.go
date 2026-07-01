package notes

import (
	"context"
	"errors"
	"strings"
	"sync"
	"testing"

	"yat-blog-notes/internal/repositories"
)

func TestNotesService_GetAuthorNotes_Success(t *testing.T) {
	var mu sync.Mutex
	var gotAuthorId, gotLimit, gotOffset, gotCountAuthorId int

	repo := &mockNotesRepo{
		getNotesByAuthorFn: func(ctx context.Context, authorId, limit, offset int) (*repositories.GetNotesByAuthorIdDto, error) {
			mu.Lock()
			defer mu.Unlock()
			gotAuthorId, gotLimit, gotOffset = authorId, limit, offset
			return &repositories.GetNotesByAuthorIdDto{
				Notes: []repositories.Note{
					{
						Id:   1,
						Text: strings.Repeat("a", 500),
					}, // long enough to be shortened
					{Id: 2, Text: "short text"},
				},
			}, nil
		},
		countNotesByAuthorFn: func(ctx context.Context, authorId int) (int, error) {
			mu.Lock()
			defer mu.Unlock()
			gotCountAuthorId = authorId
			return 45, nil
		},
	}
	svc := NewNotesService(repo)

	got, err := svc.GetAuthorNotes(context.Background(), 7, 3, 20)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if gotAuthorId != 7 || gotCountAuthorId != 7 {
		t.Errorf(
			"authorId passed to repo = (%d, %d), want 7 for both calls",
			gotAuthorId,
			gotCountAuthorId,
		)
	}
	if gotLimit != 20 {
		t.Errorf("limit passed to repo = %d, want 20", gotLimit)
	}
	if gotOffset != 40 {
		t.Errorf("offset passed to repo = %d, want 40 (page-1)*limit", gotOffset)
	}

	if got.Pagination.Total != 45 {
		t.Errorf("Pagination.Total = %d, want 45", got.Pagination.Total)
	}
	if got.Pagination.Page != 3 {
		t.Errorf("Pagination.Page = %d, want 3", got.Pagination.Page)
	}
	wantPages := (45 + 20 - 1) / 20 // pagination.TotalPages(45, 20) == 3
	if got.Pagination.Pages != wantPages {
		t.Errorf("Pagination.Pages = %d, want %d", got.Pagination.Pages, wantPages)
	}

	if len(got.Notes) != 2 {
		t.Fatalf("len(Notes) = %d, want 2", len(got.Notes))
	}
	if !strings.HasSuffix(got.Notes[0].Text, "...") {
		t.Errorf(
			"expected long note text to be shortened with an ellipsis, got %q",
			got.Notes[0].Text,
		)
	}
	if got.Notes[0].Text == strings.Repeat("a", 500) {
		t.Errorf("expected long note text to be shortened, but it was left unchanged")
	}
	if got.Notes[1].Text != "short text" {
		t.Errorf(
			"short note text = %q, want unchanged %q",
			got.Notes[1].Text,
			"short text",
		)
	}
}

func TestNotesService_GetAuthorNotes_NotesRepoError(t *testing.T) {
	wantErr := errors.New("notes query failed")
	repo := &mockNotesRepo{
		getNotesByAuthorFn: func(ctx context.Context, authorId, limit, offset int) (*repositories.GetNotesByAuthorIdDto, error) {
			return nil, wantErr
		},
		countNotesByAuthorFn: func(ctx context.Context, authorId int) (int, error) {
			return 10, nil
		},
	}
	svc := NewNotesService(repo)

	got, err := svc.GetAuthorNotes(context.Background(), 1, 1, 20)
	if !errors.Is(err, wantErr) {
		t.Errorf("err = %v, want %v", err, wantErr)
	}
	if got != nil {
		t.Errorf("result = %+v, want nil when the notes query fails", got)
	}
}

func TestNotesService_GetAuthorNotes_CountRepoError(t *testing.T) {
	wantErr := errors.New("count query failed")
	repo := &mockNotesRepo{
		getNotesByAuthorFn: func(ctx context.Context, authorId, limit, offset int) (*repositories.GetNotesByAuthorIdDto, error) {
			return &repositories.GetNotesByAuthorIdDto{Notes: nil}, nil
		},
		countNotesByAuthorFn: func(ctx context.Context, authorId int) (int, error) {
			return 0, wantErr
		},
	}
	svc := NewNotesService(repo)

	got, err := svc.GetAuthorNotes(context.Background(), 1, 1, 20)
	if !errors.Is(err, wantErr) {
		t.Errorf("err = %v, want %v", err, wantErr)
	}
	if got != nil {
		t.Errorf("result = %+v, want nil when the count query fails", got)
	}
}
