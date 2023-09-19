package mysql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Nachofra/final-esp-backend-3/internal/domain/appointment"
	"github.com/Nachofra/final-esp-backend-3/pkg/mysql"
	"log"
)

// Store wraps all the operations to the database.
type Store struct {
	db *sql.DB
}

// NewStore creates a new store.
func NewStore(db *sql.DB) *Store {
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
			log.Println(err)
		}
	}(rows)

	appointmentsList := make([]appointment.Appointment, 0)

	for rows.Next() {
		var a appointment.Appointment

		err = rows.Scan(&a.ID, &a.PatientID, &a.DentistID, &a.Date.Time, &a.Description)
		if err != nil {
			return []appointment.Appointment{}
		}

		appointmentsList = append(appointmentsList, a)
	}

	return appointmentsList
}

// GetByID returns an appointment by its ID.
func (s *Store) GetByID(_ context.Context, ID int) (appointment.Appointment, error) {
	row := s.db.QueryRow(QueryGetAppointmentByID, ID)

	var a appointment.Appointment

	err := row.Scan(&a.ID, &a.PatientID, &a.DentistID, &a.Date.Time, &a.Description)
	if err != nil {
		err := mysql.CheckError(err)
		switch {
		case errors.Is(err, mysql.ErrDBNoRows):
			return appointment.Appointment{}, appointment.ErrNotFound
		default:
			return appointment.Appointment{}, err
		}
	}

	return a, nil
}

// Create creates a new appointment.
func (s *Store) Create(_ context.Context, a appointment.Appointment) (appointment.Appointment, error) {
	statement, err := s.db.Prepare(QueryInsertAppointment)
	if err != nil {
		return appointment.Appointment{}, err
	}

	defer func(statement *sql.Stmt) {
		err = statement.Close()
		if err != nil {
			log.Println(err)
		}
	}(statement)

	result, err := statement.Exec(a.PatientID, a.DentistID, a.Date.Time, a.Description)
	if err != nil {
		err := mysql.CheckError(err)
		switch {
		case errors.Is(err, mysql.ErrDBDuplicateEntry):
			return appointment.Appointment{}, appointment.ErrAlreadyExists
		case errors.Is(err, mysql.ErrDBConflict):
			return appointment.Appointment{}, appointment.ErrConflict
		case errors.Is(err, mysql.ErrDBValueExceeded):
			return appointment.Appointment{}, appointment.ErrValueExceeded
		default:
			return appointment.Appointment{}, err
		}
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return appointment.Appointment{}, err
	}

	a.ID = int(lastId)

	return a, nil
}

// Update updates an appointment.
func (s *Store) Update(_ context.Context, a appointment.Appointment) (appointment.Appointment, error) {
	statement, err := s.db.Prepare(QueryUpdateAppointment)
	if err != nil {
		return appointment.Appointment{}, err
	}

	defer func(statement *sql.Stmt) {
		err = statement.Close()
		if err != nil {
			log.Println(err)
		}
	}(statement)

	result, err := statement.Exec(a.PatientID, a.DentistID, a.Date.Time, a.Description, a.ID)
	if err != nil {
		err := mysql.CheckError(err)
		switch {
		case errors.Is(err, mysql.ErrDBDuplicateEntry):
			return appointment.Appointment{}, appointment.ErrAlreadyExists
		case errors.Is(err, mysql.ErrDBConflict):
			return appointment.Appointment{}, appointment.ErrConflict
		case errors.Is(err, mysql.ErrDBValueExceeded):
			return appointment.Appointment{}, appointment.ErrValueExceeded
		default:
			return appointment.Appointment{}, err
		}
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
		err := mysql.CheckError(err)
		switch {
		case errors.Is(err, mysql.ErrDBConflict):
			return appointment.ErrConflict
		default:
			return err
		}
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
