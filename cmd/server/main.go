package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rodrisouzag/go-web/cmd/server/handler"
	"github.com/rodrisouzag/go-web/docs"
	"github.com/rodrisouzag/go-web/internal/users"
	"github.com/rodrisouzag/go-web/pkg/store"
	"github.com/rodrisouzag/go-web/pkg/web"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func respondWithError(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, web.NewResponse(code, nil, message))
}

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	// We want to make sure the token is set, bail if not
	if requiredToken == "" {
		log.Fatal("Please set API_TOKEN environment variable")
	}

	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")

		if token == "" {
			respondWithError(c, 401, "API token required")
			return
		}

		if token != requiredToken {
			respondWithError(c, 401, "Invalid API token")
			return
		}

		c.Next()
	}
}

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.uy/es_uy/terminos-y-condiciones
// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.uy/support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	_ = godotenv.Load()
	db := store.New(store.FileType, "users.json")
	repo := users.NewRepository(db)
	service := users.NewService(repo)
	u := handler.NewUser(service)
	router := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("saludar/:nombre", u.Saludar())

	usr := router.Group("/users")
	usr.GET("/", TokenAuthMiddleware(), u.GetAll())
	usr.GET("/:id", TokenAuthMiddleware(), u.GetUser())
	usr.POST("/", TokenAuthMiddleware(), u.Store())
	usr.PUT("/:id", TokenAuthMiddleware(), u.Update())
	usr.PATCH("/:id", TokenAuthMiddleware(), u.UpdateApellidoYEdad())
	usr.DELETE("/:id", TokenAuthMiddleware(), u.Delete())

	router.Run()
}
