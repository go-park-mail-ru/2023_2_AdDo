package artist_repository

import (
	"database/sql"
	"main/internal/pkg/artist"
)

type Postgres struct {
	Db *sql.DB
}

func NewPostgres(db *sql.DB) Postgres {
	return Postgres{
		Db: db,
	}
}

func (repo *Postgres) Create(artist artist.Artist) (uint64, error) {
	// implement me
	return 0, nil
}

func (repo *Postgres) GetById(id uint64) (artist.Artist, error) {
	// implement me
	return artist.Artist{}, nil
}

func (repo *Postgres) GetByTrackId(trackId uint64) ([]artist.Response, error) {
	result := make([]artist.Response, 0)
	query := "select id, name, avatar from artist " +
		"join artist_track at on artist.id = artist_track.artist_id " +
		"where artist_track.track_id = $1"
	rows, err := repo.Db.Query(query, trackId)
	if err != nil {
		return nil, nil
	}
	defer rows.Close()
	for rows.Next() {
		var artist artist.Response
		err := rows.Scan(&artist.Id, &artist.Name, &artist.Avatar)
		if err != nil {
			return nil, err
		}
		result = append(result, artist)
	}

	return result, nil
}

func (repo *Postgres) GetByAlbumId(albumId uint64) (artist.Response, error) {
	var result artist.Response
	query := "select id, name, avatar from artist " +
		"join album at on artist.id = album.artist_id " +
		"where album.id = $1"
	err := repo.Db.QueryRow(query, albumId).Scan(&result.Id, &result.Name, &result.Avatar)
	if err != nil {
		return artist.Response{}, nil
	}

	return result, nil
}
