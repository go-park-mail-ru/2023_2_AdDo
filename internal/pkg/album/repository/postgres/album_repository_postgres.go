package album_repository

import (
	"context"
	"main/internal/pkg/album"
	postgres "main/internal/pkg/common/pgxiface"
)

type Postgres struct {
	Pool postgres.PgxIFace
}

func NewPostgres(pool postgres.PgxIFace) Postgres {
	return Postgres{
		Pool: pool,
	}
}

func (p Postgres) GetByTrackId(trackId uint64) ([]album.Response, error) {
	result := make([]album.Response, 0)
	query := "select album.id, name, preview from album join album_track on album.id = album_track.album_id where album_track.track_id = $1"
	rows, err := p.Pool.Query(context.Background(), query, trackId)
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

	rows, err := p.Pool.Query(context.Background(), query, artistId)
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
