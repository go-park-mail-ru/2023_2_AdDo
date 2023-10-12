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

//func (db *Postgres) getTracks(query string, limit uint32) ([]track.Response, error) {
//	rows, err := db.Pool.Query(context.Background(), query, limit)
//
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	result := make([]track.Response, 0)
//	for rows.Next() {
//		var t track.Response
//		err = rows.Scan(&t.Id, &t.Name, &t.Preview, &t.Content)
//		if err != nil {
//			return nil, err
//		}
//
//		result = append(result, t)
//	}
//	return result, nil
//}

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
//
//func (db *Postgres) GetByPlaylist(playlistId uint64) ([]track.Response, error) {
//	query := "select track_id from playlist_track where playlist_id = $1"
//	return db.getTracksById(query, playlistId)
//}

func (db *Postgres) getTracksById(query string, id uint64) ([]track.Response, error) {
	rows, err := db.Pool.Query(context.Background(), query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tracks := make([]track.Response, 0)
	for rows.Next() {
		var trackId int
		err = rows.Scan(&trackId)
		if err != nil {
			return nil, err
		}
		t, err := db.getById(uint64(trackId))
		if err != nil {
			return nil, err
		}
		tracks = append(tracks, t)
	}
	return tracks, nil
}

func (db *Postgres) getById(trackId uint64) (track.Response, error) {
	query := "select id, name, preview, content from track where id = $1"

	var t track.Response
	err := db.Pool.QueryRow(context.Background(), query, trackId).Scan(&t.Id, &t.Name, &t.Preview, &t.Content)
	if err != nil {
		var empty track.Response
		return empty, err
	}
	return t, nil
}

func (db *Postgres) GetByAlbum(albumId uint64) ([]track.Response, error) {
	query := "select track_id from album_track where album_id = $1"
	return db.getTracksById(query, albumId)
}

func (db *Postgres) CreateLike(userId, trackId uint64) error {
	query := "insert into like_track (profile_id, track_id) values ($1, $2)"
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

func (db *Postgres) GetByArtist(artistId uint64) ([]track.Response, error) {
	query := "select track_id from artist_track where artist_id = $1"
	return db.getTracksById(query, artistId)
}
