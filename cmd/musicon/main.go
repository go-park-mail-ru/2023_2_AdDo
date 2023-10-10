package main

import (
	_ "github.com/lib/pq"
	logger_init "main/init/logger"
	init_db "main/init/postgres_db"
	init_redis "main/init/redis_db"
	router_init "main/init/router"
	album_repository "main/internal/pkg/album/repository/postgres"
	artist_delivery "main/internal/pkg/artist/delivery/http"
	artist_repository "main/internal/pkg/artist/repository/postgres"
	artist_usecase "main/internal/pkg/artist/usecase"
	session_repository_redis "main/internal/pkg/session/repository/redis"
	session_usecase "main/internal/pkg/session/usecase"
	track_delivery "main/internal/pkg/track/delivery/http"
	track_repository "main/internal/pkg/track/repository/postgresql"
	track_usecase "main/internal/pkg/track/usecase"
	user_delivery "main/internal/pkg/user/delivery/http"
	user_repository "main/internal/pkg/user/repository/postgresql"
	user_usecase "main/internal/pkg/user/usecase"
	"net/http"
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

	sessionRepo := session_repository_redis.NewRedis(redis)
	userRepo := user_repository.NewPostgres(postgres)
	trackRepo := track_repository.NewPostgres(postgres)
	albumRepo := album_repository.NewPostgres(postgres)
	artistRepo := artist_repository.NewPostgres(postgres)
	logger.Infoln("Repositories initialized")

	artistUseCase := artist_usecase.NewDefault(&artistRepo)
	sessionUseCase := session_usecase.NewDefault(sessionRepo)
	userUseCase := user_usecase.NewWithStatefulSessions(userRepo, sessionRepo)
	trackUseCase := track_usecase.NewDefault(trackRepo, &artistRepo, albumRepo)
	logger.Infoln("UseCases initialized")

	artistHandler := artist_delivery.NewHandler(&artistUseCase)
	userHandler := user_delivery.NewHandler(&userUseCase)
	trackHandler := track_delivery.NewHandler(&trackUseCase, &sessionUseCase)
	logger.Infoln("Deliveries initialized")

	router := router_init.New(userHandler, trackHandler, artistHandler, logger)

	logger.Fatalln(http.ListenAndServe(ServerPort, router))
}
