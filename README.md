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

#### 14. GORM & Repository Pattern

##### Step 1: Create Table

Postgres

```sql
CREATE TABLE public.albums (
	id BIGSERIAL,
	title VARCHAR(255),
	artist VARCHAR(255),
	Price DECIMAL
)

```

###### Step 2: Model class for `albums` table

models/Album.go

```go
package models

type Album struct {
    Id     int  `json:"id" binding:"required,numeric,min=1" gorm:"primaryKey;autoIncrement;notnull"`
    Title  string  `json:"title" binding:"required"`
    Artist string  `json:"artist" binding:"required"`
    Price  float64 `json:"price" binding:"required,numeric,min=0"`
}
```

and a dto class for creating album, which doesn't have `id` property

models/CreateAlbum.go

```go
package models

type CreateAlbumDto struct {
    Title  string  `json:"title" binding:"required"`
    Artist string  `json:"artist" binding:"required"`
    Price  float64 `json:"price" binding:"required,numeric,min=0"`
}
```

##### Step 3: Database connection

db/database.go

```go
package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "password"
	dbname   = "postgres"
)

func GetDbConnection() (*sql.DB, error) {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)

	return db, err
}

```

##### Step 4: Repository class

repositories/AlbumRepository.go

```go
package repositories

import (
	"database/sql"

	"acy.com/api/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AlbumRepository struct {
	dbContext *gorm.DB
}

func NewAlbumRepository(sqlDB *sql.DB) AlbumRepository {
	gormDB, err := gorm.Open(postgres.New(postgres.Config{ Conn: sqlDB,}), &gorm.Config{})

	if err != nil {
		panic("gorm connection error")
	}

	return AlbumRepository{ dbContext: gormDB }
}

func (repo *AlbumRepository) FindAll() ([]models.Album, error) {
	albums := []models.Album{}
	result := repo.dbContext.Debug().Find(&albums)
	return albums, result.Error
}

func (repo *AlbumRepository) FindById(id int) (models.Album, error) {
	album := models.Album{}
	result := repo.dbContext.Debug().Find(&album, "id", id)
	return album, result.Error
}

func (repo *AlbumRepository) Create(newAlbum models.CreateAlbumDto) (models.Album, error) {
	album := models.Album{Title: newAlbum.Title, Artist: newAlbum.Artist, Price: newAlbum.Price}
	result := repo.dbContext.Debug().Create(&album)
	return album, result.Error
}

func (repo *AlbumRepository) Delete(id int) error {
	targetAlbum := models.Album{}
	result := repo.dbContext.Debug().Find(&targetAlbum, "id", id)
	if result.Error != nil {
		return result.Error
	}
	result = repo.dbContext.Debug().Delete(&targetAlbum, id)
	return result.Error
}


```

type `AlbumRepository` has a property `dbContext`, which hold the connection to db.

Then the function `NewAlbumRepository` is responsible for creating & initializing an `albumRepository` instance, and returns `AlbumRepository{ dbContext: gormDB }`.

```go
func NewAlbumRepository(sqlDB *sql.DB) AlbumRepository {
	gormDB, err := gorm.Open(postgres.New(postgres.Config{ Conn: sqlDB,}), &gorm.Config{})

	if err != nil {
		panic("gorm connection error")
	}

	return AlbumRepository{ dbContext: gormDB }
}
```

##### Step 5: Service class

services/albumService.go

```go
package services

import (
	"acy.com/api/src/models"
	"acy.com/api/src/repositories"
)

type AlbumService struct {
	repo repositories.AlbumRepository
}

func NewAlbumService(repo repositories.AlbumRepository) AlbumService {
	return AlbumService{ repo: repo }
}

func (service *AlbumService) FindAll() ([]models.Album, error) {
	albums, err := service.repo.FindAll();
	return albums, err
}

func (service *AlbumService) FindById(id int) (models.Album, error) {
	album, err := service.repo.FindById(id);
	return album, err
}

func (service *AlbumService) Create(newAlbum models.CreateAlbumDto) (models.Album, error) {
	album, err := service.repo.Create(newAlbum);
	return album, err
}

func (service *AlbumService) Delete(id int) error {
	err := service.repo.Delete(id);
	return err
}
```

##### Step 6: Use Service in controller class

Note: We firstly initialize an instance of albumService, and then use it in controller functions

```go
var conn, _ = db.GetDbConnection()
var albumService services.AlbumService = services.NewAlbumService(repositories.NewAlbumRepository(conn))
```

