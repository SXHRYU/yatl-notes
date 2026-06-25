package postgres

import (
	"context"

	"yat-blog-notes/internal/repositories"
)

func (nr *NotesRepository) GetNote(
	ctx context.Context,
	noteId int,
) (*repositories.Note, error) {
	query := `SELECT * FROM notes WHERE id = $1;`
	stmt, err := nr.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var res repositories.Note
	if err := stmt.QueryRowContext(ctx, noteId).Scan(
		&res.Id,
		&res.AuthorId,
		&res.Text,
		&res.CreatedAt,
		&res.Title,
	); err != nil {
		return nil, err
	}

	return &res, nil
}
