package main

import (
	"google.golang.org/grpc"
	init_db "main/init/postgres_db"
	log "main/internal/common/logger"
	"main/internal/microservices/survey/proto"
	grpc_survey_server "main/internal/microservices/survey/service/server"
	survey_repository "main/internal/pkg/survey/repository"
	"net"
	"strconv"
)

const Port = 8088
const EnvPostgresQueryName = "DATABASE_URL"

var loggerSingleton = log.Singleton{}

func main() {
	logger := loggerSingleton.GetLogger()

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(Port))
	if err != nil {
		logger.Errorln("err while starting survey micros: ", err)
	}

	pool, err := init_db.InitPostgres(EnvPostgresQueryName)
	if err != nil {
		logger.Errorln("error connecting database: ", err)
	}

	surveyRepo := survey_repository.NewPostgres(pool, logger)

	surveyManager := grpc_survey_server.NewSurveyManager(&surveyRepo, logger)

	server := grpc.NewServer()
	proto.RegisterSurveyServiceServer(server, &surveyManager)

	// serviceId := microservices_init.RegisterInConsul(Port, "session-api", "session", logger)
	// defer func() {
	// 	microservices_init.UnRegisterInConsul(serviceId, logger)
	// }()

	logger.Infoln("starting server at " + strconv.Itoa(Port))
	err = server.Serve(lis)
	if err != nil {
		logger.Fatalln("terminated ERROR!!!")
	}
}
