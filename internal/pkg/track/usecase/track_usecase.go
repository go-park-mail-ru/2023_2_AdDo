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

func (useCase *Default) getTracks(trackIds []uint64) ([]track.Response, error) {
	tracks := make([]track.Response, 0)
	for _, trackId := range trackIds {
		t, err := useCase.repoTrack.GetByTrackId(trackId)
		if err != nil {
			return nil, err
		}

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

		tracks = append(tracks, t)
	}
	return tracks, nil
}

func (useCase *Default) GetByAlbum(albumId uint64) ([]track.Response, error) {
	trackIds, err := useCase.repoTrack.GetTrackIdsByAlbum(albumId)
	if err != nil {
		return nil, err
	}
	return useCase.getTracks(trackIds)
}

func (useCase *Default) GetByArtist(artistId uint64) ([]track.Response, error) {
	trackIds, err := useCase.repoTrack.GetTrackIdsByArtist(artistId)
	if err != nil {
		return nil, err
	}
	return useCase.getTracks(trackIds)
}

func (useCase *Default) GetByPlaylist(artistId uint64) ([]track.Response, error) {
	trackIds, err := useCase.repoTrack.GetTrackIdsByPlaylist(artistId)
	if err != nil {
		return nil, err
	}
	return useCase.getTracks(trackIds)
}

//func (useCase *Default) GetFavourite(userId uint64) ([]track.Response, error) {
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
