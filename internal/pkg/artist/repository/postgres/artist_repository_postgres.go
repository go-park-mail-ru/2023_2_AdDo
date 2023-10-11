package artist_repository

import (
	"context"
	"main/internal/pkg/artist"
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

func (repo *Postgres) Get(artistId uint64) (artist.Base, error) {
	var result artist.Base

	query := "select artist.id, name, avatar from artist where artist.id = $1"
	err := repo.Pool.QueryRow(context.Background(), query, artistId).Scan(&result.Id, &result.Name, &result.Avatar)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (repo *Postgres) GetByTrackId(trackId uint64) ([]artist.Base, error) {
	result := make([]artist.Base, 0)

	query := "select artist.id, name, avatar from artist join artist_track on artist.id = artist_track.artist_id where artist_track.track_id = $1"
	rows, err := repo.Pool.Query(context.Background(), query, trackId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var a artist.Base

		err := rows.Scan(&a.Id, &a.Name, &a.Avatar)
		if err != nil {
			return nil, err
		}

		result = append(result, a)
	}

	return result, nil
}

func (repo *Postgres) GetByAlbumId(albumId uint64) (artist.Base, error) {
	var result artist.Base
	query := "select artist.id, artist.name, avatar from artist join album on artist.id = album.artist_id where album.id = $1"
	err := repo.Pool.QueryRow(context.Background(), query, albumId).Scan(&result.Id, &result.Name, &result.Avatar)
	if err != nil {
		return artist.Base{}, err
	}

	return result, nil
}
