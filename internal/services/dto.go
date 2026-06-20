package notes

import (
	"time"
)

type Note struct {
	Id        int       `json:"id"`
	AuthorId  *int      `json:"author_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

type GetNotesByAuthorIdDto struct {
	Notes []Note `json:"notes"`
}
