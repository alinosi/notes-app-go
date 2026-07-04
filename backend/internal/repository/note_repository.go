package repository

import (
	"backend/internal/model"

	"github.com/jmoiron/sqlx"
)

// NoteRepository defines the contract for note database operations
type NoteRepository interface {
	CreateNote(note *model.Note) error
	UpdateNote(note *model.Note) error
	// PatchNote(note *model.Note) error
	DeleteNote(note *model.Note) error
	ReadNote(note *model.Note) error
}

// noteRepositoryImpl is the concrete implementation of NoteRepository
type noteRepositoryImpl struct {
	db *sqlx.DB
}

// CreateNote inserts a new note into the database
func (r *noteRepositoryImpl) CreateNote(note *model.Note) error {
	// The SQL query using Named Parameters (sqlx magic)
	query := `
		INSERT INTO notes (user_id, title, content) 
		VALUES (:user_id, :title, :content) 
		RETURNING id, created_at, updated_at
	`

	// Execute the named query and map the returned values back to the struct
	rows, err := r.db.NamedQuery(query, note)
	if err != nil {
		return err
	}
	defer rows.Close()

	// Fetch the generated ID and timestamps from the RETURNING clause
	if rows.Next() {
		err = rows.StructScan(note)
		if err != nil {
			return err
		}
	}

	return nil
}

// CreateNote inserts a new note into the database
func (r *noteRepositoryImpl) UpdateNote(note *model.Note) error {
	// The SQL query using Named Parameters (sqlx magic)
	query := `
		UPDATE notes (user_id, title, content) 
		VALUES (:user_id, :title, :content)
		WHERE user_id = :user_id 
		RETURNING id, updated_at
	`

	// Execute the named query and map the returned values back to the struct
	rows, err := r.db.NamedQuery(query, note)
	if err != nil {
		return err
	}
	defer rows.Close()

	// Fetch the generated ID and timestamps from the RETURNING clause
	if rows.Next() {
		err = rows.StructScan(note)
		if err != nil {
			return err
		}
	}

	return nil
}

// CreateNote inserts a new note into the database
func (r *noteRepositoryImpl) DeleteNote(note *model.Note) error {
	// The SQL query using Named Parameters (sqlx magic)
	query := `
		DELETE FROM notes WHERE notes_id = :notes_id
		`

	// Execute the named query and map the returned values back to the struct
	rows, err := r.db.NamedQuery(query, note)
	if err != nil {
		return err
	}
	defer rows.Close()

	// Fetch the generated ID and timestamps from the RETURNING clause
	if rows.Next() {
		err = rows.StructScan(note)
		if err != nil {
			return err
		}
	}

	return nil
}

// CreateNote inserts a new note into the database
func (r *noteRepositoryImpl) ReadNote(note *model.Note) error {
	// The SQL query using Named Parameters (sqlx magic)
	query := `
		INSERT INTO notes (user_id, title, content) 
		VALUES (:user_id, :title, :content) 
		RETURNING id, created_at, updated_at
	`

	// Execute the named query and map the returned values back to the struct
	rows, err := r.db.NamedQuery(query, note)
	if err != nil {
		return err
	}
	defer rows.Close()

	// Fetch the generated ID and timestamps from the RETURNING clause
	if rows.Next() {
		err = rows.StructScan(note)
		if err != nil {
			return err
		}
	}

	return nil
}

// NewNoteRepository acts as a constructor to inject the database connection
func NewNoteRepository(db *sqlx.DB) NoteRepository {
	return &noteRepositoryImpl{
		db: db,
	}
}
