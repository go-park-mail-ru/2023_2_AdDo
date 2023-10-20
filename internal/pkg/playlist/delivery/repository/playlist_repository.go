package playlist_repository

import (
	"context"
	"github.com/sirupsen/logrus"
	postgres "main/internal/pkg/common/pgxiface"
	"main/internal/pkg/playlist"
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

func (p *Postgres) Create(ctx context.Context, base playlist.Base) error {
	p.logger.Infoln("Playlist Repo Create entered")

	query := "insert into playlist (name, creator_id )values ($1, $2)"
	_, err := p.Pool.Exec(ctx, query, base.Name, base.AuthorId)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"error":         err,
			"playlist base": base,
			"query":         query,
		}).Errorln("error while creating playlist row")
		return err
	}

	return nil
}

func (p *Postgres) Get(ctx context.Context, playlistId uint64) (playlist.Base, error) {
	p.logger.Infoln("Playlist Repo Get entered")

	var result playlist.Base
	query := "select id, name, creator_id, preview from playlist where id = $1"
	err := p.Pool.QueryRow(ctx, query, playlistId).Scan(&result.Id, &result.Name, &result.AuthorId, &result.Preview)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"error":       err,
			"playlist id": playlistId,
			"query":       query,
		}).Errorln("error while selecting playlist row by id")
		return playlist.Base{}, err
	}

	return result, nil
}

func (p *Postgres) AddTrack(ctx context.Context, playlistId, trackId uint64) error {
	p.logger.Infoln("Playlist Repo AddTrack entered")

	query := "insert into playlist_track (playlist_id, track_id) values ($1, $2)"
	_, err := p.Pool.Exec(ctx, query, playlistId, trackId)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"error":       err,
			"playlist id": playlistId,
			"track id":    trackId,
			"query":       query,
		}).Errorln("error while adding track into playlist")
	}

	return nil
}

func (p *Postgres) UpdateImage(ctx context.Context, playlistId uint64, image string) error {
	p.logger.Infoln("Playlist Repo UpdateImage entered")

	query := "update playlist set preview = $1 where id = $2"
	_, err := p.Pool.Exec(ctx, query, image, playlistId)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"error":          err,
			"image url path": image,
			"playlist id":    playlistId,
			"query":          query,
		}).Errorln("error while updating image into playlist")
		return err
	}

	return nil
}

func (p *Postgres) Delete(ctx context.Context, playlistId uint64) error {
	p.logger.Infoln("Playlist Repo Delete entered")

	query := "delete from playlist where playlist.id = $1"
	_, err := p.Pool.Exec(ctx, query, playlistId)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"error":       err,
			"playlist id": playlistId,
			"query":       query,
		}).Errorln("error while deleting track from playlist relation")
		return err
	}

	return nil
}

//type Repository interface {
//	Create(playlist Base) error
//	Get(playlistId uint64) (Base, error)
//	AddTrack(playlistId, trackId uint64) error
//	UpdateImage(playlistId, image string) error
//	Delete(playlistId uint64) error
//}
