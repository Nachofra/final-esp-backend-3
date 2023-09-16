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
	Create(ctx context.Context, appointment Appointment) (Appointment, error)
	GetAll(ctx context.Context) []Appointment
	GetByID(ctx context.Context, ID int) (Appointment, error)
	Update(ctx context.Context, appointment Appointment) (Appointment, error)
	Delete(ctx context.Context, ID int) error
}
