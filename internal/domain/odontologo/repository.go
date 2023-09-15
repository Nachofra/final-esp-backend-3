package odontologo

import (
	"context"
	"database/sql"
)

type repository struct {
	db *sql.DB
}

// NewRepositoryMySql creates a new repository.
func NewRepositoryMySql(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

// Create creates a new odontologo.
func (r *repository) Create(ctx context.Context, odontologo Odontologo) (Odontologo, error) {

	statement, err := r.db.Prepare(QueryInsertOdontologo)

	if err != nil {
		return Odontologo{}, ErrStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		odontologo.FirstName,
		odontologo.LastName,
		odontologo.RegistrationNumber,
	)

	if err != nil {
		return Odontologo{}, ErrExec
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return Odontologo{}, ErrLastId
	}

	odontologo.ID = int(lastId)

	return odontologo, nil
}

// GetAll returns all odontologos.
func (r *repository) GetAll(ctx context.Context) ([]Odontologo, error) {
	rows, err := r.db.Query(QueryGetAllOdontologo)
	if err != nil {
		return []Odontologo{}, err
	}

	defer rows.Close()

	var odontologos []Odontologo

	for rows.Next() {
		var odontologo Odontologo
		err := rows.Scan(
			&odontologo.ID,
			&odontologo.FirstName,
			&odontologo.LastName,
			&odontologo.RegistrationNumber,
		)
		if err != nil {
			return []Odontologo{}, err
		}

		odontologos = append(odontologos, odontologo)
	}

	if err := rows.Err(); err != nil {
		return []Odontologo{}, err
	}

	return odontologos, nil
}

// GetByID returns a odontologo by its ID.
func (r *repository) GetByID(ctx context.Context, id int) (Odontologo, error) {
	row := r.db.QueryRow(QueryGetOdontologoById, id)

	var odontologo Odontologo
	err := row.Scan(
		&odontologo.ID,
		&odontologo.FirstName,
		&odontologo.LastName,
		&odontologo.RegistrationNumber,
	)

	if err != nil {
		return Odontologo{}, err
	}

	return odontologo, nil
}

// Update updates a odontologo.
func (r *repository) Update(ctx context.Context, odontologo Odontologo) (Odontologo, error) {
	statement, err := r.db.Prepare(QueryUpdateOdontologo)
	if err != nil {
		return Odontologo{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		odontologo.FirstName,
		odontologo.LastName,
		odontologo.RegistrationNumber,
	)

	if err != nil {
		return Odontologo{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return Odontologo{}, err
	}

	if rowsAffected < 1 {
		return Odontologo{}, ErrNotFound
	}

	return odontologo, nil
}

// Delete deletes a odontologo.
func (r *repository) Delete(ctx context.Context, id int) error {
	result, err := r.db.Exec(QueryDeleteOdontologo, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return ErrNotFound
	}

	return nil

}
