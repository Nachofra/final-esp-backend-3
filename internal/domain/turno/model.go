package turno

import "time"

// Turno describes a turno.
type Turno struct {
	ID             int         `json:"id"`
	Paciente       paciente    `json:"paciente_id"`
	Odontologo     odontologo  `json:"odontologo_id"`
	Date           time.Time   `json:"date"`
	Descrtiption   string      `json:"description"`
}

// RequestTurno describes the data needed to create a new turno.
type RequestPaciente struct {
	Paciente       paciente    `json:"paciente_id"`
	Odontologo     odontologo  `json:"odontologo_id"`
	Date           time.Time   `json:"date"`
	Descrtiption   string      `json:"description"`
}