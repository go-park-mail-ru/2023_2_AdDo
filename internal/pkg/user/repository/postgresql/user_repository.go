package user_repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sirupsen/logrus"
	"main/internal/common/pgxiface"
	"main/internal/common/utils"
	"main/internal/pkg/user"
)

type Postgres struct {
	Pool   postgres.PgxIFace
	logger *logrus.Logger
}

func NewPostgres(pool postgres.PgxIFace, logger *logrus.Logger) *Postgres {
	return &Postgres{Pool: pool, logger: logger}
}

func (db *Postgres) Create(user user_domain.User) error {
	db.logger.Infoln("UserRepo Create entered")

	query := "insert into profile (email, password, nickname, birth_date) values ($1, $2, $3, $4)"
	_, err := db.Pool.Exec(context.Background(), query, user.Email, utils.GetMD5Sum(user.Password), user.Username, user.BirthDate)
	if err != nil {
		db.logger.WithFields(logrus.Fields{
			"err":       err,
			"user_data": user,
			"query":     query,
		}).Errorln("Creating a user query completed with error")
		return err
	}
	db.logger.Infoln("User created successfully in postgres")

	return nil
}

func (db *Postgres) GetById(id string) (user_domain.User, error) {
	db.logger.Infoln("UserRepo GetById entered")

	user := user_domain.User{Id: id}
	var dt pgtype.Date
	var avatar pgtype.Text

	query := "select email, nickname, birth_date, avatar_url from profile where id = $1"
	err := db.Pool.QueryRow(context.Background(), query, id).Scan(&user.Email, &user.Username, &dt, &avatar)
	if err != nil {
		db.logger.WithFields(logrus.Fields{
			"err":       err,
			"user_data": user,
			"query":     query,
		}).Errorln("Getting user by id failed")
		return user, err
	}
	db.logger.Infoln("Getting user completed")

	user.BirthDate = dt.Time.Format("2006-01-02")
	user.Avatar = avatar.String
	db.logger.Infoln("birthday and avatar formatted")

	return user, err
}

func (db *Postgres) CheckEmailAndPassword(email, password string) (string, error) {
	db.logger.Infoln("UserRepo CheckEmailAndPassword entered")

	var userId string

	query := "select id from profile where email = $1 and password = $2"
	err := db.Pool.QueryRow(context.Background(), query, email, utils.GetMD5Sum(password)).Scan(&userId)
	if err != nil {
		db.logger.WithFields(logrus.Fields{
			"err":       err,
			"user_data": email,
			"query":     query,
		}).Errorln("Checking user crds failed")
		return "", err
	}
	db.logger.Infoln("User credentials checked for user ", email)

	return userId, nil
}
