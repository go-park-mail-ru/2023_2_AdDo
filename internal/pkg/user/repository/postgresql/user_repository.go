package user_repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
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

func (db *Postgres) GetById(id string) (user_domain.User, error) {
	user := user_domain.User{Id: id}
	var dt pgtype.Date
	var avatar pgtype.Text

	err := db.Pool.QueryRow(context.Background(), "select email, nickname, birth_date, avatar_url from profile where id = $1", id).Scan(&user.Email, &user.Username, &dt, &avatar)
	if err != nil {
		return user, err
	}

	user.BirthDate = dt.Time.Format("2006-01-02")
	user.Avatar = avatar.String

	return user, err
}

func (db *Postgres) CheckEmailAndPassword(email, password string) (string, error) {
	var userId string

	err := db.Pool.QueryRow(context.Background(), "select id from profile where email = $1 and password = $2", email, utils.GetMD5Sum(password)).Scan(&userId)
	if err != nil {
		return "", err
	}

	return userId, nil
}
