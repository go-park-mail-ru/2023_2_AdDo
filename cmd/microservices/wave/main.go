package main

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	init_db "main/init/postgres_db"
	log "main/internal/common/logger"
	proto2 "main/internal/microservices/wave/proto"
	grpc_wave_server "main/internal/microservices/wave/service/server"
	track_repository "main/internal/pkg/track/repository/postgresql"
	wave_repository "main/internal/pkg/wave/repository/postgres"
	"net"
	"strconv"
)

const EnvPostgresQueryName = "DATABASE_URL"
const Port = 8092

var loggerSingleton = log.Singleton{}

func main() {
	logger := loggerSingleton.GetLogger()
	logger.SetLevel(logrus.ErrorLevel)

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(Port))
	if err != nil {
		logger.Errorln("err while starting wave micros: ", err)
	}

	pool, err := init_db.InitPostgres(EnvPostgresQueryName)
	if err != nil {
		logger.Errorln("error connecting database: ", err)
	}

	waveRepo := wave_repository.NewRepo(pool, logger)
	trackRepo := track_repository.NewPostgres(pool, logger)

	waveManager := grpc_wave_server.NewWaveManager(trackRepo, waveRepo, logger)

	server := grpc.NewServer()
	proto2.RegisterWaveServiceServer(server, &waveManager)

	logger.Infoln("starting server at " + strconv.Itoa(Port))
	logger.Errorln("test error for deploying check")
	err = server.Serve(lis)
	if err != nil {
		logger.Fatalln("terminated ERROR!!!")
	}
}
