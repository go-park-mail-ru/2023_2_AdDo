package session

import (
	"errors"
	"time"
)

type UseCase interface {
	CheckSession(sessionId string) (bool, error)
	GetUserId(sessionId string) (string, error)
}

type Repository interface {
	Create(userId string) (string, error)
	Get(sessionId string) (string, error)
	Delete(sessionId string) error
}

const CookieName = "JSESSIONID"
const TimeToLiveCookie = 1 * time.Minute
const TimeToLiveCSRF = 24 * 60 * 60

var CSRFKey = []byte("6rOD7Jb4gwBISPPd4T2CVDHEILjr4rq2")

const XCsrfToken = "X-Csrf-Token"

var (
	ErrSessionDoesNotExist   = errors.New("session does not exist")
	ErrSessionCreatingFailed = errors.New("session hasn't created")
)
