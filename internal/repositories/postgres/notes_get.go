package postgres

import (
	"context"

	"yat-blog-notes/internal/repositories"
)

func (nr *NotesRepository) GetNote(
	ctx context.Context,
	noteId int,
) (*repositories.Note, error) {
	const query = `
		SELECT id, author_id, title, text, created_at
		FROM notes WHERE id = $1;
	`
	stmt, err := nr.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var res repositories.Note
	if err := stmt.QueryRowContext(ctx, noteId).Scan(
		&res.Id,
		&res.AuthorId,
		&res.Title,
		&res.Text,
		&res.CreatedAt,
	); err != nil {
		return nil, err
	}

	return &res, nil
}
