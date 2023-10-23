package main

import (
	consulapi "github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	init_redis "main/init/redis_db"
	session "main/internal/microservices/session/proto"
	grpc_session_server "main/internal/microservices/session/service/server"
	session_repository_redis "main/internal/pkg/session/repository/redis"
	"net"
	"strconv"
)

const EnvPostgresQueryName = "DATABASE_URL"
const Port = 8082

func main() {
	logger := logrus.New()
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(Port))
	if err != nil {
		logger.Errorln("err while starting session micros: ", err)
	}

	redis, err := init_redis.InitRedis()
	if err != nil {
		logger.Fatalf("error while conneting redis databse %s", err.Error())
	}

	sessionRepository := session_repository_redis.NewRedis(redis, logger)

	sessionManager := grpc_session_server.NewSessionManager(sessionRepository, logger)

	server := grpc.NewServer()
	session.RegisterSessionServiceServer(server, &sessionManager)

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
