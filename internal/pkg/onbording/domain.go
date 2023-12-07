package onbording

import artist "main/internal/pkg/artist"

type GenreBase struct {
	Id      uint64
	Name    string
	Preview string
}

type Repository interface {
	GetGenres() ([]GenreBase, error)
	GetArtists() []artist.Base
	SetUserGenres(userId string, genres []GenreBase)
	SetUserArtists(userId string, artists []artist.Base)
}

type UseCase interface {
	GetGenres() ([]GenreBase, error)
	GetArtists() []artist.Base
	SetUserGenres(userId string, genres []GenreBase)
	SetUserArtists(userId string, artists []artist.Base)
}
