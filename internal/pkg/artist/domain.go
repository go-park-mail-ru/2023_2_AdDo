package artist

import (
	"main/internal/pkg/album"
	"main/internal/pkg/playlist"
	"main/internal/pkg/track"
)

type Base struct {
	Id     uint64 `json:"Id" example:"1"`
	Name   string `json:"Name" example:"ArtistName"`
	Avatar string `json:"Avatar" example:"ArtistAvatar"`
}

type Response struct {
	Id     uint64           `json:"Id" example:"1"`
	Name   string           `json:"Name" example:"ArtistName"`
	Avatar string           `json:"Avatar" example:"ArtistAvatar"`
	Albums []album.Base     `json:"Albums"`
	Tracks []track.Response `json:"Tracks"`
}

type Artists struct {
	Artists []Base `json:"Artists"`
}

type SearchResponse struct {
	Playlists []playlist.Base  `json:"Playlists"`
	Albums    []album.Base     `json:"Albums"`
	Tracks    []track.Response `json:"Tracks"`
	Artists   []Base           `json:"Artists"`
}

type UseCase interface {
	GetArtistInfo(artistId uint64) (Response, error)
	Like(userId string, artistId uint64) error
	IsLike(userId string, artistId uint64) (bool, error)
	Unlike(userId string, artistId uint64) error
	FullSearch(text string) (SearchResponse, error)
	GetUserArtists(userId string) (Artists, error)
}

type Repository interface {
	GetByTrackId(trackId uint64) ([]Base, error)
	GetByAlbumId(albumId uint64) (Base, error)
	Get(artistId uint64) (Base, error)
	CreateLike(userId string, artistId uint64) error
	CheckLike(userId string, artistId uint64) (bool, error)
	DeleteLike(userId string, artistId uint64) error
	Search(text string) ([]Base, error)
	GetByUserId(userId string) ([]Base, error)
}
