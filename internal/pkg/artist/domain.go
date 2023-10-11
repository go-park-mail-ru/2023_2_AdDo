package artist

import (
	"main/internal/pkg/album"
	"main/internal/pkg/track"
)

type Base struct {
	Id     uint64 `json:"Id" example:"1"`
	Name   string `json:"Name" example:"ArtistName"`
	Avatar string `json:"Avatar" example:"ArtistAvatar"`
}

type Response struct {
	Id     uint64 `json:"Id" example:"1"`
	Name   string `json:"Name" example:"ArtistName"`
	Avatar string `json:"Avatar" example:"ArtistAvatar"`
	Albums []album.Base
	Tracks []track.Response
}

type UseCase interface {
	GetArtistInfo(artistId uint64) (Response, error)
}

type Repository interface {
	GetByTrackId(trackId uint64) ([]Base, error)
	GetByAlbumId(albumId uint64) (Base, error)
	GetTracks(artistId uint64) ([]track.Response, error)
	Get(artistId uint64) (Base, error)
}
