package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rodrisouzag/go-web/internal/domain"
	"github.com/rodrisouzag/go-web/internal/users"
	"github.com/rodrisouzag/go-web/pkg/store"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")
	db := store.New(store.FileType, "users_test.json")
	repo := users.NewRepository(db)
	service := users.NewService(repo)
	u := NewUser(service)
	r := gin.Default()

	usr := r.Group("/users")
	usr.GET("/", u.GetAll())
	usr.PUT("/:id", u.Update())
	usr.DELETE("/:id", u.Delete())
	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func Test_GetUsers_OK(t *testing.T) {
	// crear el Server y definir las Rutas
	r := createServer()
	// crear Request del tipo GET y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodGet, "/users/", "")

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	objRes := struct {
		Code string        `json:"code"`
		Data []domain.User `json:"data"`
	}{}

	err := json.Unmarshal(rr.Body.Bytes(), &objRes)

	assert.Equal(t, 200, rr.Code)
	assert.Nil(t, err)
	assert.True(t, len(objRes.Data) > 0)
}

func Test_UpdateUser_OK(t *testing.T) {
	// crear el Server y definir las Rutas
	r := createServer()
	// crear Request del tipo PUT y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodPut, "/users/3",
		`{
        "nombre": "Ana",
        "apellido": "Lopez",
        "email": "ana.lopez@mercadolibre.com",
        "edad": 35,
        "altura": 1.56,
        "activo": true,
        "fechaDeCreacion": "2021-12-25"
    }`)

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	objRes := struct {
		Code string      `json:"code"`
		Data domain.User `json:"data"`
	}{}

	expected := domain.User{
		Id:              3,
		Nombre:          "Ana",
		Apellido:        "Lopez",
		Email:           "ana.lopez@mercadolibre.com",
		Edad:            35,
		Altura:          1.56,
		Activo:          true,
		FechaDeCreacion: "2021-12-25",
	}

	err := json.Unmarshal(rr.Body.Bytes(), &objRes)

	assert.Equal(t, 200, rr.Code)
	assert.Nil(t, err)
	assert.Equal(t, expected, objRes.Data)
}

func Test_DeleteUser_OK(t *testing.T) {
	// crear el Server y definir las Rutas
	r := createServer()
	// crear Request del tipo PUT y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodDelete, "/users/3", "")

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}
