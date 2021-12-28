package handler

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rodrisouzag/go-web/internal/users"
	"github.com/rodrisouzag/go-web/pkg/web"
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
		nombre := ctx.Param("nombre")

		ctx.JSON(200, gin.H{
			"message": "Hola " + nombre,
		})
	}
}

// ListUsers godoc
// @Summary List users
// @Tags Users
// @Description get users
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /users [get]
func (cUser *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users, err := cUser.service.GetAll(ctx)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, users, ""))
	}
}

// StoreUsers godoc
// @Summary Store users
// @Tags Users
// @Description store users
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param product body request true "User to store"
// @Success 200 {object} web.Response
// @Router /users [post]
func (cUser *User) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		if req.Nombre == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "el nombre del usuario es requerido"))
			return
		}
		if req.Apellido == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "el apellido del usuario es requerido"))
			return
		}
		if req.Email == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "el email del usuario es requerido"))
			return
		}
		if req.Edad == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "la edad del usuario es requerida"))
			return
		}
		if req.Altura == 0.0 {
			ctx.JSON(400, web.NewResponse(400, nil, "la altura del usuario es requerida"))
			return
		}
		if req.FechaDeCreacion == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "la fecha de creacion del usuario es requerida"))
			return
		}

		u, err := cUser.service.Store(ctx, req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo, req.FechaDeCreacion)
		if err != nil {
			ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(201, web.NewResponse(201, u, ""))
	}
}

func (cUser *User) GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		idInt, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
			return
		}
		user, err := cUser.service.GetUser(ctx, idInt)
		if err != nil {
			ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, user, ""))
	}
}

// UpdateUsers godoc
// @Summary Update users
// @Tags Users
// @Description update users
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param product body request true "User to update"
// @Success 200 {object} web.Response
// @Router /users [put]
func (cUser *User) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idInt, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		if req.Nombre == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "el nombre del usuario es requerido"))
			return
		}
		if req.Apellido == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "el apellido del usuario es requerido"))
			return
		}
		if req.Email == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "el email del usuario es requerido"))
			return
		}
		if req.Edad == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "la edad del usuario es requerida"))
			return
		}
		if req.Altura == 0.0 {
			ctx.JSON(400, web.NewResponse(400, nil, "la altura del usuario es requerida"))
			return
		}
		if req.FechaDeCreacion == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "la fecha de creacion del usuario es requerida"))
			return
		}

		u, err := cUser.service.Update(ctx, idInt, req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo, req.FechaDeCreacion)
		if err != nil {
			ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, u, ""))
	}
}

// DeleteUsers godoc
// @Summary Delete users
// @Tags Users
// @Description delete users
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param product body request true "User to delete"
// @Success 200 {object} web.Response
// @Router /users [delete]
func (cUser *User) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idInt, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
			return
		}

		err = cUser.service.Delete(ctx, idInt)

		if err != nil {
			ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, gin.H{"data": fmt.Sprintf("el usuario %d ha sido eliminado", idInt)}, ""))
	}
}

// UpdateUsers godoc
// @Summary Update users
// @Tags Users
// @Description update users
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param product body request true "User to update"
// @Success 200 {object} web.Response
// @Router /users [patch]
func (cUser *User) UpdateApellidoYEdad() gin.HandlerFunc {
	return func(ctx *gin.Context) {

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
			ctx.JSON(400, web.NewResponse(400, nil, "el apellido del usuario es requerido"))
			return
		}

		if req.Edad == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "la edad del usuario es requerida"))
			return
		}

		u, err := cUser.service.UpdateApellidoYEdad(ctx, idInt, req.Apellido, req.Edad)

		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, u, ""))
	}
}
