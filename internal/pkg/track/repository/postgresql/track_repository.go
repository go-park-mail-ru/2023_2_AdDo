package track_repository

import (
	"context"
	postgres "main/internal/pkg/common/pgxiface"
	"main/internal/pkg/track"
)

type Postgres struct {
	Pool postgres.PgxIFace
}

func NewPostgres(pool postgres.PgxIFace) *Postgres {
	return &Postgres{Pool: pool}
}

//func (db *Postgres) GetAll() ([]track.Response, error) {
//	query := "select id, name, preview, content from track"
//	return db.getTracks(query, 0)
//}
//
//func (db *Postgres) GetPopular(limit uint32) ([]track.Response, error) {
//	query := "select id, name, preview, content from track order by play_count desc limit $1"
//	return db.getTracks(query, limit)
//}
//
//func (db *Postgres) GetLatest(limit uint32) ([]track.Response, error) {
//	query := "select id, name, preview, content from track order by release_date desc limit $1"
//	return db.getTracks(query, limit)
//}

func (db *Postgres) getWithQuery(ctx context.Context, query string, args ...any) ([]track.Response, error) {
	rows, err := db.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]track.Response, 0)
	for rows.Next() {
		var t track.Response
		err := rows.Scan(&t.Id, &t.Name, &t.Preview, &t.Content)
		if err != nil {
			return nil, err
		}
		result = append(result, t)
	}
	return result, nil
}

func (db *Postgres) GetByAlbum(albumId uint64) ([]track.Response, error) {
	query := "select track.id, name, preview, content from track join album_track on track.id = album_track.track_id where album_track.album_id = $1"
	return db.getWithQuery(context.Background(), query, albumId)
}

func (db *Postgres) GetByArtist(artistId uint64) ([]track.Response, error) {
	query := "select track.id, name, preview, content from track join artist_track on track.id = artist_track.track_id where artist_track.artist_id = $1"
	return db.getWithQuery(context.Background(), query, artistId)
}

func (db *Postgres) CreateLike(userId string, trackId uint64) error {
	query := "insert into profile_track (profile_id, track_id) values ($1, $2)"
	_, err := db.Pool.Exec(context.Background(), query, userId, trackId)
	if err != nil {
		return err
	}

	return nil
}

func (db *Postgres) AddListen(trackId uint64) error {
	query := "update track set play_count = play_count + 1 where id = $1"
	_, err := db.Pool.Exec(context.Background(), query, trackId)
	if err != nil {
		return err
	}

	return nil
}
