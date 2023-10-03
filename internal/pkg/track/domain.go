package track

import (
	"errors"
	"main/internal/pkg/album"
	"main/internal/pkg/artist"
)

type ArtistTrackResponse struct {
	Id      uint64
	Name    string
	Preview string
}

type ArtistAlbumResponse struct {
	Id      uint64
	Name    string
	Preview string
}

type Track struct {
	Id      uint64
	Name    string
	Artist  []uint64
	Album   []uint64
	Preview string
	Content string
}

type Response struct {
	Id      uint64 `json:"Id" example:"1"`
	Name    string `json:"Name" example:"TrackName"`
	Artist  []artist.Response `json:"Artist"`
	Album   []album.Response `json:"Album"`
	Preview string `json:"Preview" example:"TrackPreview"`
	Content string `json:"Content" example:"TrackContent"`
}

type UseCase interface {
	//Add(track Track) (uint64, error)
	GetAll() ([]Response, error)
	//GetPopular() ([]Response, error)
	//GetFavourite(userId uint64) ([]Response, error)
}

type Repository interface {
	GetAll() ([]Response, error)
}

var (
	ErrTrackAlreadyExist = errors.New("track already exist")
	ErrTrackDoesNotExist = errors.New("track does not exist")
	ErrNoTracks          = errors.New("track already exist")
)
