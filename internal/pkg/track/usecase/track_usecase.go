package track_usecase

import (
	"main/internal/pkg/album"
	"main/internal/pkg/artist"
	"main/internal/pkg/track"
)

type Default struct {
	repoTrack  track.Repository
	repoArtist artist.Repository
	repoAlbum  album.Repository
}

func NewDefault(trackRepo track.Repository, artistRepo artist.Repository, albumRepo album.Repository) Default {
	return Default{
		repoTrack:  trackRepo,
		repoArtist: artistRepo,
		repoAlbum:  albumRepo,
	}
}

func (useCase *Default) Add(t track.Track) (uint64, error) {
	id, err := useCase.repoTrack.Create(t)
	if err != nil {
		return 0, track.ErrTrackAlreadyExist
	}
	return id, nil
}

func (useCase *Default) GetAll() ([]track.Response, error) {
	tracks, err := useCase.repoTrack.GetAll()
	if err != nil {
		return nil, track.ErrNoTracks
	}

	for _, t := range tracks {
		artists, err := useCase.repoArtist.GetByTrackId(t.Id)
		if err != nil {
			return nil, err
		}
		t.Artist = artists

		albums, err := useCase.repoAlbum.GetByTrackId(t.Id)
		if err != nil {
			return nil, err
		}
		t.Album = albums
	}

	return tracks, nil
}

func (useCase *Default) GetFavourite(userId uint64) ([]track.Response, error) {
	tracks, err := useCase.repoTrack.GetByUserId(userId)
	if err != nil {
		return nil, track.ErrNoTracks
	}

	for _, t := range tracks {
		artists, err := useCase.repoArtist.GetByTrackId(t.Id)
		if err != nil {
			return nil, err
		}
		t.Artist = artists

		albums, err := useCase.repoAlbum.GetByTrackId(t.Id)
		if err != nil {
			return nil, err
		}
		t.Album = albums
	}

	return tracks, nil
}

func (useCase *Default) GetPopular() ([]track.Response, error) {
	return nil, nil
}
