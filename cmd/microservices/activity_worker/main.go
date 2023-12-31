package main

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	init_kafka "main/init/kafka_queue"
	init_db "main/init/postgres_db"
	log "main/internal/common/logger"
	candidate "main/internal/microservices/candidate/proto"
	candidate_service_client "main/internal/microservices/candidate/service/client"
	candidate2 "main/internal/microservices/recommendation/proto"
	recommendation_service_client "main/internal/microservices/recommendation/service/client"
	"main/internal/pkg/activity/repository/kafka/activity_repository_consumer"
	activity_repository "main/internal/pkg/activity/repository/memcached"
	activity_usecase "main/internal/pkg/activity/worker_usecase"
	track_repository "main/internal/pkg/track/repository/postgresql"
	wave_repository "main/internal/pkg/wave/repository/postgres"
)

var loggerSingleton = log.Singleton{}
var KafkaQueryConnection = []string{"kafka:9092"}

const EnvPostgresQueryName = "DATABASE_URL"

func main() {
	logger := loggerSingleton.GetLogger()
	logger.SetLevel(logrus.ErrorLevel)

	pool, err := init_db.InitPostgres(EnvPostgresQueryName)
	if err != nil {
		logger.Errorln("error connecting database: ", err)
	}

	kafkaConsumer, err := init_kafka.NewConsumer(KafkaQueryConnection)
	if err != nil {
		logger.Errorln("error connecting kafka: ", err)
	}

	mc := memcache.New("memcached:11211")
	repoTrack := track_repository.NewPostgres(pool, logger)

	activityConsumer := activity_repository_consumer.NewDefault(repoTrack, kafkaConsumer, logger)
	wavePoolRepository := wave_repository.NewRepo(pool, logger)
	recentActivityRepo := activity_repository.NewMemCached(mc, logger)

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

	activityWorkerUseCase := activity_usecase.NewWorker(&recommendationAgent, &candidateAgent, wavePoolRepository, recentActivityRepo, &activityConsumer, logger)

	activityWorkerUseCase.Run()

	select {}
}
