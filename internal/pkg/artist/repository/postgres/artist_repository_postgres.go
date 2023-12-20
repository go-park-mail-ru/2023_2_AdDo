package artist_repository

import (
	"context"
	"github.com/sirupsen/logrus"
	postgres "main/internal/common/pgxiface"
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
	if err := p.Pool.QueryRow(context.Background(), query, artistId).Scan(&result.Id, &result.Name, &result.Avatar); err != nil {
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

func (p *Postgres) getWithQuery(_ context.Context, query string, args ...any) ([]artist.Base, error) {
	p.logger.Infoln("Artist Repo getWithQuery entered")

	result := make([]artist.Base, 0)

	rows, err := p.Pool.Query(context.Background(), query, args...)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"err":   err,
			"query": query,
		}).Errorln("Getting an artist by track id query completed with error")
		return nil, err
	}
	defer rows.Close()
	p.logger.Infoln("Got artist")

	for rows.Next() {
		var a artist.Base

		if err := rows.Scan(&a.Id, &a.Name, &a.Avatar); err != nil {
			return nil, err
		}

		result = append(result, a)
	}
	p.logger.Infoln("Formed response")

	return result, nil
}

func (p *Postgres) GetByTrackId(trackId uint64) ([]artist.Base, error) {
	p.logger.Infoln("Artist Repo GetByTrackId entered")

	query := "select artist.id, name, avatar from artist join artist_track on artist.id = artist_track.artist_id where artist_track.track_id = $1"
	return p.getWithQuery(context.Background(), query, trackId)
}

func (p *Postgres) Search(text string) ([]artist.Base, error) {
	p.logger.Infoln("Artist Repo Search entered")
	query := `select artist.id, artist.name, artist.avatar from artist 
                                             where to_tsvector('russian', artist.name) @@ plainto_tsquery('russian', $1)
											 or lower(artist.name) like lower($2)
											 or similarity(artist.name, $1) > 0.5
											 limit 10`
	return p.getWithQuery(context.Background(), query, text, "%"+text+"%")
}

func (p *Postgres) GetByAlbumId(albumId uint64) (artist.Base, error) {
	p.logger.Infoln("Artist Repo GetByAlbumId entered")

	var result artist.Base

	query := "select artist.id, artist.name, artist.avatar from artist join artist_album on artist.id = artist_album.artist_id where artist_album.album_id = $1"
	if err := p.Pool.QueryRow(context.Background(), query, albumId).Scan(&result.Id, &result.Name, &result.Avatar); err != nil {
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

func (p *Postgres) GetByUserId(userId string) ([]artist.Base, error) {
	p.logger.Infoln("Artist Repo Get By User Id entered")

	query := "select artist.id, artist.name, artist.avatar from artist join profile_artist on artist.id = profile_artist.artist_id where profile_artist.profile_id = $1"
	return p.getWithQuery(context.Background(), query, userId)
}

func (p *Postgres) CreateLike(userId string, artistId uint64) error {
	p.logger.Infoln("Artist Repo CreateLike entered")

	query := "insert into profile_artist (profile_id, artist_id) values ($1, $2)"
	if _, err := p.Pool.Exec(context.Background(), query, userId, artistId); err != nil {
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

	var counter int
	query := "select count(*) from profile_artist where profile_id = $1 and artist_id = $2"
	if err := p.Pool.QueryRow(context.Background(), query, userId, artistId).Scan(&counter); err != nil {
		p.logger.Errorln(err)
		return false, err
	}
	p.logger.Infoln("like checked")

	if counter == 0 {
		return false, nil
	}

	return true, nil
}

func (p *Postgres) DeleteLike(userId string, artistId uint64) error {
	p.logger.Infoln("Artist Repo DeleteLike entered")

	query := "delete from profile_artist where profile_id = $1 and artist_id = $2"
	if _, err := p.Pool.Exec(context.Background(), query, userId, artistId); err != nil {
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
