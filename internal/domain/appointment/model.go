package appointment

import (
	"github.com/Nachofra/final-esp-backend-3/pkg/custom_time"
	"strconv"
	"time"
)

// Appointment describes an Appointment between a dentist and its patient.
type Appointment struct {
	ID          int              `json:"id"`
	PatientID   int              `json:"patient_id"`
	DentistID   int              `json:"dentist_id"`
	Date        custom_time.Time `json:"date"`
	Description string           `json:"description"`
}

// NewAppointment describes the data needed to create a new Appointment.
type NewAppointment struct {
	PatientID   int              `json:"patient_id"  validate:"required"`
	DentistID   int              `json:"dentist_id"  validate:"required"`
	Date        custom_time.Time `json:"date"        validate:"required"`
	Description string           `json:"description" validate:"required"`
}

// UpdateAppointment describes the data needed to update an Appointment.
type UpdateAppointment struct {
	PatientID   int              `json:"patient_id"  validate:"required"`
	DentistID   int              `json:"dentist_id"  validate:"required"`
	Date        custom_time.Time `json:"date"        validate:"required"`
	Description string           `json:"description" validate:"required"`
}

// PatchAppointment describes the data needed to patch an Appointment.
type PatchAppointment struct {
	PatientID   *int              `json:"patient_id"`
	DentistID   *int              `json:"dentist_id"`
	Date        *custom_time.Time `json:"date"`
	Description *string           `json:"description"`
}

// FilterAppointment describes the data needed to filter an Appointment.
type FilterAppointment struct {
	PatientID *int              `form:"patient_id"`
	DentistID *int              `form:"dentist_id"`
	DNI       *int              `form:"dni" validate:"min=10000000,max=999999999"`
	FromDate  *custom_time.Time `form:"from_date"`
	ToDate    *custom_time.Time `form:"to_date"`
}

// ToMap parses FilterAppointment to a map[string]string.
func (fa FilterAppointment) ToMap() map[string]string {
	newMap := make(map[string]string)

	if fa.PatientID != nil {
		newMap["patient_id"] = strconv.Itoa(*fa.PatientID)
	}

	if fa.DentistID != nil {
		newMap["dentist_id"] = strconv.Itoa(*fa.DentistID)
	}

	if fa.DNI != nil {
		newMap["dni"] = strconv.Itoa(*fa.DNI)
	}

	if fa.FromDate != nil {
		newMap["from_date"] = fa.FromDate.Format(time.DateTime)
	}

	if fa.ToDate != nil {
		newMap["to_date"] = fa.ToDate.Format(time.DateTime)
	}

	return newMap
}
