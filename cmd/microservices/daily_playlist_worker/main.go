package main

import (
	"github.com/robfig/cron/v3"
	log "main/internal/common/logger"
	daily_playlist_worker_usecase "main/internal/pkg/daily-playlist/worker_usecase"
)

var loggerSingleton = log.Singleton{}

func main() {
	logger := loggerSingleton.GetLogger()
	c := cron.New()

	dailyPlaylistWorkerUseCase := daily_playlist_worker_usecase.NewDefault()

	_, err := c.AddFunc("@midnight", dailyPlaylistWorkerUseCase.CreateDailyPlaylistForUsers)
	if err != nil {
		logger.Errorln("Can't add function to schedule:", err)
		return
	}

	c.Start()

	select {}
}
