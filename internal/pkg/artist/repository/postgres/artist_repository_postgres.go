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

func (p *Postgres) Get(artistId uint64) (artist.Base, error) {
	p.logger.Infoln("Artist Repo Get entered")

	var result artist.Base

	query := "select artist.id, name, avatar from artist where artist.id = $1"
	err := p.Pool.QueryRow(context.Background(), query, artistId).Scan(&result.Id, &result.Name, &result.Avatar)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"err":       err,
			"artist id": artistId,
			"query":     query,
		}).Errorln("Getting an artist query completed with error")
		return result, err
	}
	p.logger.Infoln("Got artist by id")

	return result, nil
}

func (p *Postgres) GetByTrackId(trackId uint64) ([]artist.Base, error) {
	p.logger.Infoln("Artist Repo GetByTrackId entered")

	result := make([]artist.Base, 0)

	query := "select artist.id, name, avatar from artist join artist_track on artist.id = artist_track.artist_id where artist_track.track_id = $1"
	rows, err := p.Pool.Query(context.Background(), query, trackId)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"err":      err,
			"track id": trackId,
			"query":    query,
		}).Errorln("Getting an artist by track id query completed with error")
		return nil, err
	}
	defer rows.Close()
	p.logger.Infoln("Got artist")

	for rows.Next() {
		var a artist.Base

		err := rows.Scan(&a.Id, &a.Name, &a.Avatar)
		if err != nil {
			return nil, err
		}

		result = append(result, a)
	}
	p.logger.Infoln("Formed response")

	return result, nil
}

func (p *Postgres) GetByAlbumId(albumId uint64) (artist.Base, error) {
	p.logger.Infoln("Artist Repo GetByAlbumId entered")

	var result artist.Base

	query := "select artist.id, artist.name, avatar from artist join album on artist.id = album.artist_id where album.id = $1"
	err := p.Pool.QueryRow(context.Background(), query, albumId).Scan(&result.Id, &result.Name, &result.Avatar)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"err":      err,
			"album id": albumId,
			"query":    query,
		}).Errorln("Getting an artist by album id query completed with error")
		return artist.Base{}, err
	}
	p.logger.Infoln("Got artist by album id")

	return result, nil
}

func (p *Postgres) CreateLike(userId string, artistId uint64) error {
	p.logger.Infoln("Artist Repo CreateLike entered")

	query := "insert into profile_artist (profile_id, artist_id) values ($1, $2)"
	_, err := p.Pool.Exec(context.Background(), query, userId, artistId)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"err":       err,
			"artist id": artistId,
			"query":     query,
		}).Errorln("creating like error")
		return err
	}
	p.logger.Infoln("Like Created")

	return nil
}

func (p *Postgres) CheckLike(userId string, artistId uint64) (bool, error) {
	p.logger.Infoln("Artist Repo CheckLike entered")

	query := "select * from profile_artist where profile_id = $1 and artist_id = $2"
	err := p.Pool.QueryRow(context.Background(), query, userId, artistId).Scan()
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"err":       err,
			"artist Id": artistId,
			"query":     query,
		}).Errorln("checking like error")
		return false, err
	}
	p.logger.Infoln("Like checked")

	return true, nil
}

func (p *Postgres) DeleteLike(userId string, artistId uint64) error {
	p.logger.Infoln("Artist Repo DeleteLike entered")

	query := "delete from profile_artist where profile_id = $1 and artist_id = $2"
	_, err := p.Pool.Exec(context.Background(), query, userId, artistId)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"err":       err,
			"artist Id": artistId,
			"query":     query,
		}).Errorln("delete like error")
		return err
	}
	p.logger.Infoln("Like deleted")

	return nil
}
