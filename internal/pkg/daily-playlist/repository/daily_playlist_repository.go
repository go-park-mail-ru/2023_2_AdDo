package daily_playlist_repository

import (
	"context"
	"github.com/sirupsen/logrus"
	postgres "main/internal/common/pgxiface"
	daily_playlist "main/internal/pkg/daily-playlist"
)

type Postgres struct {
	Pool   postgres.PgxIFace
	logger *logrus.Logger
}

func NewPostgres(p postgres.PgxIFace, l *logrus.Logger) Postgres {
	return Postgres{
		Pool:   p,
		logger: l,
	}
}

func (p Postgres) deleteOldUserDaily(userId string) error {
	p.logger.Infoln("Daily Playlist Get User Playlist repo entered")

	query := `delete from daily_playlist where owner_id = $1`
	_, err := p.Pool.Exec(context.Background(), query, userId)
	if err != nil {
		p.logger.Errorln("error deleting daily", err)
		return err
	}

	return nil
}

func (p Postgres) SetUserPlaylist(userId string, response daily_playlist.Response) error {
	p.logger.Infoln("Daily Playlist Set User Playlist repo entered")

	err := p.deleteOldUserDaily(userId)
	if err != nil {
		p.logger.Errorln("error deleting daily", err)
	}

	query := `insert into daily_playlist(owner_id) values ($1)`
	_, err = p.Pool.Exec(context.Background(), query, userId)
	if err != nil {
		p.logger.Errorln("error creating daily", err)
		return err
	}

	playlistId := 0
	query = `select id from daily_playlist values where owner_id = $1`
	err = p.Pool.QueryRow(context.Background(), query, userId).Scan(&playlistId)
	if err != nil {
		p.logger.Errorln("error selecting daily", err)
		return err
	}

	query = `insert into daily_playlist_track (track_id, daily_playlist_id) values ($1, $2)`
	for _, t := range response.Tracks {
		_, err = p.Pool.Exec(context.Background(), query, t.Id, playlistId)
		if err != nil {
			p.logger.Errorln("error filling daily", err)
			return err
		}
	}

	return nil
}
