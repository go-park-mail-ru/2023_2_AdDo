package main

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	init_db "main/init/postgres_db"
	init_redis "main/init/redis_db"
	log "main/internal/common/logger"
	user_proto "main/internal/microservices/user/proto"
	grpc_server_user "main/internal/microservices/user/service/server"
	session_repository_redis "main/internal/pkg/session/repository/redis"
	user_repository "main/internal/pkg/user/repository/postgresql"
	"net"
	"strconv"
)

const EnvPostgresQueryName = "DATABASE_URL"

const Port = 8081

var loggerSingleton = log.Singleton{}

func main() {
	logger := loggerSingleton.GetLogger()
	logger.SetLevel(logrus.ErrorLevel)

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(Port))
	if err != nil {
		logger.Errorln("err while starting use micros: ", err)
	}

	pool, err := init_db.InitPostgres(EnvPostgresQueryName)
	if err != nil {
		logger.Errorln("error connecting database: ", err)
	}

	redis, err := init_redis.InitRedis()
	if err != nil {
		logger.Fatalf("error while conneting redis databse %s", err.Error())
	}

	userRepository := user_repository.NewPostgres(pool, logger)
	authRepository := session_repository_redis.NewRedis(redis, logger)

	userManager := grpc_server_user.NewUserManager(userRepository, authRepository, logger)

	server := grpc.NewServer()
	user_proto.RegisterUserServiceServer(server, userManager)

	logger.Infoln("starting server at " + strconv.Itoa(Port))
	err = server.Serve(lis)
	if err != nil {
		logger.Fatalln("terminated ERROR!!!")
	}
}
