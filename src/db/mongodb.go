package db

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongodb_user     = "root"
	mongodb_password = "password"
)

func GetMongoDbConnection() *mongo.Client {
	godotenv.Load(".env")
	var (
		mongodb_host = os.Getenv("MONGODB_HOST")
		mongodb_port = os.Getenv("MONGODB_PORT")
	)
	connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%s", mongodb_user, mongodb_password, mongodb_host, mongodb_port)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
	if err != nil {
		panic(err)
	}
	return client
}

func GetMongoDb() *mongo.Database {
	godotenv.Load(".env")
	var (
		mongodb_dbname = os.Getenv("MONGODB_DBNAME")
	)
	client := GetMongoDbConnection()
	return client.Database(mongodb_dbname)
}
