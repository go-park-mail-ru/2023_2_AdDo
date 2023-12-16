package user_repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sirupsen/logrus"
	postgres "main/internal/common/pgxiface"
	"main/internal/common/utils"
	user_domain "main/internal/pkg/user"
	"time"
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
	if _, err := db.Pool.Exec(context.Background(), query, user.Email, utils.GetMD5Sum(user.Password), user.Username, user.BirthDate); err != nil {
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
	if err := db.Pool.QueryRow(context.Background(), query, id).Scan(&user.Email, &user.Username, &dt, &avatar); err != nil {
		db.logger.WithFields(logrus.Fields{
			"err":       err,
			"user_data": user,
			"query":     query,
		}).Errorln("Getting user by id failed")
		return user, err
	}
	db.logger.Infoln("Getting user completed")

	user.BirthDate = dt.Time.Format(time.DateOnly)
	user.Avatar = avatar.String
	db.logger.Infoln("birthday and images formatted")

	return user, nil
}

func (db *Postgres) CheckEmail(email string) (error) {
	db.logger.Infoln("UserRepo CheckEmail entered")

	var userId string

	query := "select id from profile where email = $1"
	if err := db.Pool.QueryRow(context.Background(), query, email).Scan(&userId); err != nil {
		db.logger.WithFields(logrus.Fields{
			"err":       err,
			"user_data": email,
			"query":     query,
		}).Errorln("Checking user email failed")
		return err
	}
	db.logger.Infoln("User email checked for user ", email)

	return nil
}


func (db *Postgres) CheckEmailAndPassword(email, password string) (string, error) {
	db.logger.Infoln("UserRepo CheckEmailAndPassword entered")

	var userId string

	query := "select id from profile where email = $1 and password = $2"
	if err := db.Pool.QueryRow(context.Background(), query, email, utils.GetMD5Sum(password)).Scan(&userId); err != nil {
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
	if _, err := db.Pool.Exec(context.Background(), query, path, userId); err != nil {
		db.logger.WithFields(logrus.Fields{
			"err":   err,
			"path":  path,
			"query": query,
		}).Errorln("images updating failed")
		return err
	}

	return nil
}

func (db *Postgres) GetAvatarPath(userId string) (string, error) {
	var path any

	query := "select avatar_url from profile where id = $1"
	if err := db.Pool.QueryRow(context.Background(), query, userId).Scan(&path); err != nil {
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

func (db *Postgres) RemoveAvatarPath(userId string) (string, error) {
	url, err := db.GetAvatarPath(userId)
	if err != nil {
		db.logger.WithFields(logrus.Fields{
			"err":     err,
			"user id": userId,
		}).Errorln("no active avatar")
	}

	query := "update profile set avatar_url = null where id = $1"
	if _, err = db.Pool.Exec(context.Background(), query, userId); err != nil {
		return "", err
	}
	db.logger.Infoln("Avatar removed", url)

	return url, nil
}

func (db *Postgres) UpdateUserInfo(user user_domain.User) error {
	db.logger.Infoln("UserRepo UpdateUserInfo entered")

	query := "update profile set nickname = $1, email = $2, birth_date = $3 where id = $4"
	if _, err := db.Pool.Exec(context.Background(), query, user.Username, user.Email, user.BirthDate, user.Id); err != nil {
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

func (db *Postgres) GetUserNameById(userId string) (string, error) {
	db.logger.Infoln("UserRepo GetUserNameById entered")

	var result string
	query := "select profile.nickname from profile where id = $1"
	if err := db.Pool.QueryRow(context.Background(), query, userId).Scan(&result); err != nil {
		db.logger.WithFields(logrus.Fields{
			"err ":   err,
			"user ":  userId,
			"query ": query,
		}).Errorln("Getting user nickname failed")
		return result, err
	}

	return result, nil
}
