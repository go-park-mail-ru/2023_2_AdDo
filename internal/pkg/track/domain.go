package track

import (
	"errors"
)

type Id struct {
	Id uint64 `json:"Id" example:"1"`
}

type Response struct {
	Id      uint64 `json:"Id" example:"1"`
	Name    string `json:"Name" example:"TrackName"`
	Preview string `json:"Preview" example:"TrackPreview"`
	Content string `json:"Content" example:"TrackContent"`
}

type UseCase interface {
	//GetAll() ([]Response, error)
	//GetPopular(limit uint32) ([]Response, error)
	//GetLatest(limit uint32) ([]Response, error)
	//GetByAlbum(albumId uint64) ([]Response, error)
	//GetByPlaylist(playlistId uint64) ([]Response, error)
	//GetByArtist(artistId uint64) ([]Response, error)
	Listen(trackId uint64) error
	Like(userId, trackId uint64) error
}

type Repository interface {
	//GetAll() ([]Response, error)
	//GetPopular(limit uint32) ([]Response, error)
	//GetLatest(limit uint32) ([]Response, error)
	//GetByPlaylist(playlistId uint64) ([]Response, error)
	GetByAlbum(albumId uint64) ([]Response, error)
	GetByArtist(artistId uint64) ([]Response, error)
	CreateLike(userId, trackId uint64) error
	AddListen(trackId uint64) error
}

var (
	ErrTrackAlreadyExist = errors.New("track already exist")
	ErrTrackDoesNotExist = errors.New("track does not exist")
	ErrNoTracks          = errors.New("track already exist")
)
