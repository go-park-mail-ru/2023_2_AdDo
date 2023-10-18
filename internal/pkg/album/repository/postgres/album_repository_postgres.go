package album_repository

import (
	"context"
	"github.com/sirupsen/logrus"
	"main/internal/pkg/album"
	postgres "main/internal/pkg/common/pgxiface"
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

func (p Postgres) GetByTrackId(trackId uint64) ([]album.Base, error) {
	p.logger.Infoln("Album Repo GetByTrackId entered")
	query := "select album.id, name, preview from album join album_track on album.id = album_track.album_id where album_track.track_id = $1"
	return p.getWithQuery(context.Background(), query, trackId)
}

func (p Postgres) GetByArtistId(artistId uint64) ([]album.Base, error) {
	p.logger.Infoln("Album Repo GetByArtistId entered")
	query := "select id, name, preview from album where artist_id = $1"
	return p.getWithQuery(context.Background(), query, artistId)
}

func (p Postgres) GetByReleaseDate(limit uint32) ([]album.Base, error) {
	p.logger.Infoln("Album Repo GetByReleaseDate entered")
	query := "select id, name, preview from album order by release_date desc limit $1"
	return p.getWithQuery(context.Background(), query, limit)
}

func (p Postgres) GetRandom(limit uint32) ([]album.Base, error) {
	p.logger.Infoln("Album Repo GetRandom entered")
	query := "select id, name, preview from album limit $1"
	return p.getWithQuery(context.Background(), query, limit)
}

func (p Postgres) GetByListenCount(limit uint32) ([]album.Base, error) {
	p.logger.Infoln("Album Repo GetByListenCount entered")
	query := "select id, name, preview from album order by play_count desc limit $1"
	return p.getWithQuery(context.Background(), query, limit)
}

func (p Postgres) GetByLikeCount(limit uint32) ([]album.Base, error) {
	p.logger.Infoln("Album Repo GetByLikeCount entered")
	query := "select id, name, preview from album order by rating desc limit $1"
	return p.getWithQuery(context.Background(), query, limit)
}

func (p Postgres) Get(albumId uint64) (album.Base, error) {
	p.logger.Infoln("Album Repo Get entered")

	var result album.Base

	query := "select id, name, preview from album where id = $1"
	err := p.Pool.QueryRow(context.Background(), query, albumId).Scan(&result.Id, &result.Name, &result.Preview)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"err":      err,
			"album id": albumId,
			"query":    query,
		}).Errorln("Getting an album completed with error")
		return result, err
	}
	p.logger.Infoln("Got album base")

	return result, nil
}

func (p Postgres) getWithQuery(ctx context.Context, query string, args ...any) ([]album.Base, error) {
	p.logger.Infoln("Album Repo getWithQuery entered")

	result := make([]album.Base, 0)

	rows, err := p.Pool.Query(ctx, query, args...)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"err":   err,
			"args":  args,
			"query": query,
		}).Errorln("Getting an album with custom query completed with error")
		return result, err
	}
	defer rows.Close()
	p.logger.Infoln("Got rows from db")

	for rows.Next() {
		var base album.Base

		err = rows.Scan(&base.Id, &base.Name, &base.Preview)
		if err != nil {
			p.logger.WithFields(logrus.Fields{
				"err":      err,
				"album id": base.Id,
				"query":    query,
			}).Errorln("error rows scanning")
			return nil, err
		}

		result = append(result, base)
	}
	p.logger.Infoln("Got albums success")

	return result, nil
}

func (p Postgres) CreateLike(userId string, albumId uint64) error {
	p.logger.Infoln("Album Repo CreateLike entered")

	query := "insert into profile_album (profile_id, album_id) values ($1, $2)"
	_, err := p.Pool.Exec(context.Background(), query, userId, albumId)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"err":      err,
			"album id": albumId,
			"query":    query,
		}).Errorln("creating like error")
		return err
	}
	p.logger.Infoln("Like Created")

	return nil
}
