package patient

import (
	"context"
	"errors"
	"log"
)

var (
	ErrEmptyList = errors.New("the list is empty")
	ErrNotFound  = errors.New("Patient not found")
	ErrStatement = errors.New("error preparing statement")
	ErrExec      = errors.New("error exect statement")
	ErrLastId    = errors.New("error getting last id")
)

type Store interface {
	Create(ctx context.Context, patient Patient) (Patient, error)
	GetAll(ctx context.Context) ([]Patient, error)
	GetByID(ctx context.Context, id int) (Patient, error)
	Update(ctx context.Context, patient Patient) (Patient, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	store Store
}

// Create creates a new patient.
func (s *service) Create(ctx context.Context, newPatient NewPatient) (Patient, error) {
	patient := requestToPatient(newPatient)
	response, err := s.store.Create(ctx, patient)
	if err != nil {
		log.Println("error creating all patients on service layer")
		return Patient{}, errors.New("service error. Method Create")
	}

	return response, nil
}

// GetAll returns all patients.
func (s *service) GetAll(ctx context.Context) ([]Patient, error) {
	patients, err := s.store.GetAll(ctx)
	if err != nil {
		log.Println("error gettin all patients on service layer", err.Error())
		return []Patient{}, errors.New("service error. Method GetAll")
	}
	return patients, nil
}

// GetByID returns a product by its ID.
func (s *service) GetByID(ctx context.Context, id int) (Patient, error) {
	patient, err := s.store.GetByID(ctx, id)
	if err != nil {
		log.Println("error getting patient on service layer", err.Error())
		return Patient{}, errors.New("service error. Method GetByID")
	}

	return patient, nil
}

// Update updates a product.
func (s *service) Update(ctx context.Context, newPatient NewPatient, id int) (Patient, error) {
	patient := requestToPatient(newPatient)
	patient.ID = id
	response, err := s.store.Update(ctx, patient)
	if err != nil {
		log.Println("error updating patient on service layer", err.Error())
		return Patient{}, errors.New("service error. Method Update")
	}

	return response, nil
}

// Delete deletes a product.
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.store.Delete(ctx, id)
	if err != nil {
		log.Println("error deleting patient on service layer", err.Error())
		return errors.New("service error. Method Delete")
	}

	return nil
}

func requestToPatient(newPatient NewPatient) Patient {
	var patient Patient
	patient.FirstName = newPatient.FirstName
	patient.LastName = newPatient.LastName
	patient.Address = newPatient.Address
	patient.DNI = newPatient.DNI
	patient.DischargeDate = newPatient.DischargeDate

	return patient
}
