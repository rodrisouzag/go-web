package users

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/rodrisouzag/go-web/internal/domain"
	"github.com/stretchr/testify/assert"
)

type stubStore struct {
}

type mockStore struct {
	readUsed bool
}

func (s *stubStore) Read(data interface{}) error {
	u1 := domain.User{1, "Rodrigo", "Souza", "rodrigo.souza@mercadolibre.com", 23, 1.68, true, "2021-12-13"}
	u2 := domain.User{2, "Juan", "Perez", "juan.perez@mercadolibre.com", 30, 1.75, false, "2010-10-10"}
	users := []domain.User{u1, u2}
	fileData, _ := json.MarshalIndent(users, "", "  ")
	json.Unmarshal(fileData, &data)
	return nil
}

func (s *stubStore) Write(data interface{}) error {
	return nil
}

func (s *mockStore) Read(data interface{}) error {
	u1 := domain.User{1, "Rodrigo", "Before", "rodrigo.souza@mercadolibre.com", 23, 1.68, true, "2021-12-13"}
	u2 := domain.User{2, "Juan", "Perez", "juan.perez@mercadolibre.com", 30, 1.75, false, "2010-10-10"}

	users := []domain.User{u1, u2}
	fileData, _ := json.MarshalIndent(users, "", "  ")
	json.Unmarshal(fileData, &data)
	s.readUsed = true
	return nil

}

func (s *mockStore) Write(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {
	db := stubStore{}
	repo := NewRepository(&db)

	var result []domain.User
	result, _ = repo.GetAll(context.Background())

	u1 := domain.User{1, "Rodrigo", "Souza", "rodrigo.souza@mercadolibre.com", 23, 1.68, true, "2021-12-13"}
	u2 := domain.User{2, "Juan", "Perez", "juan.perez@mercadolibre.com", 30, 1.75, false, "2010-10-10"}

	expected := []domain.User{u1, u2}

	assert.Equal(t, expected, result, "deben ser iguales")
}

func TestUpdateApellidoYEdad(t *testing.T) {
	db := mockStore{}
	repo := NewRepository(&db)

	result, err := repo.UpdateApellidoYEdad(context.Background(), 1, "After", 23)

	expected := domain.User{1, "Rodrigo", "After", "rodrigo.souza@mercadolibre.com", 23, 1.68, true, "2021-12-13"}
	assert.Equal(t, expected, result, "deben ser iguales")
	assert.Equal(t, true, db.readUsed, "no se ejecuto el read")
	assert.Nil(t, err, "error al actualizar")
}
