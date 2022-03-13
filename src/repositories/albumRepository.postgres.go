package repositories

import (
	"database/sql"

	"acy.com/api/src/models"
	utils "acy.com/api/src/utils"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type IAlbumRepository interface {
	FindAll() ([]models.Album, error)
	FindById(id int) (models.Album, error)
	Create(newAlbum models.CreateAlbumDto) (models.Album, error)
	Delete(id int) error
}

type AlbumRepository struct {
	dbContext *gorm.DB
	logger	*zap.Logger
}

func NewAlbumRepository(sqlDB *sql.DB) *AlbumRepository {
	gormDB, err := gorm.Open(postgres.New(postgres.Config{ Conn: sqlDB,}), &gorm.Config{})
	logger := utils.NewLogger()
	
	if err != nil {
			logger.Error("gorm connection error", 
			zap.String("error", err.Error()),
		)
	}

	return &AlbumRepository{ dbContext: gormDB, logger: logger }
}

func (repo *AlbumRepository) FindAll() ([]models.Album, error) {
	albums := []models.Album{}
	result := repo.dbContext.Debug().Find(&albums)
	return albums, result.Error
}

func (repo *AlbumRepository) FindById(id int) (models.Album, error) {
	album := models.Album{}
	result := repo.dbContext.Debug().Find(&album, "id", id)
	return album, result.Error
}

func (repo *AlbumRepository) Create(newAlbum models.Album) (models.Album, error) {
	result := repo.dbContext.Debug().Create(&newAlbum)
	return newAlbum, result.Error
}

func (repo *AlbumRepository) Delete(id int) error {
	targetAlbum := models.Album{}
	result := repo.dbContext.Debug().Find(&targetAlbum, "id", id)
	if result.Error != nil {
		return result.Error
	}
	result = repo.dbContext.Debug().Delete(&targetAlbum, id)
	return result.Error
}

