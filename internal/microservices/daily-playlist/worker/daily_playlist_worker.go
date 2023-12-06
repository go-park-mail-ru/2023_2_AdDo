package worker

import (
	"github.com/sirupsen/logrus"
	daily_playlist "main/internal/pkg/daily-playlist"
)

// здесь крон джоб на регулярное создание дейликов

type Server struct {
	dailyPlaylistRepo daily_playlist.Repository
	logger            *logrus.Logger
}

func NewServer(dp daily_playlist.Repository, logger *logrus.Logger) Server {
	return Server{
		dailyPlaylistRepo: dp,
		logger:            logger,
	}
}

func (s *Server) Run() error {
	return nil
}
