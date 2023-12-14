package activity_usecase

import (
	"github.com/sirupsen/logrus"
	"main/internal/pkg/activity"
	"main/internal/pkg/candidate"
	"main/internal/pkg/recommendation"
	"main/internal/pkg/wave"
)

type Worker struct {
	activityRepo          activity.ConsumerRepository
	recentActivityRepo    activity.KeyValueRepository
	candidateUseCase      candidate.UseCase
	trackPool             wave.PoolRepository
	recommendationUseCase recommendation.ServiceUseCase
	logger                *logrus.Logger
}

func NewWorker(ruc recommendation.ServiceUseCase, cus candidate.UseCase, up wave.PoolRepository, rar activity.KeyValueRepository, uar activity.ConsumerRepository, logger *logrus.Logger) Worker {
	return Worker{
		activityRepo:          uar,
		trackPool:             up,
		logger:                logger,
		recentActivityRepo:    rar,
		candidateUseCase:      cus,
		recommendationUseCase: ruc,
	}
}

// Цель этого воркера, в потоке получать данные из очередей(если они есть) и обрабатывать их, приводя к единому виду. Когда активности становится достаточно, то
// Мы обновляем пул треков. Обновляем мы его через кандидатов и нейронку.

// Итак, итоговый флоу волны выглядит так:
// 1. Юзер регистрируется и проходит обязательный онбординг
// 2. Этот онбординг оказывается в нашей очереди
// 3. Activity worker разгружает наши лайки из очереди и приводит их к виду LikeTrack, ListenTrack, SkipTrack
// 4. Когда Недавняя активность достигает определенного размера, например, 5(Речь идет о действиях LikeTrack, ListenTrack),
// Воркеру достаточно информации чтобы создать/пересоздать пул треков пользователя.
// * Пул треков - некоторое количество треков, пусть 50. Из этого пула мы берем пачку для "Моей волны" по сигналу от фронта.
// ** Фронт присылает сигнал, когда пользователь выслушал всю прошлую пачку. Если треков в пуле нет - берем рандомные из базы данных.
// *** Когда мы берем данные из пула, мы удаляем их оттуда, а в деливери запоминаем, какие треки мы отсылали, чтобы не прислать дублей (Это можно запомнить и в другом месте)
// 5. Создание/Пересоздание пула выглядит следующим образом:
// 5.1 Активити Воркер идет в микрос кандидатов(МК) с UserId.
// 5.2 МК собирает дополнительные данные:
// 5.2.1 Горячие треки пользователя
// 5.2.2 Недавняя Активность
// 5.2.3 Неитересные треки (треки, которые пользователь сразу скипал)
// 5.2.4 Треки на перспективу (те треки, которые получили при конвертации, они чаще всего идут из недавней активности)
// 5.3 МК выбирает несколько оптимальных объектов в нашем кластерном пространстве, собирает оттуда треки в некоторой пропорции(с этим надо экспериментировать)
// 5.4 Пул практически готов. МК возвращает эти треки воркеру и тут можно было бы их уже положить в наш новый пул. Но это не всё.
// 6. Воркер относит кандидатов нейронной сети, которая попробует предсказать какое действие совершит пользователь. Она же ранжирует их по этому действию и возвращает воркеру
// * Рекуррентная Нейронка делает предсказания только на основе последней активности
// 7. Воркер с гордостью удаляет перезаписывает старый пул на новый, из которого "Моя Волна" возьмет следующую пачку. И так по кругу.

// Итак, итоговый флоу дейликов:
// 1. Каждый день в 00:00 по Мск запускается джоба Дейли Воркера (ДВ)
// 2. Эта ДВ берет айдишники всех пользователей в нашей БД
// 3. Проходится по каждому айдишнику и делает с ним следующее:
// 3.1 Идёт к МК (микросу кандидатов) с айдишником юзера.
// 3.2 Получает кандидатов аналогично с предыдущим примером
// 3.3 Несет эти треки нейронке, которая ранжирует их аналогично с предыдущим примеров.
// 4. Сохраняет созданный плейлист в базу, который пользователь впоследствии может запросить.

// Здесь мы держим мапу, в которой лежат активности. Мы прям тут проверяем сколько их у пользователя и в случае чего тригеррим старт переварки пула
func (s *Worker) Run() {
	s.logger.Infoln("Activity worker started")
	actions := make(chan activity.UserTrackAction)

	go s.activityRepo.PopLikeAlbum(actions)
	go s.activityRepo.PopLikeTrack(actions)
	go s.activityRepo.PopLikeArtist(actions)
	go s.activityRepo.PopLikeGenre(actions)
	go s.activityRepo.PopSkipTrack(actions)
	go s.activityRepo.PopListenTrack(actions)
	s.logger.Infoln("All worker gorutines initialized")

	for action := range actions {
		isFull, err := s.recentActivityRepo.SaveActivityAndCountCheck(action, activity.RecentActivityNeedToRecreateTrackPool)
		if err != nil {
			s.logger.Errorln("error while saving recent activity", err)
		}

		if !isFull {
			continue
		}

		err = s.RecreatePool(action.UserId)
		if err != nil {
			s.logger.Errorln("error while recreating new pool", err)
		}
	}
}

func (s *Worker) RecreatePool(userId string) error {
	defer func(recentActivityRepo activity.KeyValueRepository, userId string) {
		err := recentActivityRepo.CleanLastActivityForUser(userId)
		if err != nil {
			s.logger.Errorln("error cleaning recent activities", err)
		}
	}(s.recentActivityRepo, userId)

	candidates, err := s.candidateUseCase.GetCandidateForWave(userId)
	if err != nil {
		s.logger.Errorln("error getting candidates for user", err)
		return err
	}

	classifiedCandidates, err := s.recommendationUseCase.ClassifyCandidates(userId, candidates)
	if err != nil {
		s.logger.Errorln("error while RNN magic", err)
		return err
	}

	err = s.trackPool.SaveTracksToUserPool(userId, classifiedCandidates)
	if err != nil {
		s.logger.Errorln("error while saving new pool", err)
		return err
	}

	return nil
}
