package track_repository

import (
	"database/sql"
	"main/internal/pkg/track"
)

type Postgres struct {
	database *sql.DB
}

func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{database: db}
}

func (db *Postgres) GetAll() ([]track.Response, error) {
	query := "select id, name, preview, content, play_count from track"
	rows, err := db.database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]track.Response, 0)
	for rows.Next() {
		var t track.Response
		err = rows.Scan(&t.Id, &t.Name, &t.Preview, &t.Content, &t.PlayCount)
		if err != nil {
			return nil, err
		}

		result = append(result, t)
	}
	return result, nil
}

func (db *Postgres) GetPopular(limit uint32) ([]track.Response, error) {
	query := "select id, name, preview, content, play_count from track order by play_count desc limit $1"
	rows, err := db.database.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]track.Response, 0)
	for rows.Next() {
		var t track.Response
		err = rows.Scan(&t.Id, &t.Name, &t.Preview, &t.Content, &t.PlayCount)
		if err != nil {
			return nil, err
		}

		result = append(result, t)
	}
	return result, nil
}

func (db *Postgres) getIds(query string, id uint64) ([]uint64, error) {
	rows, err := db.database.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	trackIds := make([]uint64, 0)
	for rows.Next() {
		var trackId uint64
		err = rows.Scan(&trackId)
		if err != nil {
			return nil, err
		}
		trackIds = append(trackIds, trackId)
	}
	return trackIds, nil
}

func (db *Postgres) GetTrackIdsByAlbum(albumId uint64) ([]uint64, error) {
	query := "select track_id from album_track where album_id = $1"
	return db.getIds(query, albumId)
}

func (db *Postgres) GetTrackIdsByArtist(artistId uint64) ([]uint64, error) {
	query := "select track_id from artist_track where artist_id = $1"
	return db.getIds(query, artistId)
}

func (db *Postgres) GetTrackIdsByPlaylist(playlistId uint64) ([]uint64, error) {
	query := "select track_id from playlist_track where playlist_id = $1"
	return db.getIds(query, playlistId)
}

func (db *Postgres) GetByTrackId(trackId uint64) (track.Response, error) {
	query := "select id, name, preview, content, play_count from track where id = $1"

	var t track.Response
	err := db.database.QueryRow(query, trackId).Scan(&t.Id, &t.Name, &t.Preview, &t.Content, &t.PlayCount)
	if err != nil {
		var empty track.Response
		return empty, err
	}
	return t, nil
}
