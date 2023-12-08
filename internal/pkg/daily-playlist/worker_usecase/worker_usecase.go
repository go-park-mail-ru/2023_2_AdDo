package daily_playlist_worker_usecase

import (
	"github.com/sirupsen/logrus"
	"main/internal/pkg/candidate"
	daily_playlist "main/internal/pkg/daily-playlist"
	"main/internal/pkg/recommendation"
	user_domain "main/internal/pkg/user"
)

type Default struct {
	repoUser              user_domain.Repository
	dailyPlaylistRepo     daily_playlist.Repository
	logger                *logrus.Logger
	candidateUseCase      candidate.UseCase
	recommendationUseCase recommendation.ServiceUseCase
}

func NewDefault(ruc recommendation.ServiceUseCase, ru user_domain.Repository, dpr daily_playlist.Repository, cuc candidate.UseCase, l *logrus.Logger) Default {
	return Default{
		repoUser:              ru,
		logger:                l,
		dailyPlaylistRepo:     dpr,
		candidateUseCase:      cuc,
		recommendationUseCase: ruc,
	}
}

func (d *Default) CreateDailyForUser(userId string) (daily_playlist.Response, error) {
	d.logger.Infoln("create daily for concrete user entered", userId)

	// Активность представляет собой отношение трека к действию:
	// Действие может быть нескольких типов:
	// 1. Like
	// 2. Listen
	// 3. Skip
	// 4. AddToRotation
	// Добавление трека в ротацию подразумевает, что этот трек был ни лайкнут, ни послушан, но был лайкнут альбом, артист или жанр, который
	// содержит этот трек, а сам трек является ярым представителем этой сущности
	// *Ярый представитель в данном случае - это объект лежащий близко к центру кластера(евклидово расстояние между этим объектом и центром не больше,
	// чем у N других объектов этого кластера)

	//userActivity, err := d.repoUser.GetLastUserActivity(userId, UserActivityBatchForDailyPlaylist)
	//if err != nil {
	//	d.logger.Errorln("error getting last user activity", err, userId)
	//	return daily_playlist.Response{}, err
	//}
	candidates, err := d.candidateUseCase.GetCandidateForUser(userId)
	if err != nil {
		d.logger.Errorln("got candidates for user finished with error", err, userId, candidates)
		return daily_playlist.Response{}, err
	}
	d.logger.Errorln("got candidates for user completed", userId)

	// нейронка классифицирует, какое действие произведет пользователь, лайк, прослушивание или скип, она же ранжирует их
	candidatesAfterClassify, err := d.recommendationUseCase.ClassifyCandidates(userId, candidates)
	if err != nil {
		d.logger.Errorln("classify candidates for user finished with error", err, userId, candidates)
		return daily_playlist.Response{}, err
	}
	d.logger.Errorln("classified candidates for user completed", userId)

	return daily_playlist.Response{
		OwnerId: userId,
		Tracks:  candidatesAfterClassify,
	}, nil
}

func (d *Default) CreateDailyPlaylistForUsers() {
	d.logger.Infoln("start creating user playlist")

	userIds, err := d.repoUser.GetAllUserIds()
	if err != nil {
		d.logger.Errorln("error getting user ids", err)
	}

	for _, uId := range userIds {
		dailyPlaylist, err := d.CreateDailyForUser(uId)
		if err != nil {
			d.logger.Errorln("error getting creating playlist for concrete user", err, uId)
		}

		err = d.dailyPlaylistRepo.SetUserPlaylist(uId, dailyPlaylist)
		if err != nil {
			d.logger.Errorln("error saving created daily for user", err, uId, dailyPlaylist)
		}
	}

	d.logger.Infoln("all daily created for all users successfully")
}
