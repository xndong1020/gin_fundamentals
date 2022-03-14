package dependencies

// import (
// 	"acy.com/api/src/db"
// 	"acy.com/api/src/repositories"
// 	"acy.com/api/src/services"
// 	"github.com/google/wire"
// )

// func InitializeAlbumService() *services.AlbumService {
//     wire.Build(repositories.AlbumRepository, services.AlbumService, db.PostgresDbProvider)
//     return &services.AlbumService{}
// }

// func InitializeAlbumMongoDBService() *services.AlbumMongoService {
//     wire.Build(repositories.AlbumMongoDBRepository, services.AlbumMongoService, db.GetMongoDb)
//     return &services.AlbumMongoService{}
// }