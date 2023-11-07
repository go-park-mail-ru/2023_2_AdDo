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

func (p *Postgres) Create(ctx context.Context, base playlist.Base) (playlist.Response, error) {
	p.logger.Infoln("Playlist Repo Create entered")

	var url any
	var result playlist.Response
	query := "insert into playlist (creator_id) values ($1) returning id, name, creator_id, preview"
	if err := p.Pool.QueryRow(ctx, query, base.AuthorId).Scan(&result.Id, &result.Name, &result.AuthorId, &url); err != nil {
		p.logger.WithFields(logrus.Fields{
			"error":         err,
			"playlist base": base,
			"query":         query,
		}).Errorln("error while creating playlist row")
		return playlist.Response{}, err
	}

	if url != nil {
		result.Preview = url.(string)
	}

	return result, nil
}

func (p *Postgres) Get(ctx context.Context, playlistId uint64) (playlist.Base, error) {
	p.logger.Infoln("Playlist Repo Get entered")

	var url any
	var result playlist.Base
	query := "select id, name, creator_id, preview from playlist where id = $1"
	if err := p.Pool.QueryRow(ctx, query, playlistId).Scan(&result.Id, &result.Name, &result.AuthorId, &url); err != nil {
		p.logger.WithFields(logrus.Fields{
			"error":       err,
			"playlist id": playlistId,
			"query":       query,
		}).Errorln("error while selecting playlist row by id")
		return playlist.Base{}, err
	}

	if url != nil {
		result.Preview = url.(string)
	}

	return result, nil
}

func (p *Postgres) GetByCreatorId(ctx context.Context, userId string) ([]playlist.Base, error) {
	p.logger.Infoln("Playlist Repo GetByCreatorId entered")

	var url any
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
		if err := rows.Scan(&base.Id, &base.Name, &base.AuthorId, &url); err != nil {
			p.logger.WithFields(logrus.Fields{
				"query":   query,
				"err":     err,
				"user id": userId,
			}).Errorln("error scanning row")
			return nil, err
		}

		if url != nil {
			base.Preview = url.(string)
		}

		result = append(result, base)
	}
	p.logger.Infoln("Scanning rows success")

	return result, nil
}

func (p *Postgres) AddTrack(ctx context.Context, playlistId, trackId uint64) error {
	p.logger.Infoln("Playlist Repo AddTrack entered")

	query := "insert into playlist_track (playlist_id, track_id) values ($1, $2)"
	if _, err := p.Pool.Exec(ctx, query, playlistId, trackId); err != nil {
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
	if _, err := p.Pool.Exec(ctx, query, playlistId, trackId); err != nil {
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
	if _, err := p.Pool.Exec(ctx, query, image, playlistId); err != nil {
		p.logger.WithFields(logrus.Fields{
			"error":           err,
			"images url path": image,
			"playlist id":     playlistId,
			"query":           query,
		}).Errorln("error while updating images into playlist")
		return err
	}

	return nil
}

func (p *Postgres) RemovePreviewPath(ctx context.Context, playlistId uint64) (string, error) {
	p.logger.Infoln("Playlist Repo UpdateImage entered")

	var result string
	query := "update playlist set preview = null where id = $1 returning preview"
	if err := p.Pool.QueryRow(ctx, query, playlistId).Scan(&result); err != nil {
		p.logger.WithFields(logrus.Fields{
			"error":       err,
			"playlist id": playlistId,
			"preview":     result,
			"query":       query,
		}).Errorln("error while updating images into playlist")
		return "", err
	}

	return result, nil
}

func (p *Postgres) Delete(ctx context.Context, playlistId uint64) error {
	p.logger.Infoln("Playlist Repo Delete entered")

	query := "delete from playlist where playlist.id = $1"
	if _, err := p.Pool.Exec(ctx, query, playlistId); err != nil {
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
	p.logger.Infoln("Playlist Repo CreateLike entered")

	query := "insert into profile_playlist (profile_id, playlist_id) values ($1, $2)"
	if _, err := p.Pool.Exec(context.Background(), query, userId, playlistId); err != nil {
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

func (p *Postgres) CheckLike(ctx context.Context, userId string, playlistId uint64) (bool, error) {
	p.logger.Infoln("Playlist Repo CheckLike entered")

	var counter int
	query := "select count(*) from profile_playlist where profile_id = $1 and playlist_id = $2"
	if err := p.Pool.QueryRow(context.Background(), query, userId, playlistId).Scan(&counter); err != nil {
		p.logger.Errorln(err)
		return false, err
	}
	p.logger.Infoln("like checked")

	if counter == 0 {
		return false, nil
	}

	return true, nil
}

func (p *Postgres) DeleteLike(ctx context.Context, userId string, playlistId uint64) error {
	p.logger.Infoln("playlist Repo CreateLike entered")

	query := "delete from profile_playlist where profile_id = $1 and playlist_id = $2"
	if _, err := p.Pool.Exec(context.Background(), query, userId, playlistId); err != nil {
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

func (p *Postgres) IsCreator(ctx context.Context, userId string, playlistId uint64) (bool, error) {
	p.logger.Infoln("Album Repo IsCreator entered")

	var creatorId string
	query := "select creator_id from playlist where id = $1"
	if err := p.Pool.QueryRow(context.Background(), query, playlistId).Scan(&creatorId); err != nil {
		p.logger.WithFields(logrus.Fields{
			"err":         err,
			"playlist Id": playlistId,
			"query":       query,
		}).Errorln("getting creator error")
		return false, err
	}
	p.logger.Infoln("Got creatorId")

	return creatorId == userId, nil
}

func (p *Postgres) IsPrivate(ctx context.Context, playlistId uint64) (bool, error) {
	p.logger.Infoln("Album Repo IsPrivate entered")

	var isPrivate bool
	query := "select is_private from playlist where id = $1"
	if err := p.Pool.QueryRow(context.Background(), query, playlistId).Scan(&isPrivate); err != nil {
		p.logger.WithFields(logrus.Fields{
			"err":         err,
			"playlist Id": playlistId,
			"query":       query,
		}).Errorln("getting is_private error")
		return false, err
	}
	p.logger.Infoln("Got isPrivate")

	return isPrivate, nil
}

func (p *Postgres) MakePublic(ctx context.Context, playlistId uint64) error {
	p.logger.Infoln("Album Repo MakePublic entered")

	query := "update playlist set is_private = false where id = $1"
	if _, err := p.Pool.Exec(context.Background(), query, playlistId); err != nil {
		p.logger.WithFields(logrus.Fields{
			"err":         err,
			"playlist Id": playlistId,
			"query":       query,
		}).Errorln("updating is_private error")
		return err
	}

	return nil
}

func (p *Postgres) MakePrivate(ctx context.Context, playlistId uint64) error {
	p.logger.Infoln("Album Repo MakePrivate entered")

	query := "update playlist set is_private = true where id = $1"
	if _, err := p.Pool.Exec(context.Background(), query, playlistId); err != nil {
		p.logger.WithFields(logrus.Fields{
			"err":         err,
			"playlist Id": playlistId,
			"query":       query,
		}).Errorln("updating is_private error")
		return err
	}

	return nil
}
