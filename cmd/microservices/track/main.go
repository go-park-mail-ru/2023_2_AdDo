package main

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	init_kafka "main/init/kafka_queue"
	"main/internal/pkg/activity/repository/kafka/activity_repository_producer"

	// microservices_init "main/cmd/microservices"
	init_db "main/init/postgres_db"
	log "main/internal/common/logger"
	"main/internal/microservices/track/proto"
	grpc_track_server "main/internal/microservices/track/service/server"
	album_repository "main/internal/pkg/album/repository/postgres"
	artist_repository "main/internal/pkg/artist/repository/postgres"
	track_repository "main/internal/pkg/track/repository/postgresql"
	"net"
	"strconv"
)

const EnvPostgresQueryName = "DATABASE_URL"
const Port = 8083

var loggerSingleton = log.Singleton{}
var KafkaQueryConnection = []string{"kafka:9092"}

func main() {
	logger := loggerSingleton.GetLogger()
	logger.SetLevel(logrus.ErrorLevel)

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(Port))
	if err != nil {
		logger.Errorln("err while starting use micros: ", err)
	}

	pool, err := init_db.InitPostgres(EnvPostgresQueryName)
	if err != nil {
		logger.Errorln("error connecting database: ", err)
	}

	kafkaProducerClient, err := init_kafka.NewProducer(KafkaQueryConnection)
	if err != nil {
		logger.Errorln("error connecting kafka: ", err)
	}

	trackRepository := track_repository.NewPostgres(pool, logger)
	albumRepository := album_repository.NewPostgres(pool, logger)
	artistRepository := artist_repository.NewPostgres(pool, logger)
	activityProducer := activity_repository_producer.NewDefault(kafkaProducerClient, logger)

	trackManager := grpc_track_server.NewTrackManager(&activityProducer, trackRepository, &artistRepository, albumRepository, logger)

	server := grpc.NewServer()
	proto.RegisterTrackServiceServer(server, &trackManager)

	logger.Infoln("starting server at " + strconv.Itoa(Port))
	err = server.Serve(lis)
	if err != nil {
		logger.Fatalln("terminated ERROR!!!")
	}
}
