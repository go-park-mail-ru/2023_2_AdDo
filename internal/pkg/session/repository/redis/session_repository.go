package session_repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"main/internal/pkg/session"
)

type Redis struct {
	database *redis.Client
	ctx      context.Context
}

func NewRedis(db *redis.Client, ctx context.Context) *Redis {
	return &Redis{database: db, ctx: ctx}
}

func (redis *Redis) Create() (string, error) {
	sessionId := uuid.New().String()

	err := redis.database.Set(redis.ctx, sessionId, true, session.TimeToLive).Err()
	if err != nil {
		return "", err
	}

	return sessionId, nil
}

func (redis *Redis) Get(sessionId string) (bool, error) {
	_, err := redis.database.Get(redis.ctx, sessionId).Result()
	if err != nil {
		return false, err
	}

	return true, nil
}

func (redis *Redis) Delete(sessionId string) error {
	err := redis.database.Del(redis.ctx, sessionId).Err()
	if err != nil {
		return err
	}
	return nil
}
