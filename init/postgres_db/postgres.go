package init_postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
)

func InitPostgres(env string) (*pgx.Conn, error) {
	psqlInfo := os.Getenv(env)
	log.Printf(psqlInfo)

	conn, err := pgx.Connect(context.Background(), os.Getenv(env))
	if err != nil {
		return nil, err
	}
	fmt.Println("Postgres database successfully connected!")

	err = conn.Ping(context.Background())

	if err != nil {
		return nil, err
	}
	fmt.Println("Postgres database pinged")

	return conn, nil
}
