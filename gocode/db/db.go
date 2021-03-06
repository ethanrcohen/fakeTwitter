package db

import (
	"database/sql"
	"fmt"
	"os"

	// Postgres sql driver
	_ "github.com/lib/pq"
)

const (
	host = "postgres"
	port = 5432
)

var db *sql.DB

// GetDB gets a conection to the postgres DB, opening one if necessary.
func GetDB() *sql.DB {
	if db == nil {
		db = open()
	}

	err := db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func open() *sql.DB {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"))

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	return db
}
