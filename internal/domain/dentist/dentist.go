package dentist

import (
	"context"
	"errors"
	"log"
)

var (
	ErrNotFound  = errors.New("dentist not found")
	ErrStatement = errors.New("error preparing statement")
	ErrExec      = errors.New("error executing statement")
	ErrLastId    = errors.New("error getting last id")
)

type Store interface {
	Create(ctx context.Context, dentist Dentist) (Dentist, error)
	GetAll(ctx context.Context) []Dentist
	GetByID(ctx context.Context, id int) (Dentist, error)
	Update(ctx context.Context, dentist Dentist) (Dentist, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	store Store
}

type Service interface {
	Create(ctx context.Context, newDentist NewDentist) (Dentist, error)
	GetAll(ctx context.Context) []Dentist
	GetByID(ctx context.Context, id int) (Dentist, error)
	Update(ctx context.Context, newDentist NewDentist, id int) (Dentist, error)
	Delete(ctx context.Context, id int) error
}

// NewService creates a new product service.
func NewService(store Store) Service {
	return &service{
		//Here we put the repository
		store: store,
	}
}

// Create creates a new product.
func (s *service) Create(ctx context.Context, newDentist NewDentist) (Dentist, error) {
	dentist := requestToDentist(newDentist)
	response, err := s.store.Create(ctx, dentist)
	if err != nil {
		log.Println("error en servicio. Metodo create")
		return Dentist{}, errors.New("error en servicio. Metodo create")
	}

	return response, nil
}

// GetAll returns all products.
func (s *service) GetAll(ctx context.Context) []Dentist {
	productos := s.store.GetAll(ctx)
	return productos
}

// GetByID returns a product by its ID.
func (s *service) GetByID(ctx context.Context, id int) (Dentist, error) {
	dentist, err := s.store.GetByID(ctx, id)
	if err != nil {
		log.Println("error getting dentist on service layer", err.Error())
		return Dentist{}, errors.New("service error. Method GetByID")
	}

	return dentist, nil
}

// Update updates a product.
func (s *service) Update(ctx context.Context, newDentist NewDentist, id int) (Dentist, error) {
	dentist := requestToDentist(newDentist)
	dentist.ID = id
	response, err := s.store.Update(ctx, dentist)
	if err != nil {
		log.Println("error updating dentist on service layer", err.Error())
		return Dentist{}, errors.New("service error. Method Update")
	}

	return response, nil
}

// Delete deletes a product.
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.store.Delete(ctx, id)
	if err != nil {
		log.Println("error deleting dentist on service layer", err.Error())
		return errors.New("service error. Method Delete")
	}

	return nil
}

func requestToDentist(newDentist NewDentist) Dentist {
	var dentist Dentist
	dentist.FirstName = newDentist.FirstName
	dentist.LastName = newDentist.LastName
	dentist.RegistrationNumber = newDentist.RegistrationNumber

	return dentist
}
