1. init go module

```
go mod init module acy.com/api
```

2. install

```
go get -u github.com/gin-gonic/gin
```

3. Init server file

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

4. Create a bash file for start debugging

```sh
#!/bin/bash

go build src/server.go

# export GIN_MODE=release
export GIN_MODE=debug

./server
```

5. Add Router group

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

5. Add `controllers`

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
	fmt.Println("albums", albums)
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

6. Add nested router group

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
