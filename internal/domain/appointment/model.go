package appointment

import (
	"github.com/Nachofra/final-esp-backend-3/pkg/time"
)

// Appointment describes a appointment.
type Appointment struct {
	ID          int       `json:"id"`
	PatientID   int       `json:"patient_id"`
	DentistID   int       `json:"dentist_id"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}

// NewAppointment describes the data needed to create a new Appointment.
type NewAppointment struct {
	PatientID   int       `json:"patient_id"`
	DentistID   int       `json:"dentist_id"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}

// FilterAppointment describes the data needed to filter an appointment.
type FilterAppointment struct {
	PatientID int       `json:"patient_id"`
	DentistID int       `json:"dentist_id"`
	FromDate  time.Time `json:"from_date"`
	ToDate    time.Time `json:"to_date"`
}
