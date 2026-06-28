package postgres

import (
	"context"

	"yat-blog-notes/internal/repositories"
)

func (nr *NotesRepository) GetNotesByAuthorId(
	ctx context.Context,
	authorId, limit, offset int,
) (*repositories.GetNotesByAuthorIdDto, error) {
	const query = `
		SELECT id, author_id, title, text, created_at
		FROM notes WHERE author_id = $1 LIMIT $2 OFFSET $3;
	`
	stmt, err := nr.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx, authorId, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []repositories.Note
	for rows.Next() {
		var note repositories.Note
		if err := rows.Scan(
			&note.Id,
			&note.AuthorId,
			&note.Title,
			&note.Text,
			&note.CreatedAt,
		); err != nil {
			return nil, err
		}
		res = append(res, note)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &repositories.GetNotesByAuthorIdDto{Notes: res}, nil
}
