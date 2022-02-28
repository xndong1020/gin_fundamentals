package main

import (
	"net/http"

	"acy.com/api/src/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // setup default router with some common middleware

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world!")
	})

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