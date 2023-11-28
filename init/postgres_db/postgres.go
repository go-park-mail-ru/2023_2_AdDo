package init_postgres

import (
	"context"
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

	if err = conn.Ping(context.Background()); err != nil {
		return nil, err
	}

	return conn, nil
}
