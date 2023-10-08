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
	Id        uint64
	Name      string
	Artist    []uint64
	Album     []uint64
	Preview   string
	Content   string
	PlayCount uint64
}

type Response struct {
	Id        uint64
	Name      string
	Artist    []artist.Response
	Album     []album.Response
	Preview   string
	Content   string
	PlayCount uint64
}

type UseCase interface {
	//Add(track Track) (uint64, error)
	GetAll() ([]Response, error)
	GetPopular(requiredNumOfTracks int) ([]Response, error)
	GetByAlbum(albumId uint64) ([]Response, error)
	GetByArtist(artistId uint64) ([]Response, error)
	GetByPlaylist(playlistId uint64) ([]Response, error)
	//GetFavourite(userId uint64) ([]Response, error)
}

type Repository interface {
	GetAll() ([]Response, error)
	GetTrackIdsByAlbum(albumId uint64) ([]uint64, error)
	GetTrackIdsByArtist(artistId uint64) ([]uint64, error)
	GetTrackIdsByPlaylist(playlistId uint64) ([]uint64, error)
	GetByTrackId(trackId uint64) (Response, error)
}

var (
	ErrTrackAlreadyExist = errors.New("track already exist")
	ErrTrackDoesNotExist = errors.New("track does not exist")
	ErrNoTracks          = errors.New("track already exist")
)
