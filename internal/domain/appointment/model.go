package appointment

import (
	"encoding/json"
	"github.com/Nachofra/final-esp-backend-3/pkg/time"
)

// Appointment describes an Appointment between a dentist and its patient.
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

// UpdateAppointment describes the data needed to update an Appointment.
type UpdateAppointment struct {
	PatientID   int       `json:"patient_id"`
	DentistID   int       `json:"dentist_id"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}

// PatchAppointment describes the data needed to patch an Appointment.
type PatchAppointment struct {
	PatientID   *int       `json:"patient_id"`
	DentistID   *int       `json:"dentist_id"`
	Date        *time.Time `json:"date"`
	Description *string    `json:"description"`
}

// FilterAppointment describes the data needed to filter an Appointment.
type FilterAppointment struct {
	PatientID int       `form:"patient_id"`
	DentistID int       `form:"dentist_id"`
	FromDate  time.Time `form:"from_date"`
	ToDate    time.Time `form:"to_date"`
}

// ToMap parses FilterAppointment to a map[string]string.
func (fa FilterAppointment) ToMap() map[string]string {
	var newMap map[string]string

	b, _ := json.Marshal(fa)

	err := json.Unmarshal(b, &newMap)
	if err != nil {
		return nil
	}

	return newMap
}
