package appointment

import (
	"github.com/Nachofra/final-esp-backend-3/internal/domain/dentist"
	"github.com/Nachofra/final-esp-backend-3/internal/domain/patient"
	"github.com/Nachofra/final-esp-backend-3/pkg/time"
)

// Appointment describes a appointment.
type Appointment struct {
	ID           int             `json:"id"`
	Paciente     patient.Patient `json:"patient_id"`
	Odontologo   dentist.Dentist `json:"dentist_id"`
	Date         time.Time       `json:"date"`
	Descrtiption string          `json:"description"`
}

// NewAppointment describes the data needed to create a new Appointment.
type NewAppointment struct {
	Paciente     patient.Patient `json:"patient_id"`
	Odontologo   dentist.Dentist `json:"dentist_id"`
	Date         time.Time       `json:"date"`
	Descrtiption string          `json:"description"`
}
