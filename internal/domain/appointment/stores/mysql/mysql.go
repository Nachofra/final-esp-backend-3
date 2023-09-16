package mysql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Nachofra/final-esp-backend-3/internal/domain/appointment"
)

const (
	QueryGetAllAppointment = `SELECT id, patient_id, dentist_id, date, description
	FROM clinic.appointment`

	QueryGetAppointmentByID = `SELECT id, patient_id, dentist_id, date, description
	FROM clinic.appointment WHERE id = ?`

	QueryGetAppointmentByPatientID = `SELECT id, patient_id, dentist_id, date, description
	FROM clinic.appointment WHERE patient_id = ?`

	QueryInsertAppointment = `INSERT INTO clinic.appointment(patient_id,dentist_id,date,description)
	VALUES(?,?,?,?)`

	QueryUpdateAppointment = `UPDATE clinic.appointment SET patient_id = ?, dentist_id = ?, date = ?, description = ?
	WHERE id = ?`

	QueryDeleteAppointment = `DELETE FROM clinic.appointment WHERE id = ?`
)

type Store struct {
	db *sql.DB
}

// New creates a new store.
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// GetAll returns all appointments.
func (s *Store) GetAll(_ context.Context, filters map[string]string) []appointment.Appointment {
	query := GenerateQuery(filters)

	rows, err := s.db.Query(query)
	if err != nil {
		return []appointment.Appointment{}
	}

	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			panic("IMPLEMENT LOGGER")
		}
	}(rows)

	var appointmentsList []appointment.Appointment

	for rows.Next() {
		var a appointment.Appointment

		err = rows.Scan(&a.ID, &a.PatientID, &a.DentistID, &a.Date, &a.Description)
		if err != nil {
			return []appointment.Appointment{}
		}

		appointmentsList = append(appointmentsList, a)
	}

	return appointmentsList
}

// GetByID returns an appointment by its ID.
func (s *Store) GetByID(_ context.Context, ID int) (appointment.Appointment, error) {
	stmt, err := s.db.Prepare(QueryGetAppointmentByID)
	if err != nil {
		// TODO: implement log
		return appointment.Appointment{}, appointment.ErrStatement
	}

	row := stmt.QueryRow(ID)

	var a appointment.Appointment

	err = row.Scan(&a.ID, &a.PatientID, &a.DentistID, &a.Date, &a.Description)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			// TODO: implement log
			err = appointment.ErrNotFound
		}

		return appointment.Appointment{}, err
	}

	return a, nil
}

// Create creates a new appointment.
func (s *Store) Create(_ context.Context, a appointment.Appointment) (appointment.Appointment, error) {
	statement, err := s.db.Prepare(QueryInsertAppointment)
	if err != nil {
		return appointment.Appointment{}, appointment.ErrStatement
	}

	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			panic("IMPLEMENT LOGGER")
		}
	}(statement)

	result, err := statement.Exec(a.PatientID, a.DentistID, a.Date, a.Description)
	if err != nil {
		return appointment.Appointment{}, appointment.ErrExec
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return appointment.Appointment{}, appointment.ErrLastId
	}

	a.ID = int(lastId)

	return a, nil
}

// Update updates an appointment.
func (s *Store) Update(_ context.Context, a appointment.Appointment) (appointment.Appointment, error) {
	statement, err := s.db.Prepare(QueryUpdateAppointment)
	if err != nil {
		return appointment.Appointment{}, appointment.ErrStatement
	}

	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			panic("IMPLEMENT LOGGER")
		}
	}(statement)

	result, err := statement.Exec(a.PatientID, a.DentistID, a.Date, a.Description, a.ID)
	if err != nil {
		return appointment.Appointment{}, appointment.ErrExec
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return appointment.Appointment{}, err
	}

	if rowsAffected < 1 {
		return appointment.Appointment{}, appointment.ErrNotFound
	}

	return a, nil
}

// Delete deletes an appointment.
func (s *Store) Delete(_ context.Context, ID int) error {
	result, err := s.db.Exec(QueryDeleteAppointment, ID)
	if err != nil {
		return appointment.ErrExec
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return appointment.ErrNotFound
	}

	return nil
}
