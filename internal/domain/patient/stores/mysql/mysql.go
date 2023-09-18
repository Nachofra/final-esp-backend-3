package patient

import (
	"context"
	"database/sql"

	"github.com/Nachofra/final-esp-backend-3/internal/domain/patient"
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
func (s *Store) GetAll(_ context.Context) ([]patient.Patient, error) {
	rows, err := s.db.Query(QueryGetAllPatient)
	if err != nil {
		return []patient.Patient{}, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic("IMPLEMENT LOGGER")
		}
	}(rows)

	var patientsList []patient.Patient

	for rows.Next() {
		var p patient.Patient

		err := rows.Scan(
			&p.ID,
			&p.FirstName,
			&p.LastName,
			&p.Address,
			&p.DNI,
			&p.DischargeDate,
		)
		if err != nil {
			return []patient.Patient{}, err
		}
		patientsList = append(patientsList, p)
	}
	if err := rows.Err(); err != nil {
		return []patient.Patient{}, err
	}
	return patientsList, nil
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
		&p.DischargeDate,
	)
	if err != nil {
		return patient.Patient{}, err
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
		&p.DischargeDate,
	)
	if err != nil {
		return patient.Patient{}, err
	}
	return p, nil
}

// Create creates a new patient.
func (s *Store) Create(_ context.Context, p patient.Patient) (patient.Patient, error) {
	statement, err := s.db.Prepare(QueryInsertPatient)
	if err != nil {
		return patient.Patient{}, patient.ErrStatement
	}
	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			panic("IMPLEMENT LOGGER")
		}
	}(statement)
	result, err := statement.Exec(
		p.FirstName,
		p.LastName,
		p.Address,
		p.DNI,
		p.DischargeDate,
	)
	if err != nil {
		return patient.Patient{}, patient.ErrExec
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return patient.Patient{}, patient.ErrLastId
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
		err := statement.Close()
		if err != nil {
			panic("IMPLEMENT LOGGER")
		}
	}(statement)
	result, err := statement.Exec(
		p.FirstName,
		p.LastName,
		p.Address,
		p.DNI,
		p.DischargeDate,
		p.ID,
	)
	if err != nil {
		return patient.Patient{}, err
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
		return err
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
