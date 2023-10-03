package artist_repository

import (
	"database/sql"
	"main/internal/pkg/artist"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) Postgres {
	return Postgres{
		db: db,
	}
}

func (repo *Postgres) GetByTrackId(trackId uint64) ([]artist.Response, error) {
	result := make([]artist.Response, 0)
	query := "select artist.id, name, avatar from artist join artist_track on artist.id = artist_track.artist_id where artist_track.track_id = $1"
	rows, err := repo.db.Query(query, trackId)
	if err != nil {
		return nil, err
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
	query := "select artist.id, artist.name, avatar from artist join album on artist.id = album.artist_id where album.id = $1"
	err := repo.db.QueryRow(query, albumId).Scan(&result.Id, &result.Name, &result.Avatar)
	if err != nil {
		return artist.Response{}, err
	}

	return result, nil
}
