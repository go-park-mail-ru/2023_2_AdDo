package artist_usecase

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

func (repo *Default) GetArtistInfo(artistId uint64) (artist.Response, error) {
	repo.logger.Infoln("Artist UseCase GetArtistInfo entered")

	var result artist.Response

	artistBase, err := repo.repoArtist.Get(artistId)
	if err != nil {
		return result, err
	}
	repo.logger.Infoln("Got artist base")

	result.Id = artistBase.Id
	result.Name = artistBase.Name
	result.Avatar = artistBase.Avatar

	albums, err := repo.repoAlbum.GetByArtistId(artistId)
	if err != nil {
		return result, err
	}
	repo.logger.Infoln("Got artist albums")

	result.Albums = albums

	tracks, err := repo.repoTrack.GetByArtist(artistId)
	if err != nil {
		return result, err
	}
	repo.logger.Infoln("Got artist tracks")

	result.Tracks = tracks

	return result, nil
}
