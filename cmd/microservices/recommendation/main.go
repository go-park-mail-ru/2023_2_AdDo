package main

import (
	"google.golang.org/grpc"
	log "main/internal/common/logger"
	"main/internal/microservices/recommendation/proto"
	grpc_recommendation_server "main/internal/microservices/recommendation/service/server"
	"net"
	"strconv"
)

const EnvPostgresQueryName = "DATABASE_URL"
const Port = 8089

var loggerSingleton = log.Singleton{}

func main() {
	logger := loggerSingleton.GetLogger()

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(Port))
	if err != nil {
		logger.Errorln("err while starting recommendation micros: ", err)
	}
	//
	//pool, err := init_db.InitPostgres(EnvPostgresQueryName)
	//if err != nil {
	//	logger.Errorln("error connecting database: ", err)
	//}

	//waveRepo := wave_repository.NewPostgres(pool, logger)

	recManager := grpc_recommendation_server.NewRecommendationManager(logger)

	server := grpc.NewServer()
	proto.RegisterRecommendationServiceServer(server, &recManager)

	logger.Infoln("starting server at " + strconv.Itoa(Port))
	err = server.Serve(lis)
	if err != nil {
		logger.Fatalln("terminated ERROR!!!")
	}
}
