package postgres

import "context"

func (nr *NotesRepository) CountNotesByAuthorId(
	ctx context.Context,
	authorId int,
) (int, error) {
	query := `SELECT COUNT(*) FROM notes WHERE author_id = $1;`
	stmt, err := nr.db.PrepareContext(ctx, query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var total int
	err = stmt.QueryRowContext(ctx, authorId).Scan(&total)
	return total, err
}
