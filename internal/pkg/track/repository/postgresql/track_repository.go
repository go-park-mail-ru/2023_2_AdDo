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

func (db *Postgres) Create(track track.Track) (uint64, error) {
	return 0, nil
}

func (db *Postgres) GetById(trackId uint64) (track.Track, error) {
	return track.Track{
		Id:          0,
		Name:        "",
		IsSong:      false,
		FKArtistId:  0,
		FKAlbumId:   0,
		ImagePath:   "",
		ContentPath: "",
	}, nil
}

func (db *Postgres) GetAll() ([]track.Track, error) {
	return nil, nil
}
