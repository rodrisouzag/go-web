package handler

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rodrisouzag/go-web/internal/users"
)

type request struct {
	Id              int     `json:"id"`
	Nombre          string  `json:"nombre" binding:"required"`
	Apellido        string  `json:"apellido" binding:"required"`
	Email           string  `json:"email" binding:"required"`
	Edad            int     `json:"edad" binding:"required"`
	Altura          float64 `json:"altura" binding:"required"`
	Activo          bool    `json:"activo" binding:"required"`
	FechaDeCreacion string  `json:"fechaDeCreacion" binding:"required"`
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

		users, err := cUser.service.Init()
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

		users, err := cUser.service.GetAll()
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
			v := reflect.ValueOf(req)
			tipoObtenidoDeReflection := v.Type()
			message := ""
			for i := 0; i < v.NumField(); i++ {
				if tipoObtenidoDeReflection.Field(i).Name != "Id" && v.Field(i).Interface() == reflect.Zero(v.Field(i).Type()).Interface() {
					message += fmt.Sprintf("el campo %s es requerido\n", tipoObtenidoDeReflection.Field(i).Name)
				}
			}
			ctx.String(404, message)
			return
		}
		u, err := cUser.service.Store(req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo, req.FechaDeCreacion)
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
		user, err := cUser.service.GetUser(idInt)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, user)
	}
}
