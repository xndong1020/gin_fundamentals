package services

import (
	entities "acy.com/api/src/entities"
	"acy.com/api/src/repositories"
)

type IAlbumService interface {
	FindAll() ([]entities.Album, error)
	FindById(id uint) (entities.Album, error)
	Create(newAlbum entities.Album) (entities.Album, error)
	Delete(id uint) error
}

type albumService struct {
	repo *repositories.IAlbumRepository
}

// constructor
func AlbumService(repo *repositories.IAlbumRepository) *albumService {
	return &albumService{ repo: repo }
}

/* interface implementations */
func (service *albumService) FindAll() ([]entities.Album, error) {
	repo := *service.repo
	albums, err := repo.FindAll();
	return albums, err 
}

func (service *albumService) FindById(id uint) (entities.Album, error) {
	repo := *service.repo
	album, err := repo.FindById(id);
	return album, err
}

func (service *albumService) Create(newAlbum entities.Album) (entities.Album, error) {
	repo := *service.repo
	album, err := repo.Create(newAlbum);
	return album, err
}

func (service *albumService) Delete(id uint) error {
	repo := *service.repo
	err := repo.Delete(id);
	return err
}