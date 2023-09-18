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
	GetByRegistrationNumber(ctx context.Context, rn int) (Dentist, error)
	Update(ctx context.Context, dentist Dentist) (Dentist, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	store Store
}

type Service interface {
	Create(ctx context.Context, newDentist NewDentist) (Dentist, error)
	GetAll(ctx context.Context) ([]Dentist, error)
	GetByID(ctx context.Context, id int) (Dentist, error)
	GetByRegistrationNumber(ctx context.Context, rn int) (Dentist, error)
	Update(ctx context.Context, updateDentist UpdateDentist, id int) (Dentist, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, dentist Dentist, pd PatchDentist) (Dentist, error)
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
	dentist := newToDentist(newDentist)
	response, err := s.store.Create(ctx, dentist)
	if err != nil {
		log.Println("error en servicio. Metodo create")
		return Dentist{}, errors.New("error en servicio. Metodo create")
	}

	return response, nil
}

// GetAll returns all products.
func (s *service) GetAll(ctx context.Context) ([]Dentist, error) {
	dentists, err := s.store.GetAll(ctx)
	if err != nil {
		log.Println("error getting dentists on service layer", err.Error())
		return []Dentist{}, errors.New("service error. Method GetAll")
	}

	return dentists, nil
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

// GetByRegistrationNumber returns a patient by its RegistrationNumber.
func (s *service) GetByRegistrationNumber(ctx context.Context, dni int) (Dentist, error) {
	dentist, err := s.store.GetByRegistrationNumber(ctx, dni)
	if err != nil {
		log.Println("error getting patient on service layer", err.Error())
		return Dentist{}, errors.New("service error. Method GetByRegistrationNumber")
	}

	return dentist, nil
}

// Update updates a product.
func (s *service) Update(ctx context.Context, updateDentist UpdateDentist, id int) (Dentist, error) {
	dentist := updateToDentist(updateDentist)
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

// Patch patches an appointment.
func (s *service) Patch(ctx context.Context, dentist Dentist, pd PatchDentist) (Dentist, error) {
	if pd.FirstName != nil {
		dentist.FirstName = *pd.FirstName
	}

	if pd.LastName != nil {
		dentist.LastName = *pd.LastName
	}

	if pd.RegistrationNumber != nil {
		dentist.RegistrationNumber = *pd.RegistrationNumber
	}

	a, err := s.store.Update(ctx, dentist)
	if err != nil {
		return Dentist{}, err
	}

	return a, nil
}

func newToDentist(newDentist NewDentist) Dentist {
	var dentist Dentist
	dentist.FirstName = newDentist.FirstName
	dentist.LastName = newDentist.LastName
	dentist.RegistrationNumber = newDentist.RegistrationNumber

	return dentist
}

func updateToDentist(updateDentist UpdateDentist) Dentist {
	var dentist Dentist
	dentist.FirstName = updateDentist.FirstName
	dentist.LastName = updateDentist.LastName
	dentist.RegistrationNumber = updateDentist.RegistrationNumber

	return dentist
}
