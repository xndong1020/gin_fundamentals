package services

import (
	entities "acy.com/api/src/entities"
	"acy.com/api/src/repositories"
)

type IAlbumMongoService interface {
	FindAll() []entities.AlbumMongoDB
	FindById(id uint) entities.AlbumMongoDB
	Create(newAlbum *entities.AlbumMongoDB) string
	Delete(id uint) bool
}

type albumMongoService struct {
	repo *repositories.IAlbumMongoDBRepository
}

// constructor
func AlbumMongoService(repo *repositories.IAlbumMongoDBRepository) *albumMongoService {
	return &albumMongoService{repo: repo}
}

func (service *albumMongoService) FindAll() []entities.AlbumMongoDB {
	return (*service.repo).FindAll()
}

func (service *albumMongoService) FindById(id uint) entities.AlbumMongoDB {
	return (*service.repo).FindById(id)
}

func (service *albumMongoService) Create(newAlbum *entities.AlbumMongoDB) string {
	return (*service.repo).Create(newAlbum)
}

func (service *albumMongoService) Delete(id uint) bool {
	return (*service.repo).Delete(id)
}