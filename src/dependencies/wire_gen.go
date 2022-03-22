//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package dependencies

import (
	"acy.com/api/src/db"
	"acy.com/api/src/repositories"
	"acy.com/api/src/services"
)

// Injectors from wire.go:

func InitializeAlbumService() *services.IAlbumService {
	conn := db.PostgresDbProvider()
	var albumRepository repositories.IAlbumRepository = repositories.NewAlbumRepository(conn)
	var albumService services.IAlbumService = services.AlbumService(&albumRepository)
	return &albumService
}

func InitializeAlbumMongoDBService() *services.IAlbumMongoService {
	database := db.GetMongoDb()
	var albumMongoDBRepository repositories.IAlbumMongoDBRepository = repositories.AlbumMongoDBRepository(database)
	var albumMongoService services.IAlbumMongoService = services.AlbumMongoService(&albumMongoDBRepository)
	return &albumMongoService
}
