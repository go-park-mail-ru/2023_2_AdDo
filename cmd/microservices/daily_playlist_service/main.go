package main

import (
	"google.golang.org/grpc"
	init_db "main/init/postgres_db"
	log "main/internal/common/logger"
	daily_playlist "main/internal/microservices/daily-playlist/proto"
	daily_playlist_service_server "main/internal/microservices/daily-playlist/service/server"
	daily_playlist_repository "main/internal/pkg/daily-playlist/repository"
	"net"
	"strconv"
)

const EnvPostgresQueryName = "DATABASE_URL"
const Port = 8090

var loggerSingleton = log.Singleton{}

func main() {
	logger := loggerSingleton.GetLogger()

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(Port))
	if err != nil {
		logger.Errorln("err while starting daily playlist micros: ", err)
	}

	pool, err := init_db.InitPostgres(EnvPostgresQueryName)
	if err != nil {
		logger.Errorln("error connecting database: ", err)
	}

	dailyPlaylistRepo := daily_playlist_repository.NewPostgres(pool, logger)

	dailyPlaylistManager := daily_playlist_service_server.NewDailyManager(dailyPlaylistRepo, logger)

	server := grpc.NewServer()
	daily_playlist.RegisterDailyPlaylistServiceServer(server, &dailyPlaylistManager)

	logger.Infoln("starting server at " + strconv.Itoa(Port))
	err = server.Serve(lis)
	if err != nil {
		logger.Fatalln("terminated ERROR!!!")
	}
}
