package users

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/rodrisouzag/go-web/internal/domain"
)

var users []domain.User

type Repository interface {
	Init(ctx context.Context) ([]domain.User, error)
	GetAll(ctx context.Context) ([]domain.User, error)
	GetId(ctx context.Context) (int, error)
	Store(ctx context.Context, id int, nombre string, apellido string, email string, edad int, altura float64, activo bool, fechaDeCreacion string) (domain.User, error)
	GetUser(ctx context.Context, id int) (domain.User, error)
	Update(ctx context.Context, id int, nombre string, apellido string, email string, edad int, altura float64, activo bool, fechaDeCreacion string) (domain.User, error)
	UpdateApellidoYEdad(ctx context.Context, id int, apellido string, edad int) (domain.User, error)
	Delete(ctx context.Context, id int) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Init(ctx context.Context) ([]domain.User, error) {
	datosJson, err := ioutil.ReadFile("internal/users/users.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(datosJson, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *repository) GetAll(ctx context.Context) ([]domain.User, error) {
	return users, nil
}

func (r *repository) GetId(ctx context.Context) (int, error) {
	if len(users) > 0 {
		return users[len(users)-1].Id + 1, nil
	} else {
		return 1, nil
	}
}

func (r *repository) Store(ctx context.Context, id int, nombre string, apellido string, email string, edad int, altura float64, activo bool, fechaDeCreacion string) (domain.User, error) {
	u := domain.User{id, nombre, apellido, email, edad, altura, activo, fechaDeCreacion}
	users = append(users, u)
	return u, nil
}

func (r *repository) GetUser(ctx context.Context, id int) (domain.User, error) {
	for _, u := range users {
		if u.Id == id {
			return u, nil
		}
	}
	return domain.User{}, fmt.Errorf("usuario %d no encontrado", id)
}

func (r *repository) Update(ctx context.Context, id int, nombre string, apellido string, email string, edad int, altura float64, activo bool, fechaDeCreacion string) (domain.User, error) {
	u := domain.User{id, nombre, apellido, email, edad, altura, activo, fechaDeCreacion}
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
	return u, nil
}

func (r *repository) UpdateApellidoYEdad(ctx context.Context, id int, apellido string, edad int) (domain.User, error) {
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
	return u, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	deleted := false
	var index int
	for i := range users {
		if users[i].Id == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("usuario %d no encontrado", id)
	}
	users = append(users[:index], users[index+1:]...)
	return nil
}
