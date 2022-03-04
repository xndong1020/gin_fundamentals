package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "password"
	dbname   = "docker"
	schema   = "hollywood"
)

func GetDbConnection() (*sql.DB, error) {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable search_path=%s", host, port, user, password, dbname, schema)

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