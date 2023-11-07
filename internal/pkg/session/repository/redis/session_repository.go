package session_repository_redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"main/internal/pkg/session"
)

type Redis struct {
	database *redis.Client
	logger   *logrus.Logger
}

func NewRedis(db *redis.Client, logger *logrus.Logger) *Redis {
	return &Redis{database: db, logger: logger}
}

func (redis *Redis) Create(userId string) (string, error) {
	redis.logger.Infoln("SessionRepo Create entered")

	sessionId := uuid.New().String()

	if err := redis.database.Set(context.Background(), sessionId, userId, session.TimeToLiveCookie).Err(); err != nil {
		redis.logger.WithFields(logrus.Fields{
			"err":     err.Error(),
			"user_id": userId,
		}).Errorln("error with setting new session id in redis")
		return "", err
	}
	redis.logger.Infoln("New session Id set in redis db")

	return sessionId, nil
}

func (redis *Redis) Get(sessionId string) (string, error) {
	redis.logger.Infoln("SessionRepo Get entered")

	userId, err := redis.database.Get(context.Background(), sessionId).Result()
	if err != nil {
		redis.logger.WithFields(logrus.Fields{
			"err":        err.Error(),
			"session id": sessionId,
		}).Errorln("error while getting user id from redis by session id. Are you sure you have this session Id in db?")
		return "", err
	}
	redis.logger.Infoln("Session id matched with stored one")

	return userId, nil
}

func (redis *Redis) Delete(sessionId string) error {
	redis.logger.Infoln("SessionRepo Get entered")

	if err := redis.database.Del(context.Background(), sessionId).Err(); err != nil {
		redis.logger.WithFields(logrus.Fields{
			"err":        err.Error(),
			"session id": sessionId,
		}).Errorln("error while deleting user id from redis by session id. Are you sure you have this session Id in db?")
		return err
	}
	redis.logger.Infoln("Session id successfully deleted from redis")

	return nil
}
