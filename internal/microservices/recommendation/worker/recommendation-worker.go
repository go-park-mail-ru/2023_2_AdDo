package recommendation_worker

import (
	"github.com/sirupsen/logrus"
	"main/internal/pkg/recommendation"
)

type Server struct {
	userActivityRepo recommendation.QueueRepository
	userPool         recommendation.PoolRepository
	logger           *logrus.Logger
}

func NewServer(up recommendation.PoolRepository, uar recommendation.QueueRepository, logger *logrus.Logger) Server {
	return Server{
		userActivityRepo: uar,
		userPool:         up,
		logger:           logger,
	}
}

// В этой функции мы получаем из очереди Пользовательскую активность, приводим ее к единому виду, ходим в сервис отбора кандидатов
// после чего ходим к нашей нейросети и просим предсказать, насколько пользователю понравится трек, ранжируем результаты по этому параметру и записываем в пул
// пользователя сортированные треки
// запускаем питонячий инстанс с моделью, к которому приходим с данными

func (s *Server) Run() error {
	return nil
}
