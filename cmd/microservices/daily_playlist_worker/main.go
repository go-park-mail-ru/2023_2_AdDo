package main

import (
	"github.com/robfig/cron/v3"
	log "main/internal/common/logger"
)

var loggerSingleton = log.Singleton{}

func CreateDailyPlaylistForUsers() {
	logger := loggerSingleton.GetLogger()
	logger.Infoln("Start cooking playlists")

	// получаем всех пользователей, для каждого берем последнюю активность
	// c этой активностью идем выбираем кандидатов
	// после выбора кандидатов ранжируем их по данным из нейронки
}

func main() {
	logger := loggerSingleton.GetLogger()
	c := cron.New()

	_, err := c.AddFunc("@midnight", CreateDailyPlaylistForUsers)
	if err != nil {
		logger.Errorln("Can't add function to schedule:", err)
		return
	}

	c.Start()

	select {}
}
