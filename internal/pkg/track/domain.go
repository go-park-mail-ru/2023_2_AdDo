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
	Id          uint64
	Name        string
	Artist      []uint64
	Album       []uint64
	Preview     string
	Content     string
	PlayCount   uint64
	ReleaseDate string
}

type Response struct {
	Id          uint64            `json:"Id" example:"1"`
	Name        string            `json:"Name" example:"TrackName"`
	Artist      []artist.Response `json:"Artist"`
	Album       []album.Response  `json:"Album"`
	Preview     string            `json:"Preview" example:"TrackPreview"`
	Content     string            `json:"Content" example:"TrackContent"`
	PlayCount   uint64
	ReleaseDate string
}

type UseCase interface {
	GetAll() ([]Response, error)
	GetPopular(limit uint32) ([]Response, error)
	GetLatest(limit uint32) ([]Response, error)
	GetByAlbum(albumId uint64) ([]Response, error)
	GetByArtist(artistId uint64) ([]Response, error)
	GetByPlaylist(playlistId uint64) ([]Response, error)
}

type Repository interface {
	GetAll() ([]Response, error)
	GetPopular(limit uint32) ([]Response, error)
	GetLatest(limit uint32) ([]Response, error)
	GetByAlbum(albumId uint64) ([]Response, error)
	GetByArtist(artistId uint64) ([]Response, error)
	GetByPlaylist(playlistId uint64) ([]Response, error)
}

var (
	ErrTrackAlreadyExist = errors.New("track already exist")
	ErrTrackDoesNotExist = errors.New("track does not exist")
	ErrNoTracks          = errors.New("track already exist")
)
