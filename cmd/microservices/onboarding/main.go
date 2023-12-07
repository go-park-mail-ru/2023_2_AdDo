package main

import (
	"google.golang.org/grpc"
	init_kafka "main/init/kafka_queue"
	init_db "main/init/postgres_db"
	log "main/internal/common/logger"
	onboarding "main/internal/microservices/onboarding/proto"
	grpc_onboarding_server "main/internal/microservices/onboarding/service/server"
	"main/internal/pkg/activity/repository/kafka/activity_repository_producer"
	onboarding_repository "main/internal/pkg/onboarding/repository/postgres"
	"net"
	"strconv"
)

const EnvPostgresQueryName = "DATABASE_URL"
const Port = 8091

var loggerSingleton = log.Singleton{}
var KafkaQueryConnection = []string{"kafka:9092"}

func main() {
	logger := loggerSingleton.GetLogger()

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(Port))
	if err != nil {
		logger.Errorln("err while starting onboarding micros: ", err)
	}

	pool, err := init_db.InitPostgres(EnvPostgresQueryName)
	if err != nil {
		logger.Errorln("error connecting database: ", err)
	}

	kafkaProducerClient, err := init_kafka.NewProducer(KafkaQueryConnection)
	if err != nil {
		logger.Errorln("error connecting kafka: ", err)
	}

	likeProducer := activity_repository_producer.NewDefault(kafkaProducerClient, logger)
	onboardingRepo := onboarding_repository.NewPostgres(pool, logger)

	onboardingManager := grpc_onboarding_server.NewOnboardingManager(&likeProducer, &onboardingRepo, logger)

	server := grpc.NewServer()
	onboarding.RegisterOnboardingServiceServer(server, &onboardingManager)

	logger.Infoln("starting server at " + strconv.Itoa(Port))
	err = server.Serve(lis)
	if err != nil {
		logger.Fatalln("terminated ERROR!!!")
	}
}
