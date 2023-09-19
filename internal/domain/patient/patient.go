package patient

import (
	"context"
	"errors"
)

var (
	ErrNotFound      = errors.New("patient not found")
	ErrConflict      = errors.New("constraint conflict while storing")
	ErrAlreadyExists = errors.New("patient already exists, dni must be unique")
	ErrValueExceeded = errors.New("attribute value exceed type limit")
)

type Store interface {
	Create(ctx context.Context, patient Patient) (Patient, error)
	GetAll(ctx context.Context) []Patient
	GetByID(ctx context.Context, id int) (Patient, error)
	GetByDNI(ctx context.Context, dni int) (Patient, error)
	Update(ctx context.Context, patient Patient) (Patient, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	store Store
}

// Service specifies the contract needed for the Service.
type Service interface {
	Create(ctx context.Context, newPatient NewPatient) (Patient, error)
	GetAll(ctx context.Context) []Patient
	GetByID(ctx context.Context, id int) (Patient, error)
	GetByDNI(ctx context.Context, dni int) (Patient, error)
	Update(ctx context.Context, newPatient NewPatient, id int) (Patient, error)
	Patch(ctx context.Context, patient Patient, pp PatchPatient) (Patient, error)
	Delete(ctx context.Context, id int) error
}

// NewService creates a new service.
func NewService(store Store) Service {
	return &service{
		store: store,
	}
}

// Create creates a new patient.
func (s *service) Create(ctx context.Context, newPatient NewPatient) (Patient, error) {
	patient := requestToPatient(newPatient)

	response, err := s.store.Create(ctx, patient)
	if err != nil {
		return Patient{}, err
	}

	return response, nil
}

// GetAll returns all patients.
func (s *service) GetAll(ctx context.Context) []Patient {
	patients := s.store.GetAll(ctx)
	return patients
}

// GetByID returns a patient by its ID.
func (s *service) GetByID(ctx context.Context, id int) (Patient, error) {
	patient, err := s.store.GetByID(ctx, id)
	if err != nil {
		return Patient{}, err
	}

	return patient, nil
}

// GetByDNI returns a patient by its DNI.
func (s *service) GetByDNI(ctx context.Context, dni int) (Patient, error) {
	patient, err := s.store.GetByDNI(ctx, dni)
	if err != nil {
		return Patient{}, err
	}

	return patient, nil
}

// Update updates a patient.
func (s *service) Update(ctx context.Context, newPatient NewPatient, id int) (Patient, error) {
	patient := requestToPatient(newPatient)
	patient.ID = id

	response, err := s.store.Update(ctx, patient)
	if err != nil {
		return Patient{}, err
	}

	return response, nil
}

// Patch patches a patient.
func (s *service) Patch(ctx context.Context, patient Patient, pp PatchPatient) (Patient, error) {
	if pp.FirstName != nil {
		patient.FirstName = *pp.FirstName
	}

	if pp.LastName != nil {
		patient.LastName = *pp.LastName
	}

	if pp.Address != nil {
		patient.Address = *pp.Address
	}

	if pp.DNI != nil {
		patient.DNI = *pp.DNI
	}

	if pp.DischargeDate != nil {
		patient.DischargeDate = *pp.DischargeDate
	}

	p, err := s.store.Update(ctx, patient)
	if err != nil {
		return Patient{}, err
	}

	return p, nil
}

// Delete deletes a patient.
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.store.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// requestToPatient parses NewPatient to Patient
func requestToPatient(newPatient NewPatient) Patient {
	var patient Patient
	patient.FirstName = newPatient.FirstName
	patient.LastName = newPatient.LastName
	patient.Address = newPatient.Address
	patient.DNI = newPatient.DNI
	patient.DischargeDate = newPatient.DischargeDate

	return patient
}
