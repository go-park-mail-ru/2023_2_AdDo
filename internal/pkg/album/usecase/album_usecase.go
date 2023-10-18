package album_usecase

import (
	"github.com/sirupsen/logrus"
	"main/internal/pkg/album"
	"main/internal/pkg/artist"
	"main/internal/pkg/track"
)

type Default struct {
	repoArtist artist.Repository
	repoTrack  track.Repository
	repoAlbum  album.Repository
	logger     *logrus.Logger
}

func NewDefault(artistRepo artist.Repository, repoTrack track.Repository, repoAlbum album.Repository, logger *logrus.Logger) Default {
	return Default{
		repoArtist: artistRepo,
		repoTrack:  repoTrack,
		repoAlbum:  repoAlbum,
		logger:     logger,
	}
}

func (r *Default) GetAlbum(albumId uint64) (album.Response, error) {
	r.logger.Infoln("Album UseCase GetAlbum entered")

	var result album.Response

	base, err := r.repoAlbum.Get(albumId)
	if err != nil {
		return result, err
	}
	r.logger.Infoln("Got album Base")

	result.Id = base.Id
	result.Name = base.Name
	result.Preview = base.Preview

	art, err := r.repoArtist.GetByAlbumId(albumId)
	if err != nil {
		return result, err
	}
	r.logger.Infoln("Got Artist by album Id ", art)

	result.ArtistId = art.Id
	result.ArtistName = art.Name

	tracks, err := r.repoTrack.GetByAlbum(albumId)
	if err != nil {
		return result, err
	}
	r.logger.Infoln("Got tracks by album id ", tracks)

	result.Tracks = tracks

	return result, nil
}

func (r *Default) GetRandom() ([]album.Response, error) {
	r.logger.Infoln("Album UseCase GetRandom entered")

	albums, err := r.repoAlbum.GetRandom(album.LimitForMainPage)
	if err != nil {
		return nil, err
	}
	r.logger.Infoln("Got random albums")

	return r.formResponse(albums)
}

func (r *Default) GetMostLiked() ([]album.Response, error) {
	r.logger.Infoln("Album UseCase GetMostLiked entered")

	albums, err := r.repoAlbum.GetByLikeCount(album.LimitForMainPage)
	if err != nil {
		return nil, err
	}
	r.logger.Infoln("Got album")

	return r.formResponse(albums)
}

func (r *Default) GetPopular() ([]album.Response, error) {
	r.logger.Infoln("Album UseCase GetPopular entered")

	albums, err := r.repoAlbum.GetByListenCount(album.LimitForMainPage)
	if err != nil {
		return nil, err
	}
	r.logger.Infoln("Got albums by Listen count")

	return r.formResponse(albums)
}

func (r *Default) GetNew() ([]album.Response, error) {
	r.logger.Infoln("Album UseCase GetNew entered")

	albums, err := r.repoAlbum.GetByReleaseDate(album.LimitForMainPage)
	if err != nil {
		return nil, err
	}
	r.logger.Infoln("Got new albums")

	return r.formResponse(albums)
}

func (r *Default) formResponse(albumBase []album.Base) ([]album.Response, error) {
	r.logger.Infoln("Album UseCase fromResponse entered")

	result := make([]album.Response, 0)
	for _, base := range albumBase {
		var a album.Response
		art, err := r.repoArtist.GetByAlbumId(base.Id)
		if err != nil {
			return result, nil
		}
		r.logger.Infoln("artist founded")

		a.Id = base.Id
		a.Name = base.Name
		a.Preview = base.Preview
		a.ArtistId = art.Id
		a.ArtistName = art.Name

		result = append(result, a)
	}
	r.logger.Infoln("response formed")

	return result, nil
}

func (r *Default) Like(userId string, albumId uint64) error {
	r.logger.Infoln("Album UseCase Like entered")

	err := r.repoAlbum.CreateLike(userId, albumId)
	if err != nil {
		return err
	}
	r.logger.Infoln("Like created")

	return nil
}
