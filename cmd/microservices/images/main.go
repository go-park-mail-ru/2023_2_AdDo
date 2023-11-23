package main

import (
	"google.golang.org/grpc"
	// microservices_init "main/cmd/microservices"
	init_minio "main/init/minio"
	log "main/internal/common/logger"
	image_proto "main/internal/microservices/image/proto"
	grpc_image_server "main/internal/microservices/image/service/server"
	image_repository "main/internal/pkg/image/repository/minio"
	"net"
	"strconv"
)

const Port = 8087

var loggerSingleton = log.Singleton{}

func main() {
	logger := loggerSingleton.GetLogger()

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(Port))
	if err != nil {
		logger.Errorln("err while starting images micros: ", err)
	}

	minio, err := init_minio.InitMinio()
	if err != nil {
		logger.Fatalf("error minio connecting %v", err)
	}

	imageRepository := image_repository.NewMinio(minio)

	imageManager := grpc_image_server.NewImageManager(imageRepository, logger)

	server := grpc.NewServer()
	image_proto.RegisterImageServiceServer(server, &imageManager)

	// serviceId := microservices_init.RegisterInConsul(Port, "images-api", "images", logger)
	// defer func() {
	// 	microservices_init.UnRegisterInConsul(serviceId, logger)
	// }()

	logger.Infoln("starting server at " + strconv.Itoa(Port))
	err = server.Serve(lis)
	if err != nil {
		logger.Fatalln("terminated ERROR!!!")
	}
}
