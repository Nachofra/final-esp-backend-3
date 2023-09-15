package turno

import (
	"time"

	"github.com/Nachofra/final-esp-backend-3/internal/domain/odontologo"
	"github.com/Nachofra/final-esp-backend-3/internal/domain/paciente"
)

// Appointment describes a appointment.
type Appointment struct {
	ID           int                   `json:"id"`
	Paciente     paciente.Paciente     `json:"paciente_id"`
	Odontologo   odontologo.Odontologo `json:"odontologo_id"`
	Date         time.Time             `json:"date"`
	Descrtiption string                `json:"description"`
}

// NewAppointment describes the data needed to create a new Appointment.
type NewAppointment struct {
	Paciente     paciente.Paciente     `json:"paciente_id"`
	Odontologo   odontologo.Odontologo `json:"odontologo_id"`
	Date         time.Time             `json:"date"`
	Descrtiption string                `json:"description"`
}
