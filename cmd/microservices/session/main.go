package main

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	// microservices_init "main/cmd/microservices"
	init_redis "main/init/redis_db"
	log "main/internal/common/logger"
	session "main/internal/microservices/session/proto"
	grpc_session_server "main/internal/microservices/session/service/server"
	session_repository_redis "main/internal/pkg/session/repository/redis"
	"net"
	"strconv"
)

const Port = 8082

var loggerSingleton = log.Singleton{}

func main() {
	logger := loggerSingleton.GetLogger()
	logger.SetLevel(logrus.ErrorLevel)

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

	logger.Infoln("starting server at " + strconv.Itoa(Port))
	err = server.Serve(lis)
	if err != nil {
		logger.Fatalln("terminated ERROR!!!")
	}
}
