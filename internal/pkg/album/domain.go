package album

import "main/internal/pkg/track"

type Id struct {
	Id uint64 `json:"Id" example:"1"`
}

type Base struct {
	Id      uint64 `json:"Id" example:"1"`
	Name    string `json:"Name" example:"AlbumName"`
	Preview string `json:"Preview" example:"AlbumPreview"`
}

type Response struct {
	Id         uint64           `json:"Id" example:"1"`
	Name       string           `json:"Name" example:"AlbumName"`
	Preview    string           `json:"Preview" example:"AlbumPreview"`
	ArtistId   uint64           `json:"ArtistId" example:"1"`
	ArtistName string           `json:"ArtistName" example:"ArtistName"`
	Tracks     []track.Response `json:"Tracks"`
}

type UseCase interface {
	GetRandom() ([]Response, error)
	GetMostLiked() ([]Response, error)
	GetPopular() ([]Response, error)
	GetNew() ([]Response, error)
	GetAlbum(albumId uint64) (Response, error)
	Like(userId string, albumId uint64) error
}

type Repository interface {
	Get(albumId uint64) (Base, error)
	GetByTrackId(trackId uint64) ([]Base, error)
	GetByArtistId(artistId uint64) ([]Base, error)
	GetByLikeCount(limit uint32) ([]Base, error)
	GetByReleaseDate(limit uint32) ([]Base, error)
	GetByListenCount(limit uint32) ([]Base, error)
	GetRandom(limit uint32) ([]Base, error)
	CreateLike(userId string, albumId uint64) error
}

const LimitForMainPage = 10
