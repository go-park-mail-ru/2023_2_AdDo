package onboarding_repository

import (
	"context"
	"github.com/sirupsen/logrus"
	postgres "main/internal/common/pgxiface"
	"main/internal/pkg/artist"
	"main/internal/pkg/onboarding"
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

func (p *Postgres) GetGenres() ([]onboarding.GenreBase, error) {
	p.logger.Infoln("Onboarding repo Get Genres entered")

	query := `select id, ru_name, icon_url from genre`
	result, err := p.Pool.Query(context.Background(), query)
	if err != nil {
		p.logger.Errorln("err while getting genres ", err)
		return nil, err
	}

	genres := make([]onboarding.GenreBase, 0)
	for result.Next() {
		var base onboarding.GenreBase
		err := result.Scan(&base.Id, &base.Name, &base.Preview)
		if err != nil {
			p.logger.Errorln("err while getting genres ", err)
			return nil, err
		}
		genres = append(genres, base)
	}

	return genres, nil
}

func (p *Postgres) GetArtists() ([]artist.Base, error) {
	p.logger.Infoln("Onboarding repo Get Artists entered")

	query := `select id, name, avatar from artist`
	result, err := p.Pool.Query(context.Background(), query)
	if err != nil {
		p.logger.Errorln("err while getting artists ", err)
		return nil, err
	}

	artists := make([]artist.Base, 0)
	for result.Next() {
		var base artist.Base
		err := result.Scan(&base.Id, &base.Name, &base.Avatar)
		if err != nil {
			p.logger.Errorln("err while scanning artists ", err)
			return nil, err
		}
		artists = append(artists, base)
	}

	return artists, nil
}

func (p *Postgres) SetUserGenres(userId string, genres []onboarding.GenreBase) error {
	p.logger.Infoln("Onboarding repo Set Genres entered")

	query := `insert into profile_genre (profile_id, genre_id) values ($1, $2) on conflict (profile_id, genre_id) do nothing `
	for _, genre := range genres {
		_, err := p.Pool.Exec(context.Background(), query, userId, genre.Id)
		if err != nil {
			p.logger.Errorln("err while adding profile genre ", err)
			return err
		}
	}

	return nil
}

func (p *Postgres) SetUserArtists(userId string, artists []artist.Base) error {
	p.logger.Infoln("Onboarding repo Set Artists entered")

	query := `insert into profile_artist (profile_id, artist_id) values ($1, $2) on conflict (profile_id, artist_id) do nothing `
	for _, a := range artists {
		_, err := p.Pool.Exec(context.Background(), query, userId, a.Id)
		if err != nil {
			p.logger.Errorln("err while adding profile artist ", err)
			return err
		}
	}

	return nil
}
