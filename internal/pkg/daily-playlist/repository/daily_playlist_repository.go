package daily_playlist_repository

import (
	"github.com/sirupsen/logrus"
	postgres "main/internal/common/pgxiface"
	daily_playlist "main/internal/pkg/daily-playlist"
)

type Postgres struct {
	Pool   postgres.PgxIFace
	logger *logrus.Logger
}

func (p Postgres) GetUserDailyPlaylist(userId string) (daily_playlist.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (p Postgres) SetUserPlaylist(userId string, response daily_playlist.Response) error {
	//TODO implement me
	panic("implement me")
}

func NewPostgres(pool postgres.PgxIFace, logger *logrus.Logger) *Postgres {
	return &Postgres{Pool: pool, logger: logger}
}
