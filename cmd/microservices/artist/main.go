package main

import (
	"google.golang.org/grpc"
	init_kafka "main/init/kafka_queue"
	"main/internal/pkg/activity/repository/kafka/activity_repository_producer"

	// microservices_init "main/cmd/microservices"
	init_db "main/init/postgres_db"
	log "main/internal/common/logger"
	artist "main/internal/microservices/artist/proto"
	grpc_artist_server "main/internal/microservices/artist/service/server"
	album_repository "main/internal/pkg/album/repository/postgres"
	artist_repository "main/internal/pkg/artist/repository/postgres"
	playlist_repository "main/internal/pkg/playlist/repository"
	track_repository "main/internal/pkg/track/repository/postgresql"
	"net"
	"strconv"
)

const EnvPostgresQueryName = "DATABASE_URL"
const Port = 8086

var KafkaQueryConnection = []string{"kafka:9092"}
var loggerSingleton = log.Singleton{}

func main() {
	logger := loggerSingleton.GetLogger()

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(Port))
	if err != nil {
		logger.Errorln("err while starting artist micros: ", err)
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
	playlistRepository := playlist_repository.NewPostgres(pool, logger)
	activityProducer := activity_repository_producer.NewDefault(kafkaProducerClient, logger)

	artistManager := grpc_artist_server.NewArtistManager(&activityProducer, &playlistRepository, &artistRepository, trackRepository, albumRepository, logger)

	server := grpc.NewServer()
	artist.RegisterArtistServiceServer(server, &artistManager)

	logger.Infoln("starting server at " + strconv.Itoa(Port))
	err = server.Serve(lis)
	if err != nil {
		logger.Fatalln("terminated ERROR!!!")
	}
}
