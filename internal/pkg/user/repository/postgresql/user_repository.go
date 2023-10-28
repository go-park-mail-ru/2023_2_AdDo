package user_repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sirupsen/logrus"
	"main/internal/common/pgxiface"
	"main/internal/common/utils"
	user_domain "main/internal/pkg/user"
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

func (db *Postgres) UpdateAvatarPath(userId string, path string) error {
	query := "update profile set avatar_url = $1 where id = $2"
	_, err := db.Pool.Exec(context.Background(), query, path, userId)
	if err != nil {
		db.logger.WithFields(logrus.Fields{
			"err":   err,
			"path":  path,
			"query": query,
		}).Errorln("avatar updating failed")
		return err
	}

	return nil
}

func (db *Postgres) GetAvatarPath(userId string) (string, error) {
	var path any

	query := "select avatar_url from profile where id = $1"
	err := db.Pool.QueryRow(context.Background(), query, userId).Scan(&path)
	if err != nil {
		db.logger.WithFields(logrus.Fields{
			"err":     err,
			"user id": userId,
			"query":   query,
		}).Errorln("Getting user crds failed")
		return "", err
	}
	if path != nil {
		return path.(string), nil
	}

	return "", nil
}

func (db *Postgres) RemoveAvatarPath(userId string) error {
	query := "update profile set avatar_url = null where id = $1"
	_, err := db.Pool.Exec(context.Background(), query, userId)
	if err != nil {
		return err
	}

	return nil
}

func (db *Postgres) UpdateUserInfo(user user_domain.User) error {
	db.logger.Infoln("UserRepo UpdateUserInfo entered")

	query := "update profile set nickname = $1, email = $2, birth_date = $3 where id = $4"
	_, err := db.Pool.Exec(context.Background(), query, user.Username, user.Email, user.BirthDate)
	if err != nil {
		db.logger.WithFields(logrus.Fields{
			"err ":   err,
			"user ":  user,
			"query ": query,
		}).Errorln("Getting user crds failed")
		return err
	}
	db.logger.Infoln("Updated user info")

	return nil
}
