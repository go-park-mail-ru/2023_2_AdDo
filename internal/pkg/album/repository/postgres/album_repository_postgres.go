package album_repository

import (
	"database/sql"
	"main/internal/pkg/album"
)

type Postgres struct {
	db *sql.DB
}

func (p Postgres) Create(track album.Album) (uint64, error) {
	//TODO implement me
	panic("implement me")
}

func (p Postgres) GetById(id uint64) (album.Album, error) {
	//TODO implement me
	panic("implement me")
}

func (p Postgres) GetByTrackId(trackId uint64) ([]album.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (p Postgres) GetByArtistId(albumId uint64) (album.Response, error) {
	//TODO implement me
	panic("implement me")
}

func NewPostgres(db *sql.DB) Postgres {
	return Postgres{
		db: db,
	}
}
