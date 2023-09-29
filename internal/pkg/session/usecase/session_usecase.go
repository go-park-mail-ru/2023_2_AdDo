package session_usecase

import (
	"main/internal/pkg/session"
	user_domain "main/internal/pkg/user"
)

type Default struct {
	repoSession session.Repository
}

func NewDefault(repository session.Repository) Default {
	return Default{
		repoSession: repository,
	}
}

func (s *Default) CheckSession(sessionId string, userId uint64) (bool, error) {
	isSession, err := s.CheckSession(sessionId, userId)
	if err != nil {
		return false, user_domain.ErrSessionDoesNotExist
	}
	return isSession, nil
}

func (s *Default) Expire(userId uint64) error {
	err := s.repoSession.DeleteByUserId(userId)
	if err != nil {
		return user_domain.ErrSessionDoesNotExist
	}
	return nil
}
