package album

import (
	"main/internal/pkg/artist"
	"main/internal/pkg/track"
)

type Album struct {
	Id      uint64
	Name    string
	Track   []track.Track
	Artist  artist.Artist
	Preview string
}

type Response struct {
	Id      uint64
	Name    string
	Preview string
}

type Repository interface {
	Create(track Album) (uint64, error)
	GetById(id uint64) (Album, error)
	GetByTrackId(trackId uint64) ([]Response, error)
	GetByArtistId(albumId uint64) (Response, error)
}
