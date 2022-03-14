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
	albums, err := (*service.repo).FindAll();
	return albums, err 
}

func (service *albumService) FindById(id uint) (entities.Album, error) {
	albumInDb, err := (*service.repo).FindById(id);
	if !albumInDb.HasRead {
		(*service.repo).Update(albumInDb.Id, "has_read", true)
	}
	return albumInDb, err
}

func (service *albumService) Create(newAlbum entities.Album) (entities.Album, error) {
	album, err := (*service.repo).Create(newAlbum);
	return album, err
}

func (service *albumService) Delete(id uint) error {
	err := (*service.repo).Delete(id);
	return err
}