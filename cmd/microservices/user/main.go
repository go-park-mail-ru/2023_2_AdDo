package main

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	microservices_init "main/cmd/microservices"
	init_minio "main/init/minio"
	init_db "main/init/postgres_db"
	init_redis "main/init/redis_db"
	user_proto "main/internal/microservices/user/proto"
	grpc_server_user "main/internal/microservices/user/service/server"
	avatar_repository "main/internal/pkg/avatar/repository/minio"
	avatar_usecase "main/internal/pkg/avatar/usecase"
	session_repository_redis "main/internal/pkg/session/repository/redis"
	user_repository "main/internal/pkg/user/repository/postgresql"
	"net"
	"strconv"
)

const EnvPostgresQueryName = "DATABASE_URL"

const Port = 8081

func main() {
	logger := logrus.New()
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

	minio, err := init_minio.InitMinio()
	if err != nil {
		logger.Fatalf("error minio connecting %v", err)
	}

	userRepository := user_repository.NewPostgres(pool, logger)
	authRepository := session_repository_redis.NewRedis(redis, logger)
	avatarRepository := avatar_repository.NewMinio(minio)

	avatarUseCase := avatar_usecase.NewDefault()
	userManager := grpc_server_user.NewUserManager(userRepository, authRepository, avatarRepository, avatarUseCase, logger)

	server := grpc.NewServer()
	user_proto.RegisterUserServiceServer(server, userManager)

	serviceId := microservices_init.RegisterInConsul(Port, "user-api", "user", logger)
	defer func() {
		microservices_init.UnRegisterInConsul(serviceId, logger)
	}()

	logger.Infoln("starting server at " + strconv.Itoa(Port))
	err = server.Serve(lis)
	if err != nil {
		logger.Fatalln("terminated ERROR!!!")
	}
}
