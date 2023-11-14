package init_postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
)

func InitPostgres(env string) (*pgxpool.Pool, error) {
	psqlInfo := os.Getenv(env)
	log.Print(psqlInfo)

	conn, err := pgxpool.New(context.Background(), os.Getenv(env))
	if err != nil {
		return nil, err
	}
	fmt.Println("Postgres db successfully connected!")

	if err = conn.Ping(context.Background()); err != nil {
		return nil, err
	}
	fmt.Println("Postgres db pinged")

	return conn, nil
}
