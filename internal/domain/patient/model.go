package patient

import "github.com/Nachofra/final-esp-backend-3/pkg/time"

// Patient describes a patient.
type Patient struct {
	ID            int       `json:"id"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Address       string    `json:"address"`
	DNI           int       `json:"dni"`
	DischargeDate time.Time `json:"discharge_date"`
}

// NewPatient describes the data needed to create a new Patient.
type NewPatient struct {
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Address       string    `json:"address"`
	DNI           int       `json:"dni"`
	DischargeDate time.Time `json:"discharge_date"`
}
