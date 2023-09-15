package odontologo

import (
	"context"
	"errors"
)

// Errores
var (
	ErrEmptyList = errors.New("the list is empty")
	ErrNotFound  = errors.New("odontologo not found")
	ErrStatement = errors.New("error preparing statement")
	ErrExec      = errors.New("error exect statement")
	ErrLastId    = errors.New("error getting last id")
)

type Repository interface {
	Create(ctx context.Context, odontologo Odontologo) (Odontologo, error)
	GetAll(ctx context.Context) ([]Odontologo, error)
	GetByID(ctx context.Context, id int) (Odontologo, error)
	Update(ctx context.Context, odontologo Odontologo) (Odontologo, error)
	Delete(ctx context.Context, id int) error
}
