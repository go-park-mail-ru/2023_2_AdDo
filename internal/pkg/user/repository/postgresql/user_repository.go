package user_repository

import (
	"context"
	postgres "main/internal/pkg/common/pgxiface"
	"main/internal/pkg/common/utils"
	"main/internal/pkg/user"
)

type Postgres struct {
	Pool postgres.PgxIFace
}

func NewPostgres(pool postgres.PgxIFace) *Postgres {
	return &Postgres{Pool: pool}
}

func (db *Postgres) Create(user user_domain.User) error {
	_, err := db.Pool.Exec(context.Background(), "insert into profile (email, password, nickname, birth_date) values ($1, $2, $3, $4)",
		user.Email, utils.GetMD5Sum(user.Password), user.Username, user.BirthDate)
	if err != nil {
		return err
	}

	return nil
}

func (db *Postgres) GetById(id uint64) (user_domain.User, error) {
	user := user_domain.User{Id: id}
	err := db.Pool.QueryRow(context.Background(), "select email, nickname, birth_date, avatar_url from profile where id = $1", id).Scan(&user.Email, &user.Username, &user.BirthDate, &user.Avatar)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (db *Postgres) CheckEmailAndPassword(email, password string) (uint64, error) {
	var id uint64
	err := db.Pool.QueryRow(context.Background(), "select id from profile where email = $1 and password = $2", email, utils.GetMD5Sum(password)).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
