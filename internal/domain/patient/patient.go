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
	GetByDNI(ctx context.Context, dni int) (Patient, error)
	Update(ctx context.Context, patient Patient) (Patient, error)
	Delete(ctx context.Context, id int) error
}

type Service struct {
	store Store
}

// NewService creates a new service.
func NewService(store Store) *Service {
	return &Service{
		store: store,
	}
}

// Create creates a new patient.
func (s *Service) Create(ctx context.Context, newPatient NewPatient) (Patient, error) {
	patient := requestToPatient(newPatient)
	response, err := s.store.Create(ctx, patient)
	if err != nil {
		log.Println("error creating all patients on service layer")
		return Patient{}, errors.New("service error. Method Create")
	}

	return response, nil
}

// GetAll returns all patients.
func (s *Service) GetAll(ctx context.Context) ([]Patient, error) {
	patients, err := s.store.GetAll(ctx)
	if err != nil {
		log.Println("error gettin all patients on service layer", err.Error())
		return []Patient{}, errors.New("service error. Method GetAll")
	}
	return patients, nil
}

// GetByID returns a patient by its ID.
func (s *Service) GetByID(ctx context.Context, id int) (Patient, error) {
	patient, err := s.store.GetByID(ctx, id)
	if err != nil {
		log.Println("error getting patient on service layer", err.Error())
		return Patient{}, errors.New("service error. Method GetByID")
	}

	return patient, nil
}

// GetByID returns a patient by its DNI.
func (s *Service) GetByDNI(ctx context.Context, dni int) (Patient, error) {
	patient, err := s.store.GetByDNI(ctx, dni)
	if err != nil {
		log.Println("error getting patient on service layer", err.Error())
		return Patient{}, errors.New("service error. Method GetByDNI")
	}

	return patient, nil
}

// Update updates a patient.
func (s *Service) Update(ctx context.Context, newPatient NewPatient, id int) (Patient, error) {
	patient := requestToPatient(newPatient)
	patient.ID = id
	response, err := s.store.Update(ctx, patient)
	if err != nil {
		log.Println("error updating patient on service layer", err.Error())
		return Patient{}, errors.New("service error. Method Update")
	}

	return response, nil
}

// Patch patches a patient.
func (s *Service) Patch(ctx context.Context, patient Patient, np NewPatient) (Patient, error) {
	if np.FirstName != "" {
		patient.FirstName = np.FirstName
	}

	if np.LastName != "" {
		patient.LastName = np.LastName
	}

	if np.Address != "" {
		patient.Address = np.Address
	}

	if np.DNI != 0 {
		patient.DNI = np.DNI
	}

	if !np.DischargeDate.IsZero() {
		patient.DischargeDate = np.DischargeDate
	}

	p, err := s.store.Update(ctx, patient)
	if err != nil {
		return Patient{}, err
	}

	return p, nil
}

// Delete deletes a patient.
func (s *Service) Delete(ctx context.Context, id int) error {
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
