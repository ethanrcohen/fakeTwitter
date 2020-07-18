package main

import (
	"database/sql"
	"fmt"
	"time"
	"os"
	_ "github.com/lib/pq"
)

const(
	host = "postgres"
	port = 5432
)

// TODO: organize gocode, figure out module structure
func main() {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host,
			port,
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB"))

	// TODO: use retry logic rather than sleeping
	time.Sleep(3000 * time.Millisecond)
	db, err := sql.Open("postgres",	psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	fmt.Println("we've connected!")
}
