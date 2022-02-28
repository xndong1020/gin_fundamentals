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
	fmt.Println("albums", albums)
    c.IndentedJSON(http.StatusAccepted, nil)
}