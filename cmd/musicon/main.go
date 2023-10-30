package main

import (
	logger_init "main/init/logger"
	init_minio "main/init/minio"
	init_db "main/init/postgres_db"
	init_redis "main/init/redis_db"
	router_init "main/init/router"
	album_delivery "main/internal/pkg/album/delivery/http"
	album_repository "main/internal/pkg/album/repository/postgres"
	album_usecase "main/internal/pkg/album/usecase"
	artist_delivery "main/internal/pkg/artist/delivery/http"
	artist_repository "main/internal/pkg/artist/repository/postgres"
	artist_usecase "main/internal/pkg/artist/usecase"
	avatar_repository "main/internal/pkg/image/repository/minio"
	avatar_usecase "main/internal/pkg/image/usecase"
	session_repository_redis "main/internal/pkg/session/repository/redis"
	session_usecase "main/internal/pkg/session/usecase"
	track_delivery "main/internal/pkg/track/delivery/http"
	track_repository "main/internal/pkg/track/repository/postgresql"
	track_usecase "main/internal/pkg/track/usecase"
	user_delivery "main/internal/pkg/user/delivery/http"
	user_repository "main/internal/pkg/user/repository/postgresql"
	user_usecase "main/internal/pkg/user/usecase"
	"net/http"

	_ "github.com/lib/pq"
)

const EnvPostgresQueryName = "DATABASE_URL"
const ServerPort = ":8080"

func main() {
	logger := logger_init.LogRusInit()

	postgres, err := init_db.InitPostgres(EnvPostgresQueryName)
	if err != nil {
		logger.Fatalf("error postgres_db connecting %v", err)
	}
	defer postgres.Close()

	redis, err := init_redis.InitRedis()
	if err != nil {
		logger.Fatalf("error redis_db connecting %v", err)
	}
	defer redis.Close()

	minio, err := init_minio.InitMinio()
	if err != nil {
		logger.Fatalf("error minio connecting %v", err)
	}

	sessionRepo := session_repository_redis.NewRedis(redis, logger)
	userRepo := user_repository.NewPostgres(postgres, logger)
	trackRepo := track_repository.NewPostgres(postgres, logger)
	albumRepo := album_repository.NewPostgres(postgres, logger)
	artistRepo := artist_repository.NewPostgres(postgres, logger)
	avatarRepo := avatar_repository.NewMinio(minio)
	logger.Infoln("Repositories initialized")

	albumUseCase := album_usecase.NewDefault(&artistRepo, trackRepo, albumRepo, logger)
	artistUseCase := artist_usecase.NewDefault(&artistRepo, trackRepo, albumRepo, logger)
	avatarUseCase := avatar_usecase.NewDefault()
	sessionUseCase := session_usecase.NewDefault(sessionRepo, logger)
	userUseCase := user_usecase.NewWithStatefulSessions(userRepo, sessionRepo, avatarRepo, avatarUseCase, logger)
	trackUseCase := track_usecase.NewDefault(trackRepo, &artistRepo, albumRepo, logger)
	logger.Infoln("UseCases initialized")

	albumHandler := album_delivery.NewHandler(&trackUseCase, &albumUseCase, &sessionUseCase, logger)
	artistHandler := artist_delivery.NewHandler(&artistUseCase, logger)
	userHandler := user_delivery.NewHandler(&userUseCase, &sessionUseCase, logger)
	trackHandler := track_delivery.NewHandler(&trackUseCase, &sessionUseCase, logger)
	logger.Infoln("Deliveries initialized")

	router := router_init.New(userHandler, trackHandler, artistHandler, albumHandler, logger)

	logger.Fatalln(http.ListenAndServe(ServerPort, router))
}
