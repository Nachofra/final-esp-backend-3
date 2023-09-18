package dentist

import (
	"context"
	"github.com/pkg/errors"
)

var (
	ErrNotFound  = errors.New("odontologo not found")
	ErrStatement = errors.New("error preparing statement")
	ErrExec      = errors.New("error exect statement")
	ErrLastId    = errors.New("error getting last id")
)

type Store interface {
	Create(ctx context.Context, dentist Dentist) (Dentist, error)
	GetAll(ctx context.Context) ([]Dentist, error)
	GetByID(ctx context.Context, id int) (Dentist, error)
	Update(ctx context.Context, dentist Dentist) (Dentist, error)
	Delete(ctx context.Context, id int) error
}
