package repositories

import (
	"database/sql"

	"acy.com/api/src/entities"
	libs "acy.com/api/src/lib"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type IAlbumRepository interface {
	FindAll() ([]entities.Album, error)
	FindById(id uint) (entities.Album, error)
	Create(newAlbum *entities.Album) (entities.Album, error)
	Update(id uint, column string, value interface{})
	Delete(id uint) error
}

type albumRepository struct {
	dbContext *gorm.DB
	logger    *zap.Logger
}

// AlbumRepository constructor
func AlbumRepository(sqlDB *sql.DB) *albumRepository {
	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{})
	logger := libs.NewZapLogger()

	if err != nil {
		logger.Error("gorm connection error",
			zap.String("error", err.Error()),
		)
	}

	return &albumRepository{dbContext: gormDB, logger: logger}
}

// FindAll /* interface implementations */
func (repo *albumRepository) FindAll() ([]entities.Album, error) {
	var albums []entities.Album
	result := repo.dbContext.Debug().Find(&albums)
	return albums, result.Error
}

func (repo *albumRepository) FindById(id uint) (entities.Album, error) {
	album := entities.Album{}
	result := repo.dbContext.Debug().Find(&album, "id", id)
	return album, result.Error
}

func (repo *albumRepository) Create(newAlbum *entities.Album) (entities.Album, error) {
	result := repo.dbContext.Debug().Create(&newAlbum)
	return *newAlbum, result.Error
}

func (repo *albumRepository) Update(id uint, column string, value interface{}) {
	album := entities.Album{}
	repo.dbContext.Debug().Model(&album).Where("id = ?", id).Update(column, value)
}

func (repo *albumRepository) Delete(id uint) error {
	targetAlbum := entities.Album{}
	result := repo.dbContext.Debug().Find(&targetAlbum, "id", id)
	if result.Error != nil {
		return result.Error
	}
	result = repo.dbContext.Debug().Delete(&targetAlbum, id)
	return result.Error
}
