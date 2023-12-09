package wave_repository

import (
	"github.com/sirupsen/logrus"
	postgres "main/internal/common/pgxiface"
	"main/internal/pkg/track"
)

type Postgres struct {
	Pool   postgres.PgxIFace
	logger *logrus.Logger
}

func NewPostgres(pool postgres.PgxIFace, logger *logrus.Logger) *Postgres {
	return &Postgres{Pool: pool, logger: logger}
}

func (p Postgres) SaveTracksToUserPool(userId string, tracks []track.Response) error {
	//TODO implement me
	panic("implement me")
}

func (p Postgres) GetTracksFromUserPool(userId string, count uint32) ([]track.Response, error) {
	//TODO implement me
	panic("implement me")
}
