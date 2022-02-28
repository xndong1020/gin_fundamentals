#### 1. init go module

```
go mod init module acy.com/api
```

#### 2. install

```
go get -u github.com/gin-gonic/gin
```

#### 3. Init server file

src/server.go

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // setup default router with some common middleware

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world!")
	})

    r.Run(":3000")
}
```

#### 4. Create a bash file for start debugging

```sh
#!/bin/bash

go build src/server.go

# export GIN_MODE=release
export GIN_MODE=debug

./server
```

#### 5. Add Router group

server.go

```go
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

	admin := r.Group("/albums")

	admin.GET("/", controllers.GetAlbums)
	admin.GET("/:id", controllers.GetAlbumById)
	admin.POST("/", controllers.CreateAlbum)
	admin.DELETE("/:id", controllers.DeleteAlbumById)

	r.Run(":3000")
}
```

#### 6. Add `controllers`

controllers/albumController.go

```go
package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	models "acy.com/api/src/models"
	"github.com/gin-gonic/gin"
)

var albums = []models.Album{
	{Id: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {Id: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {Id: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAlbums(c *gin.Context)  {
	 c.IndentedJSON(http.StatusOK, albums)
}

func GetAlbumById(c *gin.Context) {
    value := c.Param("id")
	id, err := strconv.ParseInt(value, 10, 0)

	if err != nil {
		log.Fatalln("err",err)
	}

    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.Id == int(id) {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func CreateAlbum(c *gin.Context) {
	var newAlbum models.Album

	// Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.ShouldBindJSON(&newAlbum); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

func DeleteAlbumById(c *gin.Context) {
	value := c.Param("id")
	id, err := strconv.ParseInt(value, 10, 0)

	if err != nil {
		log.Fatalln("err",err)
	}

	var index = -1

	for i, v := range albums {
        if v.Id == int(id) {
            index = i
			break
        }
    }

	if index == -1 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}

    // Remove the target album from the slice.
    albums = append(albums[:index], albums[index + 1:]...)
    c.IndentedJSON(http.StatusAccepted, nil)
}
```

#### 7. Add model class `Album`

models/album.go

```go
package models

type Album struct {
    Id     int  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}
```

#### 8. Add model class validation

```go
package models

type Album struct {
    Id     int  `json:"id" binding:"required,numeric,min=1"`
    Title  string  `json:"title" binding:"required"`
    Artist string  `json:"artist" binding:"required"`
    Price  float64 `json:"price" binding:"required,numeric,min=0"`
}
```

Meanwhile, in the controller we have a logic to catch the error when parse JSON from request body

```go
// Call BindJSON to bind the received JSON to newAlbum.
if err := c.ShouldBindJSON(&newAlbum); err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
}
```

An example of error message

```
status 400 Bad Request
{
    "error": "Key: 'Album.Price' Error:Field validation for 'Price' failed on the 'min' tag"
}
```

#### 9. Add nested Router Group

server.go

```go
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

		admin.GET("/", controllers.GetAlbums) // localhost:3000/api/v1/albums
		admin.GET("/:id", controllers.GetAlbumById) // localhost:3000/api/v1/album/4
		admin.POST("/", controllers.CreateAlbum) // localhost:3000/api/v1/albums
		admin.DELETE("/:id", controllers.DeleteAlbumById) // localhost:3000/api/v1/albums/4
	}

	r.Run(":3000")
}
```

#### 10. Install Swag

```
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
```

#### 11. Add annotations to api

src/server.go

```go
package main

import (
	"net/http"

	"acy.com/api/src/controllers"
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
	r := gin.Default() // setup default router with some common middleware

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
```

src/controllers/albumController.go

```go
package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	models "acy.com/api/src/models"
	"github.com/gin-gonic/gin"
)

var albums = []models.Album{
	{Id: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {Id: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {Id: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// @Summary Get Albums list
// @ID get-albums-list
// @Description Get Albums list
// @Tags Album
// @Produce json
// @Success 200 {object} []models.Album
// @Router /albums [get]
func GetAlbums(c *gin.Context)  {
	 c.IndentedJSON(http.StatusOK, albums)
}

// @Summary Get Album By Id
// @ID get-albums-by-id
// @Description Get Album By Id
// @Tags Album
// @Produce json
// @Param id path string true "album Id"
// @Success 200 {object} models.Album
// @Failure 404 {object} models.Error
// @Router /albums/{id} [get]
func GetAlbumById(c *gin.Context) {
    value := c.Param("id")
	id, err := strconv.ParseInt(value, 10, 0)

	if err != nil {
		log.Fatalln("err",err)
	}

    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.Id == int(id) {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, models.Error{Message: "album not found"})
}

// @Summary Create new album
// @ID create-new-album
// @Description Create new album
// @Tags Album
// @Produce json
// @Param data body models.Album true "album data"
// @Success 200 {object} models.Album
// @Failure 404 {object} models.Error
// @Router /albums [post]
func CreateAlbum(c *gin.Context) {
	var newAlbum models.Album

	// Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.ShouldBindJSON(&newAlbum); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Error{Message: err.Error()})
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

// @Summary Delete Album By Id
// @ID delete-albums-by-id
// @Description Delete Album By Id
// @Tags Album
// @Produce json
// @Param id path string true "album Id"
// @Success 200
// @Failure 404 {object} models.Error
// @Router /albums/{id} [delete]
func DeleteAlbumById(c *gin.Context) {
	value := c.Param("id")
	id, err := strconv.ParseInt(value, 10, 0)

	if err != nil {
		log.Fatalln("err",err)
	}

	var index = -1

	for i, v := range albums {
        if v.Id == int(id) {
            index = i
			break
        }
    }

	if index == -1 {
		c.IndentedJSON(http.StatusNotFound, models.Error{Message: "album not found"})
	}

    // Remove the target album from the slice.
    albums = append(albums[:index], albums[index + 1:]...)
	fmt.Println("albums", albums)
    c.IndentedJSON(http.StatusAccepted, nil)
}
```

#### 12: Initialize/Regenerate Swag

```
swag init

Run swag init in the project's root folder which contains the main.go file. This will parse your comments and generate the required files (docs folder and docs/docs.go


swag init -g src/server.go

If your General API annotations do not live in main.go, you can let swag know with -g flag.
```

If the generation is successful, the `docs` folder structure should be something like

```
.
├── controllers
│   └── albumController.go
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── models
│   ├── Album.go
│   └── Error.go
└── server.go
```

#### 13. Open Swagger url

Now you can open up the swagger from `http://localhost:3000/swagger/index.html`

![swagger](./images/swagger.png)
