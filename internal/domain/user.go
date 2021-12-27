package domain

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
