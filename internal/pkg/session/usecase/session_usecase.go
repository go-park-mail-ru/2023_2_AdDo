package session_usecase

import (
	"main/internal/pkg/session"
)

type Default struct {
	repoSession session.Repository
}

func NewDefault(repository session.Repository) Default {
	return Default{
		repoSession: repository,
	}
}

func (s *Default) CheckSession(sessionId string) (bool, error) {
	isAuth, err := s.repoSession.Get(sessionId)
	if err != nil {
		return false, session.ErrSessionDoesNotExist
	}
	return isAuth, nil
}
