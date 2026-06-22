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

type CreateNoteResponse struct {
	Id int `json:"id"`
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

func ToResponseCreateNote(newNoteId int) *CreateNoteResponse {
	return &CreateNoteResponse{
		Id: newNoteId,
	}
}
