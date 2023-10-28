package session_usecase

import (
	"github.com/sirupsen/logrus"
	"main/internal/pkg/session"
)

type Default struct {
	repoSession session.Repository
	logger      *logrus.Logger
}

func NewDefault(repository session.Repository, logger *logrus.Logger) Default {
	return Default{
		repoSession: repository,
		logger:      logger,
	}
}

func (s *Default) GetUserId(sessionId string) (string, error) {
	s.logger.Infoln("SessionUseCase GetUserId entered")

	userId, err := s.repoSession.Get(sessionId)
	if err != nil {
		return "", err
	}
	s.logger.Infoln("Got user id")

	return userId, nil
}

func (s *Default) CheckSession(sessionId string) (bool, error) {
	s.logger.Infoln("SessionUseCase CheckSession entered")

	_, err := s.repoSession.Get(sessionId)
	if err != nil {
		return false, session.ErrSessionDoesNotExist
	}
	s.logger.Infoln("Session id matched with an db one")

	return true, nil
}
