package session

import (
	"time"
)

type Session struct {
	SessionId  string
	ProfileId  uint64
	Expiration time.Time
}

type Usecase interface {
	CheckSession(sessionId string, userId uint64) (bool, error)
	Expire(userId uint64) error
}

type Repository interface {
	Create(userId uint64) (string, error)
	GetByUserId(userId uint64) (string, error)
	DeleteByUserId(userId uint64) error
	GetBySessionId(sessionId string) (string, error)
	DeleteBySessionId(sessionId string) error
}

const CookieName = "JSESSIONID"
const TimeToLive = 1 * time.Minute
