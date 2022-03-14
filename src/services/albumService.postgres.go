package services

import (
	"acy.com/api/src/models"
	"acy.com/api/src/repositories"
)

type IAlbumService interface {
	FindAll() ([]models.Album, error)
	FindById(id int) (models.Album, error)
	Create(newAlbum models.Album) (models.Album, error)
	Delete(id int) error
}

type albumService struct {
	repo *repositories.IAlbumRepository
}

// constructor
func AlbumService(repo *repositories.IAlbumRepository) *albumService {
	return &albumService{ repo: repo }
}

/* interface implementations */
func (service *albumService) FindAll() ([]models.Album, error) {
	repo := *service.repo
	albums, err := repo.FindAll();
	return albums, err 
}

func (service *albumService) FindById(id int) (models.Album, error) {
	repo := *service.repo
	album, err := repo.FindById(id);
	return album, err
}

func (service *albumService) Create(newAlbum models.Album) (models.Album, error) {
	repo := *service.repo
	album, err := repo.Create(newAlbum);
	return album, err
}

func (service *albumService) Delete(id int) error {
	repo := *service.repo
	err := repo.Delete(id);
	return err
}