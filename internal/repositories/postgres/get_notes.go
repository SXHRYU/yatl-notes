package postgres

import (
	"context"
)

func (nr *NotesRepository) GetNotesByAuthorId(
	ctx context.Context,
	authorId, limit, offset int,
) ([]string, error) {
	const query = `SELECT text FROM notes WHERE author_id = $1 LIMIT $2 OFFSET $3;`
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

	var res []string
	for rows.Next() {
		var s string
		if err := rows.Scan(&s); err != nil {
			return nil, err
		}
		res = append(res, s)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return res, nil
}
