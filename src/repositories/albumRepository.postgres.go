package repositories

import (
	"database/sql"

	libs "acy.com/api/src/lib"
	"acy.com/api/src/models"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type IAlbumRepository interface {
	FindAll() ([]models.Album, error)
	FindById(id int) (models.Album, error)
	Create(newAlbum models.Album) (models.Album, error)
	Delete(id int) error
}

type albumRepository  struct {
	dbContext *gorm.DB
	logger	*zap.Logger
}

// constructor
func AlbumRepository(sqlDB *sql.DB) *albumRepository {
	gormDB, err := gorm.Open(postgres.New(postgres.Config{ Conn: sqlDB,}), &gorm.Config{})
	logger := libs.NewZapLogger()
	
	if err != nil {
			logger.Error("gorm connection error", 
			zap.String("error", err.Error()),
		)
	}

	return &albumRepository{ dbContext: gormDB, logger: logger }
}

/* interface implementations */
func (repo *albumRepository) FindAll() ([]models.Album, error) {
	albums := []models.Album{}
	result := repo.dbContext.Debug().Find(&albums)
	return albums, result.Error
}

func (repo *albumRepository) FindById(id int) (models.Album, error) {
	album := models.Album{}
	result := repo.dbContext.Debug().Find(&album, "id", id)
	return album, result.Error
}

func (repo *albumRepository) Create(newAlbum models.Album) (models.Album, error) {
	result := repo.dbContext.Debug().Create(&newAlbum)
	return newAlbum, result.Error
}

func (repo *albumRepository) Delete(id int) error {
	targetAlbum := models.Album{}
	result := repo.dbContext.Debug().Find(&targetAlbum, "id", id)
	if result.Error != nil {
		return result.Error
	}
	result = repo.dbContext.Debug().Delete(&targetAlbum, id)
	return result.Error
}

