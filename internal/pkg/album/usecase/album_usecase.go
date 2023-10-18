package album_usecase

import (
	"main/internal/pkg/album"
	"main/internal/pkg/artist"
	"main/internal/pkg/track"
)

type Default struct {
	repoArtist artist.Repository
	repoTrack  track.Repository
	repoAlbum  album.Repository
}

func NewDefault(artistRepo artist.Repository, repoTrack track.Repository, repoAlbum album.Repository) Default {
	return Default{
		repoArtist: artistRepo,
		repoTrack:  repoTrack,
		repoAlbum:  repoAlbum,
	}
}

func (r *Default) GetAlbum(albumId uint64) (album.Response, error) {
	var result album.Response

	base, err := r.repoAlbum.Get(albumId)
	if err != nil {
		return result, err
	}

	result.Id = base.Id
	result.Name = base.Name
	result.Preview = base.Preview

	art, err := r.repoArtist.GetByAlbumId(albumId)
	if err != nil {
		return result, err
	}

	result.ArtistId = art.Id
	result.ArtistName = art.Name

	tracks, err := r.repoTrack.GetByAlbum(albumId)
	if err != nil {
		return result, err
	}
	result.Tracks = tracks

	return result, nil
}

func (r *Default) GetRandom() ([]album.Response, error) {
	albums, err := r.repoAlbum.GetRandom(album.LimitForMainPage)
	if err != nil {
		return nil, err
	}
	return r.formResponse(albums)
}

func (r *Default) GetMostLiked() ([]album.Response, error) {
	albums, err := r.repoAlbum.GetByLikeCount(album.LimitForMainPage)
	if err != nil {
		return nil, err
	}
	return r.formResponse(albums)
}

func (r *Default) GetPopular() ([]album.Response, error) {
	albums, err := r.repoAlbum.GetByListenCount(album.LimitForMainPage)
	if err != nil {
		return nil, err
	}
	return r.formResponse(albums)
}

func (r *Default) GetNew() ([]album.Response, error) {
	albums, err := r.repoAlbum.GetByReleaseDate(album.LimitForMainPage)
	if err != nil {
		return nil, err
	}
	return r.formResponse(albums)
}

func (r *Default) formResponse(albumBase []album.Base) ([]album.Response, error) {
	result := make([]album.Response, 0)
	for _, base := range albumBase {
		var a album.Response
		art, err := r.repoArtist.GetByAlbumId(base.Id)
		if err != nil {
			return result, nil
		}

		a.Id = base.Id
		a.Name = base.Name
		a.Preview = base.Preview
		a.ArtistId = art.Id
		a.ArtistName = art.Name

		result = append(result, a)
	}

	return result, nil
}

func (r *Default) Like(userId string, albumId uint64) error {
	err := r.repoAlbum.CreateLike(userId, albumId)
	if err != nil {
		return err
	}
	return nil
}
