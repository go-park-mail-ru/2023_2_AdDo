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

func (useCase *Default) addArtistAndAlbum(tracks []track.Response) ([]track.Response, error) {
	for index, t := range tracks {
		artists, err := useCase.repoArtist.GetByTrackId(t.Id)
		if err != nil {
			return nil, err
		}
		tracks[index].Artist = artists

		albums, err := useCase.repoAlbum.GetByTrackId(t.Id)
		if err != nil {
			return nil, err
		}
		tracks[index].Album = albums
	}
	return tracks, nil
}

func (useCase *Default) GetAll() ([]track.Response, error) {
	tracks, err := useCase.repoTrack.GetAll()
	if err != nil {
		return nil, track.ErrNoTracks
	}
	return useCase.addArtistAndAlbum(tracks)
}

func (useCase *Default) GetPopular(limit uint32) ([]track.Response, error) {
	tracks, err := useCase.repoTrack.GetPopular(limit)
	if err != nil {
		return nil, track.ErrNoTracks
	}
	return useCase.addArtistAndAlbum(tracks)
}

func (useCase *Default) GetLatest(limit uint32) ([]track.Response, error) {
	tracks, err := useCase.repoTrack.GetLatest(limit)
	if err != nil {
		return nil, track.ErrNoTracks
	}
	return useCase.addArtistAndAlbum(tracks)
}

func (useCase *Default) GetByAlbum(albumId uint64) ([]track.Response, error) {
	tracks, err := useCase.repoTrack.GetByAlbum(albumId)
	if err != nil {
		return nil, track.ErrNoTracks
	}
	return useCase.addArtistAndAlbum(tracks)
}

func (useCase *Default) GetByArtist(artistId uint64) ([]track.Response, error) {
	tracks, err := useCase.repoTrack.GetByArtist(artistId)
	if err != nil {
		return nil, track.ErrNoTracks
	}
	return useCase.addArtistAndAlbum(tracks)
}

func (useCase *Default) GetByPlaylist(artistId uint64) ([]track.Response, error) {
	tracks, err := useCase.repoTrack.GetByPlaylist(artistId)
	if err != nil {
		return nil, track.ErrNoTracks
	}
	return useCase.addArtistAndAlbum(tracks)
}

func (useCase *Default) Listen(trackId uint64) error {
	err := useCase.repoTrack.AddListen(trackId)
	if err != nil {
		return err
	}
	return nil
}

func (useCase *Default) Like(userId, trackId uint64) error {
	err := useCase.repoTrack.CreateLike(userId, trackId)
	if err != nil {
		return err
	}
	return nil
}

//func (useCase *Default) GetFavourite(userId uint64) ([]track.Base, error) {
//	tracks, err := useCase.repoTrack.GetByUserId(userId)
//	if err != nil {
//		return nil, track.ErrNoTracks
//	}
//
//	for _, t := range tracks {
//		artists, err := useCase.repoArtist.GetByTrackId(t.Id)
//		if err != nil {
//			return nil, err
//		}
//		t.Artist = artists
//
//		albums, err := useCase.repoAlbum.GetByTrackId(t.Id)
//		if err != nil {
//			return nil, err
//		}
//		t.Album = albums
//	}
//
//	return tracks, nil
//}
