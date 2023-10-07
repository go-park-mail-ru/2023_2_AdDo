package init_postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func InitPostgres(env string) (*sql.DB, error) {
	psqlInfo := os.Getenv(env)
	log.Printf(psqlInfo)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	fmt.Println("Postgres database opened")

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Postgres database successfully connected!")

	return db, nil
}
