package album_repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"main/internal/pkg/album"
)

type Postgres struct {
	db *pgx.Conn
}

func NewPostgres(db *pgx.Conn) Postgres {
	return Postgres{
		db: db,
	}
}

func (p Postgres) GetByTrackId(trackId uint64) ([]album.Response, error) {
	result := make([]album.Response, 0)
	query := "select album.id, name, preview from album join album_track on album.id = album_track.album_id where album_track.track_id = $1"
	rows, err := p.db.Query(context.Background(), query, trackId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var a album.Response
		err := rows.Scan(&a.Id, &a.Name, &a.Preview)
		if err != nil {
			return nil, err
		}
		result = append(result, a)
	}
	return result, nil
}

func (p Postgres) GetByArtistId(artistId uint64) ([]album.Response, error) {
	result := make([]album.Response, 0)
	query := "select id, name, preview from album where artist_id = $1"

	rows, err := p.db.Query(context.Background(), query, artistId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var a album.Response
		err := rows.Scan(&a.Id, &a.Name, &a.Preview)
		if err != nil {
			return nil, err
		}
		result = append(result, a)
	}

	return result, nil
}
