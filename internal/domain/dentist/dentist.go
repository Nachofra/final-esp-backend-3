package dentist

import (
	"context"
	"errors"
)

var (
	ErrNotFound      = errors.New("dentist not found")
	ErrConflict      = errors.New("constraint conflict while storing")
	ErrAlreadyExists = errors.New("dentist already exists")
	ErrValueExceeded = errors.New("attribute value exceed type limit")
)

// Store specifies the contract needed for the Store in the Service.
type Store interface {
	Create(ctx context.Context, dentist Dentist) (Dentist, error)
	GetAll(ctx context.Context) []Dentist
	GetByID(ctx context.Context, id int) (Dentist, error)
	GetByRegistrationNumber(ctx context.Context, rn int) (Dentist, error)
	Update(ctx context.Context, dentist Dentist) (Dentist, error)
	Delete(ctx context.Context, id int) error
}

// service unifies all the business operation for the domain.
type service struct {
	store Store
}

// Service specifies the contract needed for the Service.
type Service interface {
	Create(ctx context.Context, newDentist NewDentist) (Dentist, error)
	GetAll(ctx context.Context) []Dentist
	GetByID(ctx context.Context, id int) (Dentist, error)
	GetByRegistrationNumber(ctx context.Context, rn int) (Dentist, error)
	Update(ctx context.Context, updateDentist UpdateDentist, id int) (Dentist, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, dentist Dentist, nd NewDentist) (Dentist, error)
}

// NewService creates a new product service.
func NewService(store Store) Service {
	return &service{
		store: store,
	}
}

// Create creates a new product.
func (s *service) Create(ctx context.Context, newDentist NewDentist) (Dentist, error) {
	dentist := newToDentist(newDentist)
	response, err := s.store.Create(ctx, dentist)
	if err != nil {
		return Dentist{}, err
	}

	return response, nil
}

// GetAll returns all products.
func (s *service) GetAll(ctx context.Context) []Dentist {
	dentists := s.store.GetAll(ctx)
	return dentists
}

// GetByID returns a product by its ID.
func (s *service) GetByID(ctx context.Context, id int) (Dentist, error) {
	dentist, err := s.store.GetByID(ctx, id)
	if err != nil {
		return Dentist{}, err
	}

	return dentist, nil
}

// GetByRegistrationNumber returns a patient by its RegistrationNumber.
func (s *service) GetByRegistrationNumber(ctx context.Context, dni int) (Dentist, error) {
	dentist, err := s.store.GetByRegistrationNumber(ctx, dni)
	if err != nil {
		return Dentist{}, err
	}

	return dentist, nil
}

// Update updates a product.
func (s *service) Update(ctx context.Context, updateDentist UpdateDentist, id int) (Dentist, error) {
	dentist := updateToDentist(updateDentist)
	dentist.ID = id
	response, err := s.store.Update(ctx, dentist)
	if err != nil {
		return Dentist{}, err
	}

	return response, nil
}

// Delete deletes a product.
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.store.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// Patch patches an appointment.
func (s *service) Patch(ctx context.Context, dentist Dentist, nd NewDentist) (Dentist, error) {
	if nd.FirstName != "" {
		dentist.FirstName = nd.FirstName
	}

	if nd.LastName != "" {
		dentist.LastName = nd.LastName
	}

	if nd.RegistrationNumber != 0 {
		dentist.RegistrationNumber = nd.RegistrationNumber
	}

	d, err := s.store.Update(ctx, dentist)
	if err != nil {
		return Dentist{}, err
	}

	return d, nil
}

// newToDentist parses NewDentist to Dentist
func newToDentist(newDentist NewDentist) Dentist {
	var dentist Dentist
	dentist.FirstName = newDentist.FirstName
	dentist.LastName = newDentist.LastName
	dentist.RegistrationNumber = newDentist.RegistrationNumber

	return dentist
}

// updateToDentist parses UpdateDentist to Dentist
func updateToDentist(updateDentist UpdateDentist) Dentist {
	var dentist Dentist
	dentist.FirstName = updateDentist.FirstName
	dentist.LastName = updateDentist.LastName
	dentist.RegistrationNumber = updateDentist.RegistrationNumber

	return dentist
}
