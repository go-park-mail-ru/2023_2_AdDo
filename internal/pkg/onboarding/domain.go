package onboarding

import artist "main/internal/pkg/artist"

type GenreBase struct {
	Id      uint64
	Name    string
	Preview string
}

type GenreOnboarding struct {
	Genres []GenreBase
	UserId string
}

type ArtistOnboarding struct {
	Artists []artist.Base
	UserId  string
}

type Repository interface {
	GetGenres() ([]GenreBase, error)
	GetArtists() ([]artist.Base, error)
	SetUserGenres(userId string, genres []GenreBase) error
	SetUserArtists(userId string, artists []artist.Base) error
}

type UseCase interface {
	GetGenres() ([]GenreBase, error)
	GetArtists() ([]artist.Base, error)
	SetUserGenres(onboarding GenreOnboarding) error
	SetUserArtists(onboarding ArtistOnboarding) error
}
