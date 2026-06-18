package postgres

type NotesRepository struct {
	db *DB
}

func NewNotesRepository(db *DB) *NotesRepository {
	return &NotesRepository{
		db: db,
	}
}
