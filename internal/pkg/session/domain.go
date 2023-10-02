package session

import (
	"time"
)

type Session struct {
	SessionId  string
	ProfileId  uint64
	Expiration time.Time
}

type UseCase interface {
	CheckSession(sessionId string, userId uint64) (bool, error)
}

type Repository interface {
	Create(userId uint64) (string, error)
	GetByUserId(userId uint64) (string, error)
	DeleteByUserId(userId uint64) error
}

const CookieName = "JSESSIONID"
const TimeToLive = 1 * time.Minute
