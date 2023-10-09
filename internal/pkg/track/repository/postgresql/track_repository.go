package track_repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"main/internal/pkg/track"
)

type Postgres struct {
	database *pgx.Conn
}

func NewPostgres(db *pgx.Conn) *Postgres {
	return &Postgres{database: db}
}

func (db *Postgres) GetAll() ([]track.Response, error) {
	query := "select id, name, preview, content from track"
	rows, err := db.database.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]track.Response, 0)
	for rows.Next() {
		var t track.Response
		err = rows.Scan(&t.Id, &t.Name, &t.Preview, &t.Content)
		if err != nil {
			return nil, err
		}

		result = append(result, t)
	}
	return result, nil
}
