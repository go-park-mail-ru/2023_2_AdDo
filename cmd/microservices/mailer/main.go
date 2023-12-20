package main

import (
	init_redis "main/init/redis_db"
	log "main/internal/common/logger"
	mailer_proto "main/internal/microservices/mailer/proto"
	grpc_server "main/internal/microservices/mailer/service/server"
	domain "main/internal/pkg/mailer"
	mailer_repository_redis "main/internal/pkg/mailer/repository/redis"
	"net"
	"os"
	"strconv"

	"google.golang.org/grpc"
)

const Port = 8088

const (
	EnvSmtpServerPort     = "SMTP_SERVER_PORT"
	EnvSmtpServerHost     = "SMTP_SERVER_HOST"
	EnvSmtpServerUsername = "SMTP_SERVER_USERNAME"
	EnvSmtpServerPassword = "SMTP_SERVER_PASSWORD"
	EnvSmtpServerSender   = "SMTP_SERVER_SENDER"
)

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

	port, err := strconv.Atoi(os.Getenv(EnvSmtpServerPort))
	if err != nil {
		logger.Fatalf("error get env var %s: %s", EnvSmtpServerPort, err.Error())
	}

	smtp := domain.Smtp{
		Port:     port,
		Host:     os.Getenv(EnvSmtpServerHost),
		Username: os.Getenv(EnvSmtpServerUsername),
		Password: os.Getenv(EnvSmtpServerPassword),
		Sender:   os.Getenv(EnvSmtpServerSender),
	}

	mailerManager := grpc_server.NewMailerServer(smtp, mailerRepo, logger)

	server := grpc.NewServer()
	mailer_proto.RegisterMailerServiceServer(server, mailerManager)

	logger.Infoln("starting server at " + strconv.Itoa(Port))
	err = server.Serve(lis)
	if err != nil {
		logger.Fatalln("terminated ERROR!!!")
	}
}
