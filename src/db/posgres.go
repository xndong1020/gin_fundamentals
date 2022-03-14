package db

import (
	"database/sql"
	"fmt"
	"os"

	"acy.com/api/src/utils"
	_ "github.com/lib/pq"
)

const (
	user     = "root"
	password = "password"
)

func GetDbConnection() (*sql.DB, error) {
	utils.InitEnv()

	var (
		host   = os.Getenv("POSTGRES_HOST")
		port   = os.Getenv("POSTGRES_PORT")
		dbname = os.Getenv("POSTGRES_DBNAME")
		schema = os.Getenv("POSTGRES_SCHEMA")
	)

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=%s", host, port, user, password, dbname, schema)

	// open database
	db, err := sql.Open("postgres", psqlconn)

	return db, err
}

func PostgresDbProvider() *sql.DB {
	db, err := GetDbConnection()
	if err != nil {
		panic(err.Error())
	}
	return db
}
