package paciente

import "time"

// Paciente describes a paciente.
type Paciente struct {
	ID             int       `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Address        string    `json:"address"`
	DNI            int       `json:"dni"`
	DischargeDate  time.Time `json:"discharge_date"`
}

// RequestPaciente describes the data needed to create a new paciente.
type RequestPaciente struct {
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Address        string    `json:"address"`
	DNI            int       `json:"dni"`
	DischargeDate  time.Time `json:"discharge_date"`
}