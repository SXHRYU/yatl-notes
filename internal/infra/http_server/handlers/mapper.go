package http_server

import (
	"time"

	notes "yat-blog-notes/internal/services"
)

type Note struct {
	Id        int       `json:"id"`
	AuthorId  *int      `json:"author_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

type GetAuthorNotesResponse struct {
	Notes []Note `json:"notes"`
}

func ToResponseAuthorNotes(
	serviceDto *notes.GetNotesByAuthorIdDto,
) *GetAuthorNotesResponse {
	notes := make([]Note, len(serviceDto.Notes))
	for i := range serviceDto.Notes {
		notes[i] = Note(serviceDto.Notes[i])
	}
	return &GetAuthorNotesResponse{
		Notes: notes,
	}
}
