package session

import (
	"errors"
	"time"
)

type UseCase interface {
	CheckSession(sessionId string) (bool, error)
}

type Repository interface {
	Create(userId uint64) (string, error)
	Get(sessionId string) (uint64, error)
	Delete(sessionId string) error
}

const CookieName = "JSESSIONID"
const TimeToLive = 1 * time.Minute

var (
	ErrSessionDoesNotExist   = errors.New("session does not exist")
	ErrSessionCreatingFailed = errors.New("session hasn't created")
)
