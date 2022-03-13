package main

import (
	"net/http"

	"acy.com/api/src/controllers"
	libs "acy.com/api/src/lib"
	"acy.com/api/src/middlewares"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "acy.com/api/src/docs"
)

// @title           Swagger ACY API
// @version         1.0
// @description     This is a web api server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath  /api/v1
func main() {
	logger := libs.NewZapLogger()
	r := gin.New() // disable default router and some common middleware
	r.Use(middlewares.GinLogger(logger), middlewares.GinRecovery(logger, true))

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world!")
	})

	// docs route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		admin := v1.Group("/albums")

		admin.GET("/", controllers.GetAlbums)
		admin.GET("/:id", controllers.GetAlbumById)
		admin.POST("/", controllers.CreateAlbum)
		admin.DELETE("/:id", controllers.DeleteAlbumById)
	}

	r.Run(":3000")
}
