package repositories

import (
	"time"
)

type Note struct {
	Id        int       `json:"id"`
	AuthorId  *int      `json:"author_id"`
	Slug      string    `json:"slug"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

type GetNotesByAuthorIdDto struct {
	Notes []Note `json:"notes"`
}
