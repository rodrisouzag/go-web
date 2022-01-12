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

func TestServiceUpdate(t *testing.T) {
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
	service := NewService(repo)

	result, err := service.Update(context.Background(), 1, "Rodri", "Souza", "rodri.souza@mercadolibre.com", 24, 1.68, true, "2022-01-12")

	expected := domain.User{Id: 1, Nombre: "Rodri", Apellido: "Souza", Email: "rodri.souza@mercadolibre.com", Edad: 24, Altura: 1.68, Activo: true, FechaDeCreacion: "2022-01-12"}
	assert.Equal(t, expected, result, "deben ser iguales")
	assert.Equal(t, true, db.Mock.ReadUsed, "no se ejecuto el read")
	assert.Nil(t, err, "error al actualizar")
}

func TestServiceDelete(t *testing.T) {
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
	service := NewService(repo)

	expectedError := errors.New("usuario no encontrado")
	err := service.Delete(context.Background(), 2)

	err2 := service.Delete(context.Background(), 1)
	users, _ := service.GetAll(context.Background())

	assert.Equal(t, expectedError, err)
	assert.Equal(t, []domain.User{}, users, "deben ser iguales")
	assert.Nil(t, err2, "error al eliminar")
}
