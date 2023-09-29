package track

import "errors"

type Track struct {
	Id          uint64
	Name        string
	IsSong      bool
	FKArtistId  uint64
	FKAlbumId   uint64
	ImagePath   string
	ContentPath string
}

type Usecase interface {
	Add(track Track) (uint64, error)
	GetAll() ([]Track, error)
	GetPopular() ([]Track, error)
	GetFavourite(userId uint64) ([]Track, error)
}

type Repository interface {
	Create(track Track) (uint64, error)
	GetById(id uint64) (Track, error)
	GetAll() ([]Track, error)
	//GetMostLiked() ([]Track, error)
	//GetByUserId(userId uint64) ([]Track, error)
	//GetByArtistId(artistId uint64) ([]Track, error)
	//GetByPlaylistId(playlistId uint64) ([]Track, error)
	//GetByAlbumId(albumId uint64) ([]Track, error)
}

var (
	ErrTrackAlreadyExist = errors.New("track already exist")
	ErrTrackDoesNotExist = errors.New("track does not exist")
	ErrNoTracks          = errors.New("track already exist")
)
