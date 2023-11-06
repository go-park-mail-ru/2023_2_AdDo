package track

import (
	"errors"
)

type Id struct {
	Id uint64 `json:"Id" example:"1"`
}

type LikedTracks struct {
	Tracks []Response `json:"Tracks"`
}

type Response struct {
	Id         uint64 `json:"Id" example:"1"`
	ArtistId   uint64 `json:"ArtistId" example:"1"`
	Name       string `json:"Name" example:"TrackName"`
	Preview    string `json:"Preview" example:"TrackPreview"`
	Content    string `json:"Content" example:"TrackContent"`
	ArtistName string `json:"ArtistName" example:"ArtiName"`
	Duration   uint32 `json:"Duration" example:"0"`
	IsLiked    bool   `json:"IsLiked" example:"false"`
}

type UseCase interface {
	GetUserLikedTracks(userId string) ([]Response, error)
	Listen(trackId uint64) error
	Like(userId string, trackId uint64) error
	IsLike(userId string, trackId uint64) (bool, error)
	Unlike(userId string, trackId uint64) error
}

type Repository interface {
	GetByUser(userId string) ([]Response, error)
	GetByPlaylist(playlistId uint64) ([]Response, error)
	GetByAlbum(albumId uint64) ([]Response, error)
	GetByArtist(artistId uint64) ([]Response, error)
	CreateLike(userId string, trackId uint64) error
	CheckLike(userId string, trackId uint64) (bool, error)
	DeleteLike(userId string, trackId uint64) error
	AddListen(trackId uint64) error
}

var (
	ErrTrackAlreadyExist = errors.New("track already exist")
	ErrTrackDoesNotExist = errors.New("track does not exist")
	ErrNoTracks          = errors.New("track already exist")
)
