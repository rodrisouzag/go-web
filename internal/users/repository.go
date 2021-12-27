package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type User struct {
	Id              int     `json:"id"`
	Nombre          string  `json:"nombre" binding:"required"`
	Apellido        string  `json:"apellido" binding:"required"`
	Email           string  `json:"email" binding:"required"`
	Edad            int     `json:"edad" binding:"required"`
	Altura          float64 `json:"altura" binding:"required"`
	Activo          bool    `json:"activo" binding:"required"`
	FechaDeCreacion string  `json:"fechaDeCreacion" binding:"required"`
}

var users []User

type Repository interface {
	Init() ([]User, error)
	GetAll() ([]User, error)
	GetId() (int, error)
	Store(id int, nombre string, apellido string, email string, edad int, altura float64, activo bool, fechaDeCreacion string) (User, error)
	GetUser(id int) (User, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Init() ([]User, error) {
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

func (r *repository) GetAll() ([]User, error) {
	return users, nil
}

func (r *repository) GetId() (int, error) {
	if len(users) > 0 {
		return users[len(users)-1].Id + 1, nil
	} else {
		return 1, nil
	}
}

func (r *repository) Store(id int, nombre string, apellido string, email string, edad int, altura float64, activo bool, fechaDeCreacion string) (User, error) {
	u := User{id, nombre, apellido, email, edad, altura, activo, fechaDeCreacion}
	users = append(users, u)
	return u, nil
}

func (r *repository) GetUser(id int) (User, error) {
	for _, u := range users {
		if u.Id == id {
			return u, nil
		}
	}
	return User{}, fmt.Errorf("Usuario no existe")
}
