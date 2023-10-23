package main

import (
	consulapi "github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	init_db "main/init/postgres_db"
	init_redis "main/init/redis_db"
	user_proto "main/internal/microservices/user/proto"
	grpc_server_user "main/internal/microservices/user/service/server"
	session_repository_redis "main/internal/pkg/session/repository/redis"
	user_repository "main/internal/pkg/user/repository/postgresql"
	"net"
)

const EnvPostgresQueryName = "DATABASE_URL"

func main() {
	logger := logrus.New()
	lis, err := net.Listen("tcp", ":8081")
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

	config := consulapi.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	consul, err := consulapi.NewClient(config)
	if err != nil {
		logger.Fatalf("error while creating consul client %s", err.Error())
	}

	serviceID := "SAPI_127.0.0.1:8081"

	err = consul.Agent().ServiceRegister(&consulapi.AgentServiceRegistration{
		ID:      serviceID,
		Name:    "session-api",
		Port:    8081,
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

	logger.Infoln("starting server at 8081")
	err = server.Serve(lis)
	if err != nil {
		logger.Fatalln("terminated ERROR!!!")
	}
}
