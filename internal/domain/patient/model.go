package patient

import "github.com/Nachofra/final-esp-backend-3/pkg/custom_time"

// Patient describes a patient.
type Patient struct {
	ID            int              `json:"id"`
	FirstName     string           `json:"first_name"`
	LastName      string           `json:"last_name"`
	Address       string           `json:"address"`
	DNI           int              `json:"dni"`
	DischargeDate custom_time.Time `json:"discharge_date"`
}

// NewPatient describes the data needed to create a new Patient.
type NewPatient struct {
	FirstName     string           `json:"first_name"     validate:"required"`
	LastName      string           `json:"last_name"      validate:"required"`
	Address       string           `json:"address"        validate:"required"`
	DNI           int              `json:"dni"            validate:"required,min=10000000,max=99999999"`
	DischargeDate custom_time.Time `json:"discharge_date" validate:"required"`
}

// PatchPatient describes the data needed to patch a Patient.
type PatchPatient struct {
	FirstName     *string           `json:"first_name"`
	LastName      *string           `json:"last_name"`
	Address       *string           `json:"address"`
	DNI           *int              `json:"dni" validate:"min=10000000,max=99999999"`
	DischargeDate *custom_time.Time `json:"discharge_date"`
}