```go
package controllers

import (
	"net/http"
	"strconv"

	"acy.com/api/src/db"
	models "acy.com/api/src/models"
	"acy.com/api/src/repositories"
	"acy.com/api/src/services"
	"github.com/gin-gonic/gin"
)

var conn, _ = db.GetDbConnection()
var albumService services.AlbumService = services.NewAlbumService(repositories.NewAlbumRepository(conn))

// @Summary Get Albums list
// @ID get-albums-list
// @Description Get Albums list
// @Tags Album
// @Produce json
// @Success 200 {object} []models.Album
// @Router /albums [get]
func GetAlbums(c *gin.Context)  {
	albums, err := albumService.FindAll()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.Error{Message: err.Error()})
	}
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
		c.IndentedJSON(http.StatusBadRequest, models.Error{Message: "Invalid Album Id"})
	}

    album, err := albumService.FindById(int(id))

	if err != nil {
		// log.Fatalln("err",err)
		c.IndentedJSON(http.StatusBadRequest, models.Error{Message: err.Error()})
	}

	c.IndentedJSON(http.StatusOK, album)
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
	var newAlbum models.CreateAlbumDto

	// Call BindJSON to bind the received JSON to newAlbum.
    if err := c.ShouldBindJSON(&newAlbum); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Error{Message: err.Error()})
        return
    }

    // Add the new album to the slice.
    album, err := albumService.Create(newAlbum)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.Error{Message: err.Error()})
	}
    c.IndentedJSON(http.StatusCreated, album)
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
		c.IndentedJSON(http.StatusBadRequest, models.Error{Message: "Invalid Album Id"})
	}

	err = albumService.Delete(int(id))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, models.Error{Message: err.Error()})
	}

    c.IndentedJSON(http.StatusAccepted, nil)
}
```

#### 15. Dependency Injection in GO with Wire

Step 1: install google wire

```
go get -d github.com/google/wire/cmd/wire
```

Step 2:
For db connection, we need to create a provider function

```go
func GetDbConnection() (*sql.DB, error) {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)

	return db, err
}


func PostgresDbProvider() *sql.DB {
	db, err := GetDbConnection()
	if err != nil {
		panic(err.Error())
	}
	return db
}
```

Step 3:
Previously the albumController uses below code to create an instance of `albumService`

src/controllers/albumController

```go
var conn, _ = db.GetDbConnection()
var albumService services.AlbumService = services.NewAlbumService(repositories.NewAlbumRepository(conn))
```

to make it more readable:

```go
var conn = db.PostgresDbProvider()
var serviceRepository = repositories.NewAlbumRepository(conn)
var albumService services.AlbumService = services.NewAlbumService(serviceRepository)
```

Now with google wire, we can convert it to

```go
var albumService services.AlbumService = InitializeAlbumService();
```

and under `src` folder, create a `dependencies` folder, then create a file `wire.go`, passing in the initializers we want to use. There is **no rule on what order you should pass the initializers**.

src/dependencies.wire.go

```go
package dependencies

import (
	"acy.com/api/src/db"
	"acy.com/api/src/repositories"
	"acy.com/api/src/services"
	"github.com/google/wire"
)

func InitializeAlbumService() services.AlbumService {
    wire.Build(repositories.NewAlbumRepository, services.NewAlbumService, db.PostgresDbProvider)
    return services.AlbumService{}
}
```

Step 4: generate `wire_gen` file

From `src/dependencies` folder, run `wire` command, should see something like below

```
wire: acy.com/api/src/dependencies: wrote /home/isdance/Desktop/golang_projects/gin_fundamentals/src/dependencies/wire_gen.go
```

The generated `wire_gen` file looks like

```go
// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package dependencies

import (
	"acy.com/api/src/db"
	"acy.com/api/src/repositories"
	"acy.com/api/src/services"
)

// Injectors from wire.go:

func InitializeAlbumService() services.AlbumService {
	sqlDB := db.PostgresDbProvider()
	albumRepository := repositories.NewAlbumRepository(sqlDB)
	albumService := services.NewAlbumService(albumRepository)
	return albumService
}


```

From the albumController, import the `"acy.com/api/src/dependencies"` package

```go

package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"acy.com/api/src/dependencies"
	models "acy.com/api/src/models"
	"acy.com/api/src/services"
	"github.com/gin-gonic/gin"
)

// var conn = db.PostgresDbProvider()
// var serviceRepository = repositories.NewAlbumRepository(conn)
// var albumService services.AlbumService = services.NewAlbumService(serviceRepository)

var albumService services.AlbumService = dependencies.InitializeAlbumService();

...business logics....
```

Step 6: Commend out the code in `src/dependencies/wire.go`

Otherwise there will be a conflict between the template function `InitializeAlbumService` from `src/dependencies/wire.go`, with the generated function `InitializeAlbumService` from the `src/dependencies/wire_gen.go`

```
# acy.com/api/src/dependencies
src/dependencies/wire_gen.go:17:6: InitializeAlbumService redeclared in this block
        /home/isdance/Desktop/golang_projects/gin_fundamentals/src/dependencies/wire.go:10:31: previous declaration
```

Reference:
[Dependency Injection in GO with Wire](https://medium.com/wesionary-team/dependency-injection-in-go-with-wire-74f81cd222f6)
