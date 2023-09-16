package patient

import (
	"context"
	"errors"
)

var (
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
