package notes

import (
	"time"
)

type PagedPagination struct {
	Page  int `json:"page"`
	Pages int `json:"pages"`
	Total int `json:"total"`
}

type Note struct {
	Id        int       `json:"id"`
	AuthorId  *int      `json:"author_id"`
	Slug      string    `json:"slug"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

type GetNotesByAuthorIdDto struct {
	Notes      []Note          `json:"notes"`
	Pagination PagedPagination `json:"pagination"`
}
