package artist_repository

import (
	"context"
	"github.com/sirupsen/logrus"
	"main/internal/common/pgxiface"
	"main/internal/pkg/artist"
)

type Postgres struct {
	Pool   postgres.PgxIFace
	logger *logrus.Logger
}

func NewPostgres(pool postgres.PgxIFace, logger *logrus.Logger) Postgres {
	return Postgres{
		Pool:   pool,
		logger: logger,
	}
}

func (repo *Postgres) Get(artistId uint64) (artist.Base, error) {
	repo.logger.Infoln("Artist Repo Get entered")

	var result artist.Base

	query := "select artist.id, name, image from artist where artist.id = $1"
	err := repo.Pool.QueryRow(context.Background(), query, artistId).Scan(&result.Id, &result.Name, &result.Avatar)
	if err != nil {
		repo.logger.WithFields(logrus.Fields{
			"err":       err,
			"artist id": artistId,
			"query":     query,
		}).Errorln("Getting an artist query completed with error")
		return result, err
	}
	repo.logger.Infoln("Got artist by id")

	return result, nil
}

func (repo *Postgres) GetByTrackId(trackId uint64) ([]artist.Base, error) {
	repo.logger.Infoln("Artist Repo GetByTrackId entered")

	result := make([]artist.Base, 0)

	query := "select artist.id, name, image from artist join artist_track on artist.id = artist_track.artist_id where artist_track.track_id = $1"
	rows, err := repo.Pool.Query(context.Background(), query, trackId)
	if err != nil {
		repo.logger.WithFields(logrus.Fields{
			"err":      err,
			"track id": trackId,
			"query":    query,
		}).Errorln("Getting an artist by track id query completed with error")
		return nil, err
	}
	defer rows.Close()
	repo.logger.Infoln("Got artist")

	for rows.Next() {
		var a artist.Base

		err := rows.Scan(&a.Id, &a.Name, &a.Avatar)
		if err != nil {
			return nil, err
		}

		result = append(result, a)
	}
	repo.logger.Infoln("Formed response")

	return result, nil
}

func (repo *Postgres) GetByAlbumId(albumId uint64) (artist.Base, error) {
	repo.logger.Infoln("Artist Repo GetByAlbumId entered")

	var result artist.Base

	query := "select artist.id, artist.name, image from artist join album on artist.id = album.artist_id where album.id = $1"
	err := repo.Pool.QueryRow(context.Background(), query, albumId).Scan(&result.Id, &result.Name, &result.Avatar)
	if err != nil {
		repo.logger.WithFields(logrus.Fields{
			"err":      err,
			"album id": albumId,
			"query":    query,
		}).Errorln("Getting an artist by album id query completed with error")
		return artist.Base{}, err
	}
	repo.logger.Infoln("Got artist by album id")

	return result, nil
}
