package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrisouzag/go-web/cmd/server/handler"
	"github.com/rodrisouzag/go-web/internal/users"
)

/*

// Ejercicio 1 C1 TT

func SearchUsers(ctx *gin.Context) {

	var res []User

	id := ctx.Query("id")
	nombre := ctx.Query("nombre")
	apellido := ctx.Query("apellido")
	email := ctx.Query("email")
	edad := ctx.Query("edad")
	altura := ctx.Query("altura")
	activo := ctx.Query("activo")
	fechaDeCreacion := ctx.Query("fechaDeCreacion")

	if id != "" {
		idInt, err := strconv.Atoi(id)
		if err == nil {
			ctx.JSON(200, usersMap[idInt])
		}
	} else if nombre != "" {
		for _, u := range usersMap {
			if u.Nombre == nombre {
				res = append(res, u)
			}
		}
	} else if apellido != "" {
		for _, u := range usersMap {
			if u.Apellido == apellido {
				res = append(res, u)
			}
		}
	} else if email != "" {
		for _, u := range usersMap {
			if u.Email == email {
				res = append(res, u)
			}
		}
	} else if edad != "" {
		for _, u := range usersMap {
			if fmt.Sprint(u.Edad) == edad {
				res = append(res, u)
			}
		}
	} else if altura != "" {
		for _, u := range usersMap {
			if fmt.Sprint(u.Altura) == altura {
				res = append(res, u)
			}
		}
	} else if activo != "" {
		for _, u := range usersMap {
			if fmt.Sprint(u.Activo) == activo {
				res = append(res, u)
			}
		}
	} else if fechaDeCreacion != "" {
		for _, u := range usersMap {
			if u.FechaDeCreacion == fechaDeCreacion {
				res = append(res, u)
			}
		}
	}

	if res == nil {
		ctx.String(404, "No existe un usuario con esas caracterisitcas!")
	} else {
		ctx.JSON(200, res)
	}
}

*/

func main() {
	repo := users.NewRepository()
	service := users.NewService(repo)
	u := handler.NewUser(service)
	router := gin.Default()

	router.GET("saludar/:nombre", u.Saludar())

	usr := router.Group("/users")
	usr.POST("/", u.Store())
	usr.GET("/", u.GetAll())
	usr.GET("/init", u.Init())
	usr.GET("/user/:id", u.GetUser())

	router.Run()
}
