package playlist_repository

import (
	"context"
	"github.com/sirupsen/logrus"
	"main/internal/common/pgxiface"
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

func (p *Postgres) GetByCreatorId(ctx context.Context, userId string) ([]playlist.Base, error) {
	p.logger.Infoln("Playlist Repo GetByCreatorId entered")

	result := make([]playlist.Base, 0)
	query := "select id, name, creator_id, preview from playlist where creator_id = $1"
	rows, err := p.Pool.Query(ctx, query, userId)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"query":   query,
			"err":     err,
			"user id": userId,
		}).Errorln("error while getting playlists by user id")
		return nil, err
	}
	p.logger.Infoln("Success getting rows from db")

	for rows.Next() {
		var base playlist.Base
		err := rows.Scan(&base.Id, &base.Name, &base.AuthorId, &base.Preview)
		if err != nil {
			p.logger.WithFields(logrus.Fields{
				"query":   query,
				"err":     err,
				"user id": userId,
			}).Errorln("error scanning row")
			return nil, err
		}
		result = append(result, base)
	}
	p.logger.Infoln("Scanning rows success")

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

func (p *Postgres) RemoveTrack(ctx context.Context, playlistId, trackId uint64) error {
	p.logger.Infoln("Playlist Repo RemoveTrack entered")

	query := "delete from playlist_track where playlist_id = $1 and track_id = $2"
	_, err := p.Pool.Exec(ctx, query, playlistId, trackId)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"error":       err,
			"playlist id": playlistId,
			"track id":    trackId,
			"query":       query,
		}).Errorln("error while removing track into playlist")
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

func (p *Postgres) RemovePreviewPath(ctx context.Context, playlistId uint64) (string, error) {
	p.logger.Infoln("Playlist Repo UpdateImage entered")

	query := "update playlist set preview = null where id = $1 returning preview"
	result, err := p.Pool.Exec(ctx, query, playlistId)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"error":       err,
			"playlist id": playlistId,
			"image":       result.String(),
			"query":       query,
		}).Errorln("error while updating image into playlist")
		return "", err
	}

	return result.String(), nil
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

func (p *Postgres) CreateLike(ctx context.Context, userId string, playlistId uint64) error {
	p.logger.Infoln("Album Repo CreateLike entered")

	query := "insert into musicon.public.profile_playlist (profile_id, playlist_id) values ($1, $2)"
	_, err := p.Pool.Exec(context.Background(), query, userId, playlistId)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"err":         err,
			"playlist Id": playlistId,
			"query":       query,
		}).Errorln("creating like error")
		return err
	}
	p.logger.Infoln("Like Created")

	return nil
}
