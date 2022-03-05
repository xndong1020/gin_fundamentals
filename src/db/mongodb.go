package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongodb_host         = "localhost"
	mongodb_port         = 27017
	mongodb_user         = "root"
	mongodb_password     = "password"
	mongodb_dbname       = "albumsDb"
	mongodb_collection   = "albums"
)

func GetMongoDbConnection() *mongo.Client {
		connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%d", mongodb_user, mongodb_password, mongodb_host, mongodb_port)
        client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
        if err != nil {
            panic(err)
        }
		return client
}

func GetMongoDb() *mongo.Database {
	client := GetMongoDbConnection()
	return client.Database(mongodb_dbname)
}