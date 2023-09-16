package appointment

import (
	"github.com/Nachofra/final-esp-backend-3/internal/domain/dentist"
	"github.com/Nachofra/final-esp-backend-3/internal/domain/patient"
	"github.com/Nachofra/final-esp-backend-3/pkg/time"
)

// Appointment describes a appointment.
type Appointment struct {
	ID           int             `json:"id"`
	Patient      patient.Patient `json:"patient_id"`
	Dentist      dentist.Dentist `json:"dentist_id"`
	Date         time.Time       `json:"date"`
	Descrtiption string          `json:"description"`
}

// NewAppointment describes the data needed to create a new Appointment.
type NewAppointment struct {
	Patient      patient.Patient `json:"patient_id"`
	Dentist      dentist.Dentist `json:"dentist_id"`
	Date         time.Time       `json:"date"`
	Descrtiption string          `json:"description"`
}
