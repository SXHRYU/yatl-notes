package postgres

import "context"

func (nr *NotesRepository) CreateNote(
	ctx context.Context,
	authorId int,
	text string,
) (int, error) {
	query := "INSERT INTO notes(author_id, text) VALUES($1, $2) RETURNING id;"
	stmt, err := nr.db.PrepareContext(ctx, query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var newId int
	if err := stmt.QueryRowContext(ctx, authorId, text).Scan(&newId); err != nil {
		return 0, err
	}
	return newId, nil
}
