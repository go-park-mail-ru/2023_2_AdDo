package main

import (
	_ "github.com/lib/pq"
)

const EnvPostgresQueryName = "DATABASE_URL"
const ServerPort = ":8080"

//func main() {
//	logger := logger_init.LogRusInit()
//
//	postgres, err := init_db.InitPostgres(EnvPostgresQueryName)
//	if err != nil {
//		logger.Fatalf("error postgres_db connecting %v", err)
//	}
//	defer postgres.Close()
//
//	redis, err := init_redis.InitRedis()
//	if err != nil {
//		logger.Fatalf("error redis_db connecting %v", err)
//	}
//	defer redis.Close()
//
//	minio, err := init_minio.InitMinio()
//	if err != nil {
//		logger.Fatalf("error minio connecting %v", err)
//	}
//
//	sessionRepo := session_repository_redis.NewRedis(redis, logger)
//	trackRepo := track_repository.NewPostgres(postgres, logger)
//	albumRepo := album_repository.NewPostgres(postgres, logger)
//	artistRepo := artist_repository.NewPostgres(postgres, logger)
//	logger.Infoln("Repositories initialized")
//
//	albumUseCase := album_usecase.NewDefault(&artistRepo, trackRepo, albumRepo, logger)
//	artistUseCase := artist_usecase.NewDefault(&artistRepo, trackRepo, albumRepo, logger)
//	sessionUseCase := session_usecase.NewDefault(sessionRepo, logger)
//	trackUseCase := track_usecase.NewDefault(trackRepo, &artistRepo, albumRepo, logger)
//	logger.Infoln("UseCases initialized")
//
//	albumHandler := album_delivery.NewHandler(&trackUseCase, &albumUseCase, &sessionUseCase, logger)
//	artistHandler := artist_delivery.NewHandler(&artistUseCase, logger)
//	trackHandler := track_delivery.NewHandler(&trackUseCase, &sessionUseCase, logger)
//	logger.Infoln("Deliveries initialized")
//
//
//	logger.Fatalln(http.ListenAndServe(ServerPort, router))
//}
