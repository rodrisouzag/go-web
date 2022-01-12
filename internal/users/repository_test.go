package users

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/rodrisouzag/go-web/internal/domain"
	"github.com/rodrisouzag/go-web/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	input := []domain.User{
		{
			Id:              1,
			Nombre:          "Rodrigo",
			Apellido:        "Souza",
			Email:           "rodrigo.souza@mercadolibre.com",
			Edad:            23,
			Altura:          1.68,
			Activo:          true,
			FechaDeCreacion: "2021-12-13",
		}, {
			Id:              2,
			Nombre:          "Juan",
			Apellido:        "Perez",
			Email:           "juan.perez@mercadolibre.com",
			Edad:            30,
			Altura:          1.75,
			Activo:          false,
			FechaDeCreacion: "2010-10-10",
		},
	}

	dataJson, _ := json.MarshalIndent(input, "", "  ")
	dbMock := store.Mock{
		Data: dataJson,
	}
	stubStore := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}

	repo := NewRepository(&stubStore)
	var result []domain.User
	result, _ = repo.GetAll(context.Background())

	u1 := domain.User{Id: 1, Nombre: "Rodrigo", Apellido: "Souza", Email: "rodrigo.souza@mercadolibre.com", Edad: 23, Altura: 1.68, Activo: true, FechaDeCreacion: "2021-12-13"}
	u2 := domain.User{Id: 2, Nombre: "Juan", Apellido: "Perez", Email: "juan.perez@mercadolibre.com", Edad: 30, Altura: 1.75, Activo: false, FechaDeCreacion: "2010-10-10"}

	expected := []domain.User{u1, u2}

	assert.Equal(t, expected, result, "deben ser iguales")
}

func TestGetUser(t *testing.T) {
	input := []domain.User{
		{
			Id:              1,
			Nombre:          "Rodrigo",
			Apellido:        "Souza",
			Email:           "rodrigo.souza@mercadolibre.com",
			Edad:            23,
			Altura:          1.68,
			Activo:          true,
			FechaDeCreacion: "2021-12-13",
		}, {
			Id:              2,
			Nombre:          "Juan",
			Apellido:        "Perez",
			Email:           "juan.perez@mercadolibre.com",
			Edad:            30,
			Altura:          1.75,
			Activo:          false,
			FechaDeCreacion: "2010-10-10",
		},
	}

	dataJson, _ := json.MarshalIndent(input, "", "  ")
	dbMock := store.Mock{
		Data: dataJson,
	}
	stubStore := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}

	repo := NewRepository(&stubStore)

	result, _ := repo.GetUser(context.Background(), 1)

	expected := domain.User{Id: 1, Nombre: "Rodrigo", Apellido: "Souza", Email: "rodrigo.souza@mercadolibre.com", Edad: 23, Altura: 1.68, Activo: true, FechaDeCreacion: "2021-12-13"}

	assert.Equal(t, expected, result, "deben ser iguales")
}

func TestUpdateApellidoYEdad(t *testing.T) {
	input := []domain.User{
		{
			Id:              1,
			Nombre:          "Rodrigo",
			Apellido:        "Before",
			Email:           "rodrigo.souza@mercadolibre.com",
			Edad:            23,
			Altura:          1.68,
			Activo:          true,
			FechaDeCreacion: "2021-12-13",
		}, {
			Id:              2,
			Nombre:          "Juan",
			Apellido:        "Perez",
			Email:           "juan.perez@mercadolibre.com",
			Edad:            30,
			Altura:          1.75,
			Activo:          false,
			FechaDeCreacion: "2010-10-10",
		},
	}

	dataJson, _ := json.MarshalIndent(input, "", "  ")
	dbMock := store.Mock{
		Data: dataJson,
	}
	db := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}

	repo := NewRepository(&db)

	result, err := repo.UpdateApellidoYEdad(context.Background(), 1, "After", 23)

	expected := domain.User{Id: 1, Nombre: "Rodrigo", Apellido: "After", Email: "rodrigo.souza@mercadolibre.com", Edad: 23, Altura: 1.68, Activo: true, FechaDeCreacion: "2021-12-13"}
	assert.Equal(t, expected, result, "deben ser iguales")
	assert.Equal(t, true, db.Mock.ReadUsed, "no se ejecuto el read")
	assert.Nil(t, err, "error al actualizar apellido")
}

func TestUpdate(t *testing.T) {
	input := []domain.User{
		{
			Id:              1,
			Nombre:          "Rodrigo",
			Apellido:        "Souza",
			Email:           "rodrigo.souza@mercadolibre.com",
			Edad:            23,
			Altura:          1.68,
			Activo:          true,
			FechaDeCreacion: "2021-12-13",
		}, {
			Id:              2,
			Nombre:          "Juan",
			Apellido:        "Perez",
			Email:           "juan.perez@mercadolibre.com",
			Edad:            30,
			Altura:          1.75,
			Activo:          false,
			FechaDeCreacion: "2010-10-10",
		},
	}

	dataJson, _ := json.MarshalIndent(input, "", "  ")
	dbMock := store.Mock{
		Data: dataJson,
	}
	db := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}

	repo := NewRepository(&db)

	result, err := repo.Update(context.Background(), 1, "Rodri", "Souza", "rodri.souza@mercadolibre.com", 24, 1.68, true, "2022-01-12")

	expected := domain.User{Id: 1, Nombre: "Rodri", Apellido: "Souza", Email: "rodri.souza@mercadolibre.com", Edad: 24, Altura: 1.68, Activo: true, FechaDeCreacion: "2022-01-12"}
	assert.Equal(t, expected, result, "deben ser iguales")
	assert.Equal(t, true, db.Mock.ReadUsed, "no se ejecuto el read")
	assert.Nil(t, err, "error al actualizar")
}

func TestDelete(t *testing.T) {
	input := []domain.User{
		{
			Id:              1,
			Nombre:          "Rodrigo",
			Apellido:        "Souza",
			Email:           "rodrigo.souza@mercadolibre.com",
			Edad:            23,
			Altura:          1.68,
			Activo:          true,
			FechaDeCreacion: "2021-12-13",
		},
	}

	dataJson, _ := json.MarshalIndent(input, "", "  ")
	dbMock := store.Mock{
		Data: dataJson,
	}
	db := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}

	repo := NewRepository(&db)

	expectedError := errors.New("usuario no encontrado")
	err := repo.Delete(context.Background(), 2)

	err2 := repo.Delete(context.Background(), 1)

	users, _ := repo.GetAll(context.Background())

	assert.Equal(t, expectedError, err)
	assert.Equal(t, []domain.User{}, users, "deben ser iguales")
	assert.Nil(t, err2, "error al eliminar")
}
