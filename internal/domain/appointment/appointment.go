package appointment

import (
	"context"
	"errors"
)

var (
	ErrNotFound  = errors.New("appointment not found")
	ErrStatement = errors.New("error preparing statement")
	ErrExec      = errors.New("error executing statement")
	ErrLastId    = errors.New("error getting last id")
)

type Store interface {
	GetAll(ctx context.Context, filters map[string]string) []Appointment
	GetByID(ctx context.Context, ID int) (Appointment, error)
	Create(ctx context.Context, appointment Appointment) (Appointment, error)
	Update(ctx context.Context, appointment Appointment) (Appointment, error)
	Delete(ctx context.Context, ID int) error
}

type Service struct {
	store Store
}

func NewStore(store Store) *Service {
	return &Service{
		store: store,
	}
}

// GetAll returns appointments by filter.
func (s *Service) GetAll(ctx context.Context, filters map[string]string) []Appointment {
	appointments := s.store.GetAll(ctx, filters)
	return appointments
}

// GetByID returns an appointment by its ID.
func (s *Service) GetByID(ctx context.Context, ID int) (Appointment, error) {
	appointment, err := s.store.GetByID(ctx, ID)
	if err != nil {
		return Appointment{}, err
	}

	return appointment, nil
}

// Create creates a new appointment.
func (s *Service) Create(ctx context.Context, appointment Appointment) (Appointment, error) {
	a, err := s.store.Create(ctx, appointment)
	if err != nil {
		return Appointment{}, err
	}

	return a, nil
}

// Update updates an appointment.
func (s *Service) Update(ctx context.Context, appointment Appointment, ID int) (Appointment, error) {
	appointment.ID = ID
	a, err := s.store.Update(ctx, appointment)
	if err != nil {
		return Appointment{}, err
	}

	return a, nil
}

// Delete deletes an appointment.
func (s *Service) Delete(ctx context.Context, ID int) error {
	err := s.store.Delete(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}
