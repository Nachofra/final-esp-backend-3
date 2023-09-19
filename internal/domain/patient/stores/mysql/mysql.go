package patient

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Nachofra/final-esp-backend-3/internal/domain/patient"
	"github.com/Nachofra/final-esp-backend-3/pkg/mysql"
	"log"
)

var (
	QueryInsertPatient = `INSERT INTO clinic.patient(first_name,last_name,address,dni,discharge_date)
	VALUES(?,?,?,?,?)`
	QueryGetAllPatient = `SELECT id, first_name, last_name, address, dni, discharge_date
	FROM clinic.patient`
	QueryDeletePatient  = `DELETE FROM clinic.patient WHERE id = ?`
	QueryGetPatientByID = `SELECT id, first_name, last_name, address, dni, discharge_date
	FROM clinic.patient WHERE id = ?`
	QueryGetPatientByDNI = `SELECT id, first_name, last_name, address, dni, discharge_date
	FROM clinic.patient WHERE dni = ?`
	QueryUpdatePatient = `UPDATE clinic.patient SET first_name = ?, last_name = ?, address = ? , dni = ?, discharge_date = ?
	WHERE id = ?`
)

type Store struct {
	db *sql.DB
}

// NewStore creates a new repository.
func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// GetAll returns all patients.
func (s *Store) GetAll(_ context.Context) []patient.Patient {
	rows, err := s.db.Query(QueryGetAllPatient)
	if err != nil {
		return []patient.Patient{}
	}

	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			log.Println(err)
		}
	}(rows)

	patientsList := make([]patient.Patient, 0)

	for rows.Next() {
		var p patient.Patient

		err = rows.Scan(
			&p.ID,
			&p.FirstName,
			&p.LastName,
			&p.Address,
			&p.DNI,
			&p.DischargeDate.Time,
		)
		if err != nil {
			return []patient.Patient{}
		}

		patientsList = append(patientsList, p)
	}

	return patientsList
}

// GetByID returns a patient by its ID.
func (s *Store) GetByID(_ context.Context, id int) (patient.Patient, error) {
	row := s.db.QueryRow(QueryGetPatientByID, id)

	var p patient.Patient

	err := row.Scan(
		&p.ID,
		&p.FirstName,
		&p.LastName,
		&p.Address,
		&p.DNI,
		&p.DischargeDate.Time,
	)
	if err != nil {
		err := mysql.CheckError(err)
		switch {
		case errors.Is(err, mysql.ErrDBNoRows):
			return patient.Patient{}, patient.ErrNotFound
		default:
			return patient.Patient{}, err
		}
	}

	return p, nil
}

// GetByDNI returns a patient by its DNI.
func (s *Store) GetByDNI(_ context.Context, dni int) (patient.Patient, error) {
	row := s.db.QueryRow(QueryGetPatientByDNI, dni)

	var p patient.Patient

	err := row.Scan(
		&p.ID,
		&p.FirstName,
		&p.LastName,
		&p.Address,
		&p.DNI,
		&p.DischargeDate.Time,
	)
	if err != nil {
		err := mysql.CheckError(err)
		switch {
		case errors.Is(err, mysql.ErrDBNoRows):
			return patient.Patient{}, patient.ErrNotFound
		default:
			return patient.Patient{}, err
		}
	}

	return p, nil
}

// Create creates a new patient.
func (s *Store) Create(_ context.Context, p patient.Patient) (patient.Patient, error) {
	statement, err := s.db.Prepare(QueryInsertPatient)
	if err != nil {
		return patient.Patient{}, err
	}

	defer func(statement *sql.Stmt) {
		err = statement.Close()
		if err != nil {
			log.Println(err)
		}
	}(statement)

	result, err := statement.Exec(
		p.FirstName,
		p.LastName,
		p.Address,
		p.DNI,
		p.DischargeDate.Time,
	)
	if err != nil {
		err := mysql.CheckError(err)
		switch {
		case errors.Is(err, mysql.ErrDBDuplicateEntry):
			return patient.Patient{}, patient.ErrAlreadyExists
		case errors.Is(err, mysql.ErrDBConflict):
			return patient.Patient{}, patient.ErrConflict
		case errors.Is(err, mysql.ErrDBValueExceeded):
			return patient.Patient{}, patient.ErrValueExceeded
		default:
			return patient.Patient{}, err
		}
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return patient.Patient{}, err
	}

	p.ID = int(lastId)

	return p, nil
}

// Update updates a patient.
func (s *Store) Update(_ context.Context, p patient.Patient) (patient.Patient, error) {
	statement, err := s.db.Prepare(QueryUpdatePatient)
	if err != nil {
		return patient.Patient{}, err
	}

	defer func(statement *sql.Stmt) {
		err = statement.Close()
		if err != nil {
			log.Println(err)
		}
	}(statement)

	result, err := statement.Exec(
		p.FirstName,
		p.LastName,
		p.Address,
		p.DNI,
		p.DischargeDate.Time,
		p.ID,
	)
	if err != nil {
		err := mysql.CheckError(err)
		switch {
		case errors.Is(err, mysql.ErrDBDuplicateEntry):
			return patient.Patient{}, patient.ErrAlreadyExists
		case errors.Is(err, mysql.ErrDBConflict):
			return patient.Patient{}, patient.ErrConflict
		case errors.Is(err, mysql.ErrDBValueExceeded):
			return patient.Patient{}, patient.ErrValueExceeded
		default:
			return patient.Patient{}, err
		}
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return patient.Patient{}, err
	}

	if rowsAffected < 1 {
		return patient.Patient{}, patient.ErrNotFound
	}

	return p, nil
}

// Delete deletes a patient.
func (s *Store) Delete(_ context.Context, id int) error {
	result, err := s.db.Exec(QueryDeletePatient, id)
	if err != nil {
		err := mysql.CheckError(err)
		switch {
		case errors.Is(err, mysql.ErrDBConflict):
			return patient.ErrConflict
		default:
			return err
		}
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return patient.ErrNotFound
	}

	return nil
}
