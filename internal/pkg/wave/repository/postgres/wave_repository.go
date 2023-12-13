package wave_repository

import (
	"context"
	"github.com/sirupsen/logrus"
	postgres "main/internal/common/pgxiface"
	"main/internal/pkg/track"
)

type Postgres struct {
	Pool   postgres.PgxIFace
	logger *logrus.Logger
}

func NewRepo(p postgres.PgxIFace, logger *logrus.Logger) Postgres {
	return Postgres{Pool: p, logger: logger}
}

func (p Postgres) deleteOldPool(userId string) error {
	p.logger.Infoln("Get Tracks From User Pool entered")
	query := `delete from wave where owner_id = $1`
	_, err := p.Pool.Exec(context.Background(), query, userId)
	if err != nil {
		p.logger.Errorln("Delete old wave", err)
		return err
	}

	return nil
}

func (p Postgres) SaveTracksToUserPool(userId string, tracks []track.Response) error {
	p.logger.Infoln("Save Tracks To User Pool entered")

	err := p.deleteOldPool(userId)
	if err != nil {
		p.logger.Errorln("Delete old wave", err)
		return err
	}

	query := `insert into wave (owner_id) values ($1)`
	_, err = p.Pool.Exec(context.Background(), query, userId)
	if err != nil {
		p.logger.Errorln("create new wave error", err)
		return err
	}

	waveId := 0
	query = `select id from wave where owner_id = $1`
	err = p.Pool.QueryRow(context.Background(), query, userId).Scan(&waveId)
	if err != nil {
		p.logger.Errorln("select wave error", err)
		return err
	}

	query = `insert into wave_track (track_id, wave_id) values ($1, $2)`
	for _, t := range tracks {
		_, err = p.Pool.Exec(context.Background(), query, t.Id, waveId)
		if err != nil {
			p.logger.Errorln("add track to wave", err)
			return err
		}
	}

	return nil
}
