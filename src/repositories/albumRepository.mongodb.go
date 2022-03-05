package repositories

import (
	"context"

	"acy.com/api/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AlbumMongoDBRepository struct {
	dbContext *mongo.Database
}

func NewAlbumMongoDBRepository(db *mongo.Database) *AlbumMongoDBRepository {
	 albumMongoDBRepository := AlbumMongoDBRepository{dbContext: db}
	 return &albumMongoDBRepository
}

func (albumRepo *AlbumMongoDBRepository) FindAll() []models.AlbumMongoDB  {
	var results []models.AlbumMongoDB
	cursor, err := albumRepo.dbContext.Collection("albums").Find(context.TODO(), bson.D{})
	if err != nil {
        panic(err)
	}

	 for cursor.Next(context.TODO()) {
        //Create a value into which the single document can be decoded
        var elem models.AlbumMongoDB
        err := cursor.Decode(&elem)
        if err != nil {
            panic(err)
        }

        results = append(results, elem)
    }

	return results
}

func (albumRepo *AlbumMongoDBRepository) FindById(id primitive.ObjectID) models.AlbumMongoDB  {
	var album models.AlbumMongoDB
	if err :=  albumRepo.dbContext.Collection("albums").FindOne(context.TODO(), bson.M{"_id": id}).Decode(&album); err != nil {
        panic(err)
	}
	return album
}

func (albumRepo *AlbumMongoDBRepository) Create(newAlbum models.AlbumMongoDB) string {
	result, err := albumRepo.dbContext.Collection("albums").InsertOne(context.TODO(), newAlbum)

	if err != nil {
		panic(err)
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
    	return oid.Hex()
	}

	return ""
}

func (albumRepo *AlbumMongoDBRepository) Delete(id primitive.ObjectID) bool {
	result, err :=  albumRepo.dbContext.Collection("albums").DeleteOne(context.TODO(), bson.M{"_id": id}) 
	
	if err != nil {
        panic(err)
	}

	return result.DeletedCount > 0
}