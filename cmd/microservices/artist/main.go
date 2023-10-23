package main

import (
	consulapi "github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	init_db "main/init/postgres_db"
	artist "main/internal/microservices/artist/proto"
	grpc_artist_server "main/internal/microservices/artist/service/server"
	album_repository "main/internal/pkg/album/repository/postgres"
	artist_repository "main/internal/pkg/artist/repository/postgres"
	track_repository "main/internal/pkg/track/repository/postgresql"
	"net"
	"strconv"
)

const EnvPostgresQueryName = "DATABASE_URL"

const Port = 8086

func main() {
	logger := logrus.New()
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(Port))
	if err != nil {
		logger.Errorln("err while starting artist micros: ", err)
	}

	pool, err := init_db.InitPostgres(EnvPostgresQueryName)
	if err != nil {
		logger.Errorln("error connecting database: ", err)
	}

	trackRepository := track_repository.NewPostgres(pool, logger)
	albumRepository := album_repository.NewPostgres(pool, logger)
	artistRepository := artist_repository.NewPostgres(pool, logger)

	artistManager := grpc_artist_server.NewArtistManager(&artistRepository, trackRepository, albumRepository, logger)

	server := grpc.NewServer()
	artist.RegisterArtistServiceServer(server, &artistManager)

	config := consulapi.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	consul, err := consulapi.NewClient(config)
	if err != nil {
		logger.Fatalf("error while creating consul client %s", err.Error())
	}

	serviceID := "SAPI_127.0.0.1:" + strconv.Itoa(Port)

	err = consul.Agent().ServiceRegister(&consulapi.AgentServiceRegistration{
		ID:      serviceID,
		Name:    "session-api",
		Port:    Port,
		Address: "127.0.0.1",
	})
	if err != nil {
		logger.Fatalln("cant add service to consul", err)
	}

	logger.Infoln("registered in consul", serviceID)

	defer func() {
		err := consul.Agent().ServiceDeregister(serviceID)
		if err != nil {
			logger.Fatalln("cant delete service from consul", err)
		}
		logger.Fatalf("sevice deleted from consul %s", serviceID)
	}()

	logger.Infoln("starting server at " + strconv.Itoa(Port))
	err = server.Serve(lis)
	if err != nil {
		logger.Fatalln("terminated ERROR!!!")
	}
}
