package users

import (
	"context"
	"errors"
	"fmt"

	"github.com/rodrisouzag/go-web/internal/domain"
	"github.com/rodrisouzag/go-web/pkg/store"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.User, error)
	GetId(ctx context.Context) (int, error)
	Store(ctx context.Context, id int, nombre string, apellido string, email string, edad int, altura float64, activo bool, fechaDeCreacion string) (domain.User, error)
	GetUser(ctx context.Context, id int) (domain.User, error)
	Update(ctx context.Context, id int, nombre string, apellido string, email string, edad int, altura float64, activo bool, fechaDeCreacion string) (domain.User, error)
	UpdateApellidoYEdad(ctx context.Context, id int, apellido string, edad int) (domain.User, error)
	Delete(ctx context.Context, id int) error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	err := r.db.Read(&users)
	if err != nil {
		return users, err
	}
	return users, nil
}

func (r *repository) GetId(ctx context.Context) (int, error) {
	var users []domain.User
	if err := r.db.Read(&users); err != nil {
		return 0, err
	}

	if len(users) > 0 {
		return users[len(users)-1].Id + 1, nil
	} else {
		return 1, nil
	}
}

func (r *repository) Store(ctx context.Context, id int, nombre string, apellido string, email string, edad int, altura float64, activo bool, fechaDeCreacion string) (domain.User, error) {
	var users []domain.User
	if err := r.db.Read(&users); err != nil {
		return domain.User{}, err
	}
	u := domain.User{Id: id, Nombre: nombre, Apellido: apellido, Email: email, Edad: edad, Altura: altura, Activo: activo, FechaDeCreacion: fechaDeCreacion}
	users = append(users, u)
	if err := r.db.Write(users); err != nil {
		return domain.User{}, err
	}
	return u, nil
}

func (r *repository) GetUser(ctx context.Context, id int) (domain.User, error) {
	var users []domain.User
	if err := r.db.Read(&users); err != nil {
		return domain.User{}, err
	}
	for _, u := range users {
		if u.Id == id {
			return u, nil
		}
	}
	return domain.User{}, fmt.Errorf("usuario %d no encontrado", id)
}

func (r *repository) Update(ctx context.Context, id int, nombre string, apellido string, email string, edad int, altura float64, activo bool, fechaDeCreacion string) (domain.User, error) {
	var users []domain.User
	if err := r.db.Read(&users); err != nil {
		return domain.User{}, err
	}
	u := domain.User{Id: id, Nombre: nombre, Apellido: apellido, Email: email, Edad: edad, Altura: altura, Activo: activo, FechaDeCreacion: fechaDeCreacion}
	updated := false
	for i := range users {
		if users[i].Id == id {
			users[i] = u
			updated = true
		}
	}
	if !updated {
		return domain.User{}, fmt.Errorf("usuario %d no encontrado", id)
	}
	if err := r.db.Write(users); err != nil {
		return domain.User{}, err
	}
	return u, nil
}

func (r *repository) UpdateApellidoYEdad(ctx context.Context, id int, apellido string, edad int) (domain.User, error) {
	var users []domain.User
	if err := r.db.Read(&users); err != nil {
		return domain.User{}, err
	}
	updated := false
	var u domain.User
	for i := range users {
		if users[i].Id == id {
			users[i].Apellido = apellido
			users[i].Edad = edad
			u = users[i]
			updated = true
		}
	}
	if !updated {
		return domain.User{}, fmt.Errorf("usuario %d no encontrado", id)
	}
	if err := r.db.Write(users); err != nil {
		return domain.User{}, err
	}
	return u, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	var users []domain.User
	if err := r.db.Read(&users); err != nil {
		return err
	}
	deleted := false
	var index int
	for i := range users {
		if users[i].Id == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return errors.New("usuario no encontrado")
	}
	users = append(users[:index], users[index+1:]...)

	if err := r.db.Write(users); err != nil {
		return err
	}
	return nil
}
