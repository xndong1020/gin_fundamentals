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

// var albums = []models.Album{
// 	{Id: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
//     {Id: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
//     {Id: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }
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
		// log.Fatalln("err",err)
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
		// log.Fatalln("err",err)
		c.IndentedJSON(http.StatusBadRequest, models.Error{Message: "Invalid Album Id"})
	}

	err = albumService.Delete(int(id))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, models.Error{Message: err.Error()})
	}
	
    c.IndentedJSON(http.StatusAccepted, nil)
}