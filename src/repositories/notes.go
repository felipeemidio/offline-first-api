package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type Notes struct {
	db *sql.DB
}

func CreateNotesRepository(db *sql.DB) *Notes {
	return &Notes{db}
}

func (repository Notes) Create(note models.Note) (uint64, error) {
	statement, err := repository.db.Prepare("insert into notes (content) value(?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(note.Content)
	if err != nil {
		return 0, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedId), nil
}

func (repository Notes) GetAll() ([]models.Note, error) {
	rows, err := repository.db.Query("select * from notes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []models.Note

	for rows.Next() {
		var note models.Note
		if err = rows.Scan(
			&note.ID,
			&note.Content,
			&note.CreatedAt,
		); err != nil {
			return nil, err
		}

		notes = append(notes, note)
	}

	return notes, nil
}

func (repository Notes) GetById(noteId uint64) (models.Note, error) {
	rows, err := repository.db.Query(fmt.Sprintf("select * from notes where id = %d", noteId))
	if err != nil {
		return models.Note{}, err
	}
	defer rows.Close()

	var note models.Note
	if rows.Next() {
		if err = rows.Scan(
			&note.ID,
			&note.Content,
			&note.CreatedAt,
		); err != nil {
			return models.Note{}, err
		}
	}

	return note, nil
}

func (repository Notes) EditNote(noteId uint64, note models.Note) error {
	statement, err := repository.db.Prepare("update notes set content = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(note.Content, noteId); err != nil {
		return err
	}

	return nil
}

func (repository Notes) DeleteNote(noteId uint64) error {
	statement, err := repository.db.Prepare("delete from notes where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(noteId); err != nil {
		return err
	}

	return nil
}
