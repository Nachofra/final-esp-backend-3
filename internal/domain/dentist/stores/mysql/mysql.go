package mysql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Nachofra/final-esp-backend-3/pkg/mysql"
	"log"

	"github.com/Nachofra/final-esp-backend-3/internal/domain/dentist"
)

const (
	QueryGetAllDentist = `SELECT id, first_name, last_name, registration_number
	FROM clinic.dentist`

	QueryGetDentistById = `SELECT id, first_name, last_name, registration_number
	FROM clinic.dentist WHERE id = ?`

	QueryGetDentistByRegistrationNumber = `SELECT id, first_name, last_name, registration_number
	FROM clinic.dentist WHERE registration_number = ?`

	QueryInsertDentist = `INSERT INTO clinic.dentist(first_name,last_name,registration_number)
	VALUES(?,?,?)`

	QueryUpdateDentist = `UPDATE clinic.dentist SET first_name = ?, last_name = ?, registration_number = ?
	WHERE id = ?`

	QueryDeleteDentist = `DELETE FROM clinic.dentist WHERE id = ?`
)

type Store struct {
	db *sql.DB
}

// New creates a new repository.
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// GetAll returns all dentists.
func (s *Store) GetAll(_ context.Context) []dentist.Dentist {
	rows, err := s.db.Query(QueryGetAllDentist)
	if err != nil {
		return []dentist.Dentist{}
	}

	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			log.Println(err)
		}
	}(rows)

	dentistsList := make([]dentist.Dentist, 0)

	for rows.Next() {
		var d dentist.Dentist

		err := rows.Scan(
			&d.ID,
			&d.FirstName,
			&d.LastName,
			&d.RegistrationNumber,
		)
		if err != nil {
			return []dentist.Dentist{}
		}

		dentistsList = append(dentistsList, d)
	}

	return dentistsList
}

// GetByID returns a dentist by its ID.
func (s *Store) GetByID(_ context.Context, ID int) (dentist.Dentist, error) {
	row := s.db.QueryRow(QueryGetDentistById, ID)

	var d dentist.Dentist

	err := row.Scan(
		&d.ID,
		&d.FirstName,
		&d.LastName,
		&d.RegistrationNumber,
	)
	if err != nil {
		err := mysql.CheckError(err)
		switch {
		case errors.Is(err, mysql.ErrDBNoRows):
			return dentist.Dentist{}, dentist.ErrNotFound
		default:
			return dentist.Dentist{}, err
		}
	}

	return d, nil
}

// GetByRegistrationNumber returns a dentist by its RegistrationNumber.
func (s *Store) GetByRegistrationNumber(_ context.Context, rn int) (dentist.Dentist, error) {
	row := s.db.QueryRow(QueryGetDentistByRegistrationNumber, rn)

	var d dentist.Dentist

	err := row.Scan(
		&d.ID,
		&d.FirstName,
		&d.LastName,
		&d.RegistrationNumber,
	)
	if err != nil {
		err := mysql.CheckError(err)
		switch {
		case errors.Is(err, mysql.ErrDBNoRows):
			return dentist.Dentist{}, dentist.ErrNotFound
		default:
			return dentist.Dentist{}, err
		}
	}

	return d, nil
}

// Create creates a new dentist.
func (s *Store) Create(_ context.Context, d dentist.Dentist) (dentist.Dentist, error) {
	statement, err := s.db.Prepare(QueryInsertDentist)
	if err != nil {
		return dentist.Dentist{}, err
	}

	defer func(statement *sql.Stmt) {
		err = statement.Close()
		if err != nil {
			log.Println(err)
		}
	}(statement)

	result, err := statement.Exec(
		d.FirstName,
		d.LastName,
		d.RegistrationNumber,
	)
	if err != nil {
		err := mysql.CheckError(err)
		switch {
		case errors.Is(err, mysql.ErrDBDuplicateEntry):
			return dentist.Dentist{}, dentist.ErrAlreadyExists
		case errors.Is(err, mysql.ErrDBConflict):
			return dentist.Dentist{}, dentist.ErrConflict
		case errors.Is(err, mysql.ErrDBValueExceeded):
			return dentist.Dentist{}, dentist.ErrValueExceeded
		default:
			return dentist.Dentist{}, err
		}
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return dentist.Dentist{}, err
	}

	d.ID = int(lastId)

	return d, nil
}

// Update updates a dentist.
func (s *Store) Update(_ context.Context, d dentist.Dentist) (dentist.Dentist, error) {
	statement, err := s.db.Prepare(QueryUpdateDentist)
	if err != nil {
		return dentist.Dentist{}, err
	}

	defer func(statement *sql.Stmt) {
		err = statement.Close()
		if err != nil {
			log.Println(err)
		}
	}(statement)

	result, err := statement.Exec(
		d.FirstName,
		d.LastName,
		d.RegistrationNumber,
		d.ID,
	)
	if err != nil {
		err := mysql.CheckError(err)
		switch {
		case errors.Is(err, mysql.ErrDBConflict):
			return dentist.Dentist{}, dentist.ErrConflict
		case errors.Is(err, mysql.ErrDBValueExceeded):
			return dentist.Dentist{}, dentist.ErrValueExceeded
		default:
			return dentist.Dentist{}, err
		}
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return dentist.Dentist{}, err
	}

	if rowsAffected < 1 {
		return dentist.Dentist{}, dentist.ErrNotFound
	}

	return d, nil
}

// Delete deletes a dentist.
func (s *Store) Delete(_ context.Context, id int) error {
	result, err := s.db.Exec(QueryDeleteDentist, id)
	if err != nil {
		err := mysql.CheckError(err)
		switch {
		case errors.Is(err, mysql.ErrDBConflict):
			return dentist.ErrConflict
		default:
			return err
		}
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return dentist.ErrNotFound
	}

	return nil
}
