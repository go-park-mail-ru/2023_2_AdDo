package main

import (
	"google.golang.org/grpc"
	microservices_init "main/cmd/microservices"
	init_db "main/init/postgres_db"
	log "main/internal/common/logger"
	proto2 "main/internal/microservices/playlist/proto"
	grpc_playlist_server "main/internal/microservices/playlist/service/server"
	playlist_repository "main/internal/pkg/playlist/repository"
	track_repository "main/internal/pkg/track/repository/postgresql"
	"net"
	"strconv"
)

const EnvPostgresQueryName = "DATABASE_URL"

const Port = 8085

var loggerSingleton = log.Singleton{}

func main() {
	logger := loggerSingleton.GetLogger()

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(Port))
	if err != nil {
		logger.Errorln("err while starting playlist micros: ", err)
	}

	pool, err := init_db.InitPostgres(EnvPostgresQueryName)
	if err != nil {
		logger.Errorln("error connecting database: ", err)
	}

	trackRepository := track_repository.NewPostgres(pool, logger)
	playlistRepository := playlist_repository.NewPostgres(pool, logger)

	playlistManager := grpc_playlist_server.NewPlaylistManager(&playlistRepository, trackRepository, logger)

	server := grpc.NewServer()
	proto2.RegisterPlaylistServiceServer(server, &playlistManager)

	serviceId := microservices_init.RegisterInConsul(Port, "playlist-api", "playlist", logger)
	defer func() {
		microservices_init.UnRegisterInConsul(serviceId, logger)
	}()

	logger.Infoln("starting server at " + strconv.Itoa(Port))
	err = server.Serve(lis)
	if err != nil {
		logger.Fatalln("terminated ERROR!!!")
	}
}
