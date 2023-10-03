package init_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// db_connection_string = "host=db port=5432 user=postgres password=music dbname=wave sslmode=disable TimeZone=Europe/Moscow"

func InitPostgres(env string) (*sql.DB, error) {
	psqlInfo := os.Getenv(env)
	log.Printf(psqlInfo)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	fmt.Println("Database opened")

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Database successfully connected!")

	return db, nil
}
