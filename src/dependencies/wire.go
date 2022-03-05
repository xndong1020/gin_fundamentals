package dependencies

// import (
// 	"acy.com/api/src/db"
// 	"acy.com/api/src/repositories"
// 	"acy.com/api/src/services"
// 	"github.com/google/wire"
// )

// func InitializeAlbumService() *services.AlbumService {
//     wire.Build(repositories.NewAlbumRepository, services.NewAlbumService, db.PostgresDbProvider)
//     return &services.AlbumService{}
// }

// func InitializeAlbumMongoDBService() *services.AlbumMongoService {
//     wire.Build(repositories.NewAlbumMongoDBRepository, services.NewAlbumMongoService, db.GetMongoDb)
//     return &services.AlbumMongoService{}
// }