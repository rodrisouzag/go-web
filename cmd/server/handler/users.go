package handler

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rodrisouzag/go-web/internal/users"
)

type request struct {
	Id              int     `json:"id"`
	Nombre          string  `json:"nombre"`
	Apellido        string  `json:"apellido"`
	Email           string  `json:"email"`
	Edad            int     `json:"edad"`
	Altura          float64 `json:"altura"`
	Activo          bool    `json:"activo"`
	FechaDeCreacion string  `json:"fechaDeCreacion"`
}

type User struct {
	service users.Service
}

func NewUser(u users.Service) *User {
	return &User{
		service: u,
	}
}

func (cUser *User) Saludar() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}

		nombre := ctx.Param("nombre")

		ctx.JSON(200, gin.H{
			"message": "Hola " + nombre,
		})
	}
}

func (cUser *User) Init() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}

		users, err := cUser.service.Init(ctx)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, users)
	}
}

func (cUser *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}

		users, err := cUser.service.GetAll(ctx)
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, users)
	}
}

func (cUser *User) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if req.Nombre == "" {
			ctx.JSON(400, gin.H{"error": "el nombre del usuario es requerido"})
			return
		}
		if req.Apellido == "" {
			ctx.JSON(400, gin.H{"error": "el apellido del usuario es requerido"})
			return
		}
		if req.Email == "" {
			ctx.JSON(400, gin.H{"error": "el email del usuario es requerido"})
			return
		}
		if req.Edad == 0 {
			ctx.JSON(400, gin.H{"error": "la edad del usuario es requerida"})
			return
		}
		if req.Altura == 0.0 {
			ctx.JSON(400, gin.H{"error": "la altura del usuario es requerida"})
			return
		}
		if req.FechaDeCreacion == "" {
			ctx.JSON(400, gin.H{"error": "la fecha de creacion del usuario es requerida"})
			return
		}

		u, err := cUser.service.Store(ctx, req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo, req.FechaDeCreacion)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, u)
	}
}

func (cUser *User) GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}

		idInt, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		user, err := cUser.service.GetUser(ctx, idInt)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, user)
	}
}

func (cUser *User) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}

		idInt, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if req.Nombre == "" {
			ctx.JSON(400, gin.H{"error": "el nombre del usuario es requerido"})
			return
		}
		if req.Apellido == "" {
			ctx.JSON(400, gin.H{"error": "el apellido del usuario es requerido"})
			return
		}
		if req.Email == "" {
			ctx.JSON(400, gin.H{"error": "el email del usuario es requerido"})
			return
		}
		if req.Edad == 0 {
			ctx.JSON(400, gin.H{"error": "la edad del usuario es requerida"})
			return
		}
		if req.Altura == 0.0 {
			ctx.JSON(400, gin.H{"error": "la altura del usuario es requerida"})
			return
		}
		if req.FechaDeCreacion == "" {
			ctx.JSON(400, gin.H{"error": "la fecha de creacion del usuario es requerida"})
			return
		}

		u, err := cUser.service.Update(ctx, idInt, req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo, req.FechaDeCreacion)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, u)
	}
}

func (cUser *User) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}

		idInt, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		err = cUser.service.Delete(ctx, idInt)

		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"data": fmt.Sprintf("el usuario %d ha sido eliminado", idInt)})
	}
}

func (cUser *User) UpdateApellidoYEdad() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}

		idInt, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if req.Apellido == "" {
			ctx.JSON(400, gin.H{"error": "el apellido del usuario es requerido"})
			return
		}

		if req.Edad == 0 {
			ctx.JSON(400, gin.H{"error": "la edad del usuario es requerida"})
			return
		}

		u, err := cUser.service.UpdateApellidoYEdad(ctx, idInt, req.Apellido, req.Edad)

		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, u)
	}
}
