package controllers

import (
	"net/http"
	"strconv"

	"acy.com/api/src/dependencies"
	models "acy.com/api/src/models"
	"acy.com/api/src/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// var albums = []models.Album{
// 	{Id: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
//     {Id: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
//     {Id: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }
// var conn = db.PostgresDbProvider()
// var serviceRepository = repositories.NewAlbumRepository(conn)
// var albumService services.AlbumService = services.NewAlbumService(serviceRepository)

var albumService *services.AlbumService = dependencies.InitializeAlbumService()

// var mongoDb = db.GetMongoDb()
// var albumMongoRepository = repositories.NewAlbumMongoDBRepository(mongoDb)
// var albumMongoService *services.AlbumMongoService = services.NewAlbumMongoService(albumMongoRepository)
var albumMongoService *services.AlbumMongoService = dependencies.InitializeAlbumMongoDBService()


// @Summary Get Albums list
// @ID get-albums-list
// @Description Get Albums list
// @Tags Album
// @Produce json
// @Success 200 {object} []models.AlbumResponse
// @Router /albums [get]
func GetAlbums(c *gin.Context)  {
	var response []models.AlbumResponse
	albums, err := albumService.FindAll()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.Error{Message: err.Error()})
	}
	albumsInMongo := albumMongoService.FindAll();

	// create a lookup map
	albumsInMongoLookup := map[string]string{}
	// convert the albumsInMongo into a map, the key is ObjectId(in postgres it is ContentId), value is content in mongodb
	for _, v := range albumsInMongo {
		albumsInMongoLookup[v.ID.Hex()] = v.Content
	}

	for _, v := range albums {
		if val, ok := albumsInMongoLookup[v.ContentId]; ok {
			response = append(response, models.AlbumResponse{ Id: v.Id, Title: v.Title, Artist: v.Artist, Price: v.Price, Content: val })
		}
	}

	c.IndentedJSON(http.StatusOK, response)
}

// @Summary Get Album By Id
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

    album, _ := albumService.FindById(int(id))
	objectId, err := primitive.ObjectIDFromHex(album.ContentId);

	albumInMongoDb := albumMongoService.FindById(objectId)

	response := models.AlbumResponse{Id: album.Id, Title: album.Title, Artist: album.Artist, Price: album.Price, Content: albumInMongoDb.Content}

	if err != nil {
		// log.Fatalln("err",err)
		c.IndentedJSON(http.StatusBadRequest, models.Error{Message: err.Error()})
	}

	c.IndentedJSON(http.StatusOK, response)
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

	contentId := albumMongoService.Create(models.AlbumMongoDB{Name: newAlbum.Title, Content: newAlbum.Content})

	if contentId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Error{Message: "Unable to save content to db"})
		return
	}

	album := models.Album{Title: newAlbum.Title, Artist: newAlbum.Artist, Price: newAlbum.Price, ContentId: contentId}
    album, err := albumService.Create(album)
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

	albumInDb, err := albumService.FindById(int(id))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound,  models.Error{Message: err.Error()})
	}

	err = albumService.Delete(int(id))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, models.Error{Message: err.Error()})
	}

	objectId, _ := primitive.ObjectIDFromHex(albumInDb.ContentId);
	isDeleteOk := albumMongoService.Delete(objectId)

	if isDeleteOk {
		c.IndentedJSON(http.StatusAccepted, nil)
		return
	} 
    c.IndentedJSON(http.StatusNotFound, models.Error{Message: "Invalid Content Id"})
}