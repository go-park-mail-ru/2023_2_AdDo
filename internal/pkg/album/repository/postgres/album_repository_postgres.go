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

func (p Postgres) GetByTrackId(trackId uint64) ([]album.Base, error) {
	query := "select album.id, name, preview from album join album_track on album.id = album_track.album_id where album_track.track_id = $1"
	return p.getWithQuery(context.Background(), query, trackId)
}

func (p Postgres) GetByArtistId(artistId uint64) ([]album.Base, error) {
	query := "select id, name, preview from album where artist_id = $1"
	return p.getWithQuery(context.Background(), query, artistId)
}

// query := "select id, name, preview, content from track order by release_date desc limit $1"

func (p Postgres) GetByReleaseDate(limit uint32) ([]album.Base, error) {
	query := "select id, name, preview from album order by release_date desc limit $1"
	return p.getWithQuery(context.Background(), query, limit)
}

func (p Postgres) GetRandom(limit uint32) ([]album.Base, error) {
	//query := "select id, name, preview from album group by rating limit $1"
	query := "select id, name, preview from album limit $1"
	return p.getWithQuery(context.Background(), query, limit)
}

func (p Postgres) GetByListenCount(limit uint32) ([]album.Base, error) {
	query := "select id, name, preview from album order by play_count desc limit $1"
	return p.getWithQuery(context.Background(), query, limit)
}

func (p Postgres) GetByLikeCount(limit uint32) ([]album.Base, error) {
	query := "select id, name, preview from album order by rating desc limit $1"
	return p.getWithQuery(context.Background(), query, limit)
}

func (p Postgres) Get(albumId uint64) (album.Base, error) {
	var result album.Base

	query := "select id, name, preview from album where id = $1"
	err := p.Pool.QueryRow(context.Background(), query, albumId).Scan(&result.Id, &result.Name, &result.Preview)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (p Postgres) getWithQuery(ctx context.Context, query string, args ...any) ([]album.Base, error) {
	result := make([]album.Base, 0)

	rows, err := p.Pool.Query(ctx, query, args...)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		var base album.Base

		err = rows.Scan(&base.Id, &base.Name, &base.Preview)
		if err != nil {
			return nil, err
		}

		result = append(result, base)
	}

	return result, nil
}

func (p Postgres) CreateLike(userId, albumId uint64) error {
	query := "insert into profile_album (profile_id, album_id) values ($1, $2)"
	_, err := p.Pool.Exec(context.Background(), query, userId, albumId)
	if err != nil {
		return err
	}

	return nil
}
