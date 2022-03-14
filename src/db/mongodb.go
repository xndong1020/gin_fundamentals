package db

import (
	"context"
	"fmt"
	"os"

	"acy.com/api/src/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MONGODB_USER     = "root"
	MONGODB_PASSWORD = "password"
)

func GetMongoDbConnection() *mongo.Client {
	utils.InitEnv()
	var (
		MONGODB_HOST = os.Getenv("MONGODB_HOST")
		MONGODB_PORT = os.Getenv("MONGODB_PORT")
	)
	connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%s", MONGODB_USER, MONGODB_PASSWORD, MONGODB_HOST, MONGODB_PORT)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
	if err != nil {
		panic(err)
	}
	return client
}

func GetMongoDb() *mongo.Database {
	utils.InitEnv()
	var (
		MONGODB_DBNAME = os.Getenv("MONGODB_DBNAME")
	)
	client := GetMongoDbConnection()
	return client.Database(MONGODB_DBNAME)
}
