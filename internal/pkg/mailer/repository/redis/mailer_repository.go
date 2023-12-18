package mailer_repository_redis

import (
	"context"
	"main/internal/common/utils"
	domain "main/internal/pkg/mailer"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/thanhpk/randstr"
)

type Redis struct {
	database *redis.Client
	logger   *logrus.Logger
}

func NewRedis(db *redis.Client, logger *logrus.Logger) *Redis {
	return &Redis{database: db, logger: logger}
}

func (redis *Redis) CreateToken(email string) (string, error) {
	redis.logger.Infoln("MailerRepo Create entered")

	resetToken := randstr.String(20)
	passwordResetToken := utils.Encode(resetToken)

	if err := redis.database.Set(context.Background(), passwordResetToken, email, domain.ResetTokenTimeToLive).Err(); err != nil {
		redis.logger.WithFields(logrus.Fields{
			"err":   err.Error(),
			"email": email,
		}).Errorln("error with setting new reset token in redis")
		return "", err
	}
	redis.logger.Infoln("New reset token add in redis db")

	return passwordResetToken, nil
}

func (redis *Redis) CheckToken(passwordResetToken string) (string, error) {
	redis.logger.Infoln("SessionRepo Get entered")

	resetToken, err := utils.Decode(passwordResetToken)
	if err != nil {
		return "", err
	}

	email, err := redis.database.Get(context.Background(), resetToken).Result()
	if err != nil {
		redis.logger.WithFields(logrus.Fields{
			"err":        err.Error(),
			"resetToken": resetToken,
		}).Errorln("error while getting email from redis by reset token. Are you sure you have this reset token in db?")
		return "", err
	}
	redis.logger.Infoln("Reset token matched with stored one")

	return email, nil
}
