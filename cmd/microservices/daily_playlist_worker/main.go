package main

import (
	"github.com/robfig/cron/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	init_db "main/init/postgres_db"
	log "main/internal/common/logger"
	candidate "main/internal/microservices/candidate/proto"
	candidate_service_client "main/internal/microservices/candidate/service/client"
	candidate2 "main/internal/microservices/recommendation/proto"
	recommendation_service_client "main/internal/microservices/recommendation/service/client"
	daily_playlist_repository "main/internal/pkg/daily-playlist/repository"
	daily_playlist_worker_usecase "main/internal/pkg/daily-playlist/worker_usecase"
	user_repository "main/internal/pkg/user/repository/postgresql"
)

var loggerSingleton = log.Singleton{}

const EnvPostgresQueryName = "DATABASE_URL"

func main() {
	logger := loggerSingleton.GetLogger()
	c := cron.New()

	pool, err := init_db.InitPostgres(EnvPostgresQueryName)
	if err != nil {
		logger.Errorln("error connecting database: ", err)
	}

	repoUser := user_repository.NewPostgres(pool, logger)
	dailyPlaylistRepo := daily_playlist_repository.NewPostgres(pool, logger)

	candidateConnection, err := grpc.Dial("candidate:8088", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatalln("error connecting to candidate micros ", err)
	}

	candidateAgent := candidate_service_client.NewClient(candidate.NewCandidateServiceClient(candidateConnection), logger)

	recommendationConnection, err := grpc.Dial("recommendation:8089", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatalln("error connecting to recommendation micros ", err)
	}

	recommendationAgent := recommendation_service_client.NewClient(candidate2.NewRecommendationServiceClient(recommendationConnection), logger)
	dailyPlaylistWorkerUseCase := daily_playlist_worker_usecase.NewDefault(&recommendationAgent, repoUser, dailyPlaylistRepo, &candidateAgent, logger)

	_, err = c.AddFunc("@midnight", dailyPlaylistWorkerUseCase.CreateDailyPlaylistForUsers)
	if err != nil {
		logger.Errorln("Can't add function to schedule:", err)
		return
	}

	c.Start()

	select {}
}
