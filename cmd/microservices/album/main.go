package main

import (
	"google.golang.org/grpc"
	// microservices_init "main/cmd/microservices"
	init_db "main/init/postgres_db"
	log "main/internal/common/logger"
	proto2 "main/internal/microservices/album/proto"
	grpc_album_server "main/internal/microservices/album/service/server"
	album_repository "main/internal/pkg/album/repository/postgres"
	artist_repository "main/internal/pkg/artist/repository/postgres"
	track_repository "main/internal/pkg/track/repository/postgresql"
	"net"
	"strconv"
)

const EnvPostgresQueryName = "DATABASE_URL"
const Port = 8084

var loggerSingleton = log.Singleton{}

func main() {
	logger := loggerSingleton.GetLogger()

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(Port))
	if err != nil {
		logger.Errorln("err while starting album micros: ", err)
	}

	pool, err := init_db.InitPostgres(EnvPostgresQueryName)
	if err != nil {
		logger.Errorln("error connecting database: ", err)
	}

	trackRepository := track_repository.NewPostgres(pool, logger)
	albumRepository := album_repository.NewPostgres(pool, logger)
	artistRepository := artist_repository.NewPostgres(pool, logger)

	albumManager := grpc_album_server.NewAlbumManager(trackRepository, &artistRepository, albumRepository, logger)

	server := grpc.NewServer()
	proto2.RegisterAlbumServiceServer(server, &albumManager)

	logger.Infoln("starting server at " + strconv.Itoa(Port))
	if err = server.Serve(lis); err != nil {
		logger.Fatalln("terminated ERROR!!!")
	}
}
