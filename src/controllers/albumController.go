package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"acy.com/api/src/dependencies"
	"acy.com/api/src/entities"
	"acy.com/api/src/models"
	"github.com/gin-gonic/gin"
)

// var albums = []models.Album{
// 	{Id: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
//     {Id: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
//     {Id: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }
// sqlDB := db.PostgresDbProvider()
// var albumRepository repositories.IAlbumRepository = repositories.AlbumRepository(sqlDB)
// var albumService services.IAlbumService = services.AlbumService(&albumRepository)

var albumService = dependencies.InitializeAlbumService()

// database := db.GetMongoDb()
// var albumMongoDBRepository repositories.IAlbumMongoDBRepository = repositories.AlbumMongoDBRepository(database)
// var albumMongoService services.IAlbumMongoService = services.AlbumMongoService(&albumMongoDBRepository)
var albumMongoService = dependencies.InitializeAlbumMongoDBService()

// GetAlbums @Summary Get Albums list
// @ID get-albums-list
// @Description Get Albums list with pagination
// @Tags Album
// @Accept  json
// @Produce json
// @Param page query int true "pagination current page" default(0)
// @Param page_size query int true "pagination page_size" default(0)
// @Success 200 {object} []models.AlbumResponse
// @Router /albums [get]
func GetAlbums(c *gin.Context) {
	var response []models.AlbumResponse
	page, _ := strconv.Atoi(c.Query("page"))

    if page == 0 {
      page = 1
    }

	pageSize, _ := strconv.Atoi(c.Query("page_size"))

	albums, err := (*albumService).FindAll(page, pageSize)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.Error{Message: err.Error()})
	}
	albumsInMongo := (*albumMongoService).FindAll()

	// create a lookup map
	albumsInMongoLookup := map[uint]string{}
	// convert the albumsInMongo into a map, the key is ObjectId(in postgres it is ContentId), value is content in mongodb
	for _, v := range albumsInMongo {
		albumsInMongoLookup[v.AlbumId] = v.Content
	}

	for _, v := range albums {
		if val, ok := albumsInMongoLookup[v.Id]; ok {
			response = append(response, models.AlbumResponse{Id: v.Id, Title: v.Title, Artist: v.Artist, Price: v.Price, HasRead: v.HasRead, Content: val})
		} else {
			response = append(response, models.AlbumResponse{Id: v.Id, Title: v.Title, Artist: v.Artist, Price: v.Price, HasRead: v.HasRead})
		}
	}

	c.IndentedJSON(http.StatusOK, response)
}

// GetAlbumById @Summary Get Album By Id
// @ID get-albums-by-id
// @Description Get Album By Id
// @Tags Album
// @Produce json
// @Param id path string true "album Id"
// @Success 200 {object} models.AlbumResponse
// @Failure 404 {object} models.Error
// @Router /albums/{id} [get]
func GetAlbumById(c *gin.Context) {
	value := c.Param("id")
	id, err := strconv.ParseInt(value, 10, 0)

	if err != nil {
		// log.Fatalln("err",err)
		c.IndentedJSON(http.StatusBadRequest, models.Error{Message: "Invalid Album Id"})
	}

	album, _ := (*albumService).FindById(uint(id))
	// objectId, err := primitive.ObjectIDFromHex(album.ContentId);

	albumInMongoDb := (*albumMongoService).FindById(album.Id)

	response := models.AlbumResponse{Id: album.Id, Title: album.Title, Artist: album.Artist, Price: album.Price, HasRead: album.HasRead, Content: albumInMongoDb.Content}

	if err != nil {
		// log.Fatalln("err",err)
		c.IndentedJSON(http.StatusBadRequest, models.Error{Message: err.Error()})
	}

	c.IndentedJSON(http.StatusOK, response)
}

// CreateAlbum @Summary Create new album
// @ID create-new-album
// @Description Create new album
// @Tags Album
// @Produce json
// @Param data body models.CreateAlbumDto true "album data"
// @Success 200 {object} entities.Album
// @Failure 404 {object} models.Error
// @Router /albums [post]
func CreateAlbum(c *gin.Context) {
	var newAlbum models.CreateAlbumDto

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}

	album := entities.Album{Title: newAlbum.Title, Artist: newAlbum.Artist, Price: newAlbum.Price}
	album, err := (*albumService).Create(&album)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.Error{Message: err.Error()})
	}

	contentId := (*albumMongoService).Create(&entities.AlbumMongoDB{Name: newAlbum.Title, Content: newAlbum.Content, AlbumId: album.Id})

	if contentId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Error{Message: "Unable to save content to db"})
		return
	}

	c.IndentedJSON(http.StatusCreated, album)
}

// DeleteAlbumById @Summary Delete Album By Id
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
	fmt.Println("value", value)
	id, err := strconv.ParseInt(value, 10, 0)

	if err != nil {
		// log.Fatalln("err",err)
		c.IndentedJSON(http.StatusBadRequest, models.Error{Message: "Invalid Album Id"})
	}

	// albumInDb, err := (*albumService).FindById(uint(id))

	// if err != nil {
	// 	c.IndentedJSON(http.StatusNotFound,  models.Error{Message: err.Error()})
	// }

	err = (*albumService).Delete(uint(id))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, models.Error{Message: err.Error()})
	}

	// objectId, _ := primitive.ObjectIDFromHex(albumInDb.ContentId);
	isDeleteOk := (*albumMongoService).Delete(uint(id))

	if isDeleteOk {
		c.IndentedJSON(http.StatusAccepted, nil)
		return
	}
	c.IndentedJSON(http.StatusNotFound, models.Error{Message: "Invalid Content Id"})
}
