package services

import (
	"acy.com/api/src/models"
	"acy.com/api/src/repositories"
)

type IAlbumService interface {
	// repositories.IAlbumRepository
	FindAll() ([]models.Album, error)
	FindById(id int) (models.Album, error)
	Create(newAlbum models.CreateAlbumDto) (models.Album, error)
	Delete(id int) error
}

type AlbumService struct {
	repo *repositories.AlbumRepository
}

func NewAlbumService(repo *repositories.AlbumRepository) *AlbumService {
	return &AlbumService{ repo: repo }
}

func (service *AlbumService) FindAll() ([]models.Album, error) {
	albums, err := service.repo.FindAll();
	return albums, err 
}

func (service *AlbumService) FindById(id int) (models.Album, error) {
	album, err := service.repo.FindById(id);
	return album, err
}

func (service *AlbumService) Create(newAlbum models.Album) (models.Album, error) {
	album, err := service.repo.Create(newAlbum);
	return album, err
}

func (service *AlbumService) Delete(id int) error {
	err := service.repo.Delete(id);
	return err
}