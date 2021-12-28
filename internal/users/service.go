package users

import (
	"context"

	"github.com/rodrisouzag/go-web/internal/domain"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.User, error)
	Store(ctx context.Context, nombre string, apellido string, email string, edad int, altura float64, activo bool, fechaDeCreacion string) (domain.User, error)
	GetUser(ctx context.Context, id int) (domain.User, error)
	Update(ctx context.Context, id int, nombre string, apellido string, email string, edad int, altura float64, activo bool, fechaDeCreacion string) (domain.User, error)
	UpdateApellidoYEdad(ctx context.Context, id int, apellido string, edad int) (domain.User, error)
	Delete(ctx context.Context, id int) error
}
type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll(ctx context.Context) ([]domain.User, error) {
	users, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *service) Store(ctx context.Context, nombre string, apellido string, email string, edad int, altura float64, activo bool, fechaDeCreacion string) (domain.User, error) {

	id, err := s.repository.GetId(ctx)

	if err != nil {
		return domain.User{}, nil
	}

	user, err := s.repository.Store(ctx, id, nombre, apellido, email, edad, altura, activo, fechaDeCreacion)

	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (s *service) GetUser(ctx context.Context, id int) (domain.User, error) {
	user, err := s.repository.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (s *service) Update(ctx context.Context, id int, nombre string, apellido string, email string, edad int, altura float64, activo bool, fechaDeCreacion string) (domain.User, error) {
	return s.repository.Update(ctx, id, nombre, apellido, email, edad, altura, activo, fechaDeCreacion)
}

func (s *service) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}

func (s *service) UpdateApellidoYEdad(ctx context.Context, id int, apellido string, edad int) (domain.User, error) {
	return s.repository.UpdateApellidoYEdad(ctx, id, apellido, edad)
}
