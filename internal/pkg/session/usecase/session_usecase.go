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

func (s *Default) GetUserId(sessionId string) (string, error) {
	userId, err := s.repoSession.Get(sessionId)
	if err != nil {
		return "", err
	}
	return userId, nil
}

func (s *Default) CheckSession(sessionId string) (bool, error) {
	_, err := s.repoSession.Get(sessionId)
	if err != nil {
		return false, session.ErrSessionDoesNotExist
	}

	return true, nil
}
