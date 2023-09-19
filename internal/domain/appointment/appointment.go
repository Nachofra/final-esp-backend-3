package appointment

import (
	"context"
	"errors"
)

var (
	ErrNotFound      = errors.New("appointment not found")
	ErrConflict      = errors.New("constraint conflict while doing an action with the store layer")
	ErrAlreadyExists = errors.New("appointment already exists")
	ErrValueExceeded = errors.New("attribute value exceed type limit")
)

// Store specifies the contract needed for the Store in the Service.
type Store interface {
	GetAll(ctx context.Context, filters map[string]string) []Appointment
	GetByID(ctx context.Context, ID int) (Appointment, error)
	Create(ctx context.Context, appointment Appointment) (Appointment, error)
	Update(ctx context.Context, appointment Appointment) (Appointment, error)
	Delete(ctx context.Context, ID int) error
}

// service unifies all the business operation for the domain.
type service struct {
	store Store
}

// Service specifies the contract needed for the Service.
type Service interface {
	GetAll(ctx context.Context, filters FilterAppointment) []Appointment
	GetByID(ctx context.Context, ID int) (Appointment, error)
	Create(ctx context.Context, newAppointment NewAppointment) (Appointment, error)
	Update(ctx context.Context, ID int, ua UpdateAppointment) (Appointment, error)
	Patch(ctx context.Context, appointment Appointment, pa PatchAppointment) (Appointment, error)
	Delete(ctx context.Context, ID int) error
}

// NewService creates a new service.
func NewService(store Store) Service {
	return &service{
		store: store,
	}
}

// GetAll returns appointments by filter.
func (s *service) GetAll(ctx context.Context, filters FilterAppointment) []Appointment {
	f := filters.ToMap()

	appointments := s.store.GetAll(ctx, f)
	return appointments
}

// GetByID returns an appointment by its ID.
func (s *service) GetByID(ctx context.Context, ID int) (Appointment, error) {
	appointment, err := s.store.GetByID(ctx, ID)
	if err != nil {
		return Appointment{}, err
	}

	return appointment, nil
}

// Create creates a new appointment.
func (s *service) Create(ctx context.Context, newAppointment NewAppointment) (Appointment, error) {
	appointment := Appointment{
		PatientID:   newAppointment.PatientID,
		DentistID:   newAppointment.DentistID,
		Date:        newAppointment.Date,
		Description: newAppointment.Description,
	}

	a, err := s.store.Create(ctx, appointment)
	if err != nil {
		return Appointment{}, err
	}

	return a, nil
}

// Update updates an appointment.
func (s *service) Update(ctx context.Context, ID int, ua UpdateAppointment) (Appointment, error) {
	appointment := Appointment{
		ID:          ID,
		PatientID:   ua.PatientID,
		DentistID:   ua.DentistID,
		Date:        ua.Date,
		Description: ua.Description,
	}

	a, err := s.store.Update(ctx, appointment)
	if err != nil {
		return Appointment{}, err
	}

	return a, nil
}

// Patch patches an appointment.
func (s *service) Patch(ctx context.Context, appointment Appointment, pa PatchAppointment) (Appointment, error) {
	if pa.PatientID != nil {
		appointment.PatientID = *pa.PatientID
	}

	if pa.DentistID != nil {
		appointment.DentistID = *pa.DentistID
	}

	if pa.Date != nil {
		appointment.Date = *pa.Date
	}

	if pa.Description != nil {
		appointment.Description = *pa.Description
	}

	a, err := s.store.Update(ctx, appointment)
	if err != nil {
		return Appointment{}, err
	}

	return a, nil
}

// Delete deletes an appointment.
func (s *service) Delete(ctx context.Context, ID int) error {
	err := s.store.Delete(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}
