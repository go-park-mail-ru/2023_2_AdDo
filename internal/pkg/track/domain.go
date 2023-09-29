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
	Id      uint64
	Name    string
	Artist  []artist.Response
	Album   []album.Response
	Preview string
	Content string
}

type UseCase interface {
	Add(track Track) (uint64, error)
	GetAll() ([]Response, error)
	GetPopular() ([]Response, error)
	GetFavourite(userId uint64) ([]Response, error)
}

type Repository interface {
	Create(track Track) (uint64, error)
	GetById(id uint64) (Track, error)
	GetAll() ([]Response, error)
	GetByUserId(userId uint64) ([]Response, error)
	//GetMostLiked() ([]Track, error)
	//GetByArtistId(artistId uint64) ([]Track, error)
	//GetByPlaylistId(playlistId uint64) ([]Track, error)
	//GetByAlbumId(albumId uint64) ([]Track, error)
}

var (
	ErrTrackAlreadyExist = errors.New("track already exist")
	ErrTrackDoesNotExist = errors.New("track does not exist")
	ErrNoTracks          = errors.New("track already exist")
)
