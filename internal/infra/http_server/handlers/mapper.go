package http_server

import (
	"time"

	notes "yat-blog-notes/internal/services"
)

// not a lot of records in db so it's ok
type PagedPagination struct {
	Page  int `json:"page"`
	Pages int `json:"pages"`
	Total int `json:"total"`
}

type Note struct {
	Id        int       `json:"id"`
	AuthorId  *int      `json:"author_id"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

type GetAuthorNotesResponse struct {
	Notes      []Note          `json:"notes"`
	Pagination PagedPagination `json:"pagination"`
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
		Pagination: PagedPagination{
			Page:  serviceDto.Pagination.Page,
			Pages: serviceDto.Pagination.Pages,
			Total: serviceDto.Pagination.Total,
		},
	}
}

func ToResponseCreateNote(newNoteId int) *CreateNoteResponse {
	return &CreateNoteResponse{
		Id: newNoteId,
	}
}

func ToResponseGetNote(note *notes.Note) *Note {
	return &Note{
		Id:        note.Id,
		AuthorId:  note.AuthorId,
		Title:     note.Title,
		Text:      note.Text,
		CreatedAt: note.CreatedAt,
	}
}
