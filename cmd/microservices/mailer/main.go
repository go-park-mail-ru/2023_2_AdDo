package main

import (
	init_redis "main/init/redis_db"
	log "main/internal/common/logger"
	mailer_proto "main/internal/microservices/mailer/proto"
	grpc_server "main/internal/microservices/mailer/service/server"
	domain "main/internal/pkg/mailer"
	mailer_repository_redis "main/internal/pkg/mailer/repository/redis"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

const Port = 8088

var loggerSingleton = log.Singleton{}

func main() {
	logger := loggerSingleton.GetLogger()

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(Port))
	if err != nil {
		logger.Errorln("err while starting use micros: ", err)
	}

	redis, err := init_redis.InitRedis()
	if err != nil {
		logger.Fatalf("error while conneting redis databse %s", err.Error())
	}

	mailerRepo := mailer_repository_redis.NewRedis(redis, logger)

	// TODO: add smtp config
	smtp := domain.Smtp{}

	mailerManager := grpc_server.NewMailerServer(smtp, mailerRepo, logger)

	server := grpc.NewServer()
	mailer_proto.RegisterMailerServiceServer(server, mailerManager)

	logger.Infoln("starting server at " + strconv.Itoa(Port))
	err = server.Serve(lis)
	if err != nil {
		logger.Fatalln("terminated ERROR!!!")
	}
}
