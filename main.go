package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id              int     `json:"id"`
	Nombre          string  `json:"nombre"`
	Apellido        string  `json:"apellido"`
	Email           string  `json:"email"`
	Edad            int     `json:"edad"`
	Altura          float64 `json:"altura"`
	Activo          bool    `json:"activo"`
	FechaDeCreacion string  `json:"fechaDeCreacion"`
}

var users []User

func GetAll(ctx *gin.Context) {
	ctx.JSON(200, users)
}

func main() {
	router := gin.Default()

	// Ejercicio 1

	router.GET("saludar/:nombre", func(ctx *gin.Context) {
		nombre := ctx.Param("nombre")

		ctx.JSON(200, gin.H{
			"message": "Hola " + nombre,
		})
	})

	//Ejercicio 2

	datosJson, err := ioutil.ReadFile("users.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(datosJson, &users)
	if err != nil {
		log.Fatal(err)
	}

	router.GET("/users", GetAll)

	router.Run()

}
