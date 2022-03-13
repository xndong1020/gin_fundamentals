package services

import (
	"acy.com/api/src/models"
	"acy.com/api/src/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AlbumMongoService struct {
	repo *repositories.AlbumMongoDBRepository
}

func NewAlbumMongoService(repo *repositories.AlbumMongoDBRepository) *AlbumMongoService {
	return &AlbumMongoService{repo: repo}
}

func (service *AlbumMongoService) FindAll() []models.AlbumMongoDB {
	return service.repo.FindAll()
}

func (service *AlbumMongoService) FindById(id primitive.ObjectID) models.AlbumMongoDB {
	return service.repo.FindById(id)
}

func (service *AlbumMongoService) Create(newAlbum models.AlbumMongoDB) string {
	return service.repo.Create(newAlbum)
}

func (service *AlbumMongoService) Delete(id primitive.ObjectID) bool {
	return service.repo.Delete(id)
}