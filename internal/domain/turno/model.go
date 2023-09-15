package turno

import (
	"time"

	"github.com/Nachofra/final-esp-backend-3/internal/domain/odontologo"
	"github.com/Nachofra/final-esp-backend-3/internal/domain/paciente"
)

// Turno describes a turno.
type Turno struct {
	ID           int                   `json:"id"`
	Paciente     paciente.Paciente     `json:"paciente_id"`
	Odontologo   odontologo.Odontologo `json:"odontologo_id"`
	Date         time.Time             `json:"date"`
	Descrtiption string                `json:"description"`
}

// RequestTurno describes the data needed to create a new turno.
type RequestTurno struct {
	Paciente     paciente.Paciente     `json:"paciente_id"`
	Odontologo   odontologo.Odontologo `json:"odontologo_id"`
	Date         time.Time             `json:"date"`
	Descrtiption string                `json:"description"`
}
