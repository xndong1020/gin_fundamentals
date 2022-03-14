package services

import (
	"acy.com/api/src/models"
	"acy.com/api/src/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IAlbumMongoService interface {
	FindAll() []models.AlbumMongoDB
	FindById(id primitive.ObjectID) models.AlbumMongoDB
	Create(newAlbum models.AlbumMongoDB) string
	Delete(id primitive.ObjectID) bool
}

type albumMongoService struct {
	repo *repositories.IAlbumMongoDBRepository
}

// constructor
func AlbumMongoService(repo *repositories.IAlbumMongoDBRepository) *albumMongoService {
	return &albumMongoService{repo: repo}
}

func (service *albumMongoService) FindAll() []models.AlbumMongoDB {
	return (*service.repo).FindAll()
}

func (service *albumMongoService) FindById(id primitive.ObjectID) models.AlbumMongoDB {
	return (*service.repo).FindById(id)
}

func (service *albumMongoService) Create(newAlbum models.AlbumMongoDB) string {
	return (*service.repo).Create(newAlbum)
}

func (service *albumMongoService) Delete(id primitive.ObjectID) bool {
	return (*service.repo).Delete(id)
}