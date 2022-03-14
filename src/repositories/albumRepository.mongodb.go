package repositories

import (
	"context"

	entities "acy.com/api/src/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IAlbumMongoDBRepository interface {
	FindAll() []entities.AlbumMongoDB
	FindById(albumId uint) entities.AlbumMongoDB
	Create(newAlbum *entities.AlbumMongoDB) string
	Delete(albumId uint) bool
}

type albumMongoDBRepository struct {
	dbContext *mongo.Database
}

// AlbumMongoDBRepository constructor
func AlbumMongoDBRepository(db *mongo.Database) *albumMongoDBRepository {
	albumMongoDBRepository := albumMongoDBRepository{dbContext: db}
	return &albumMongoDBRepository
}

func (albumRepo *albumMongoDBRepository) FindAll() []entities.AlbumMongoDB {
	var results []entities.AlbumMongoDB
	cursor, err := albumRepo.dbContext.Collection("albums").Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	for cursor.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var elem entities.AlbumMongoDB
		err := cursor.Decode(&elem)
		if err != nil {
			panic(err)
		}

		results = append(results, elem)
	}

	return results
}

func (albumRepo *albumMongoDBRepository) FindById(albumId uint) entities.AlbumMongoDB {
	var album entities.AlbumMongoDB
	if err := albumRepo.dbContext.Collection("albums").FindOne(context.TODO(), bson.M{"albumId": albumId}).Decode(&album); err != nil {
		panic(err)
	}
	return album
}

func (albumRepo *albumMongoDBRepository) Create(newAlbum *entities.AlbumMongoDB) string {
	result, err := albumRepo.dbContext.Collection("albums").InsertOne(context.TODO(), *newAlbum)

	if err != nil {
		panic(err)
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex()
	}

	return ""
}

func (albumRepo *albumMongoDBRepository) Delete(albumId uint) bool {
	result, err := albumRepo.dbContext.Collection("albums").DeleteOne(context.TODO(), bson.M{"albumId": albumId})

	if err != nil {
		panic(err)
	}
	return result.DeletedCount > 0
}
