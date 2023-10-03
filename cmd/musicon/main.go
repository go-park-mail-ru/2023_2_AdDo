package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	init_db "main/init/database"
	logger_init "main/init/logger"
	router_init "main/init/router"
	album_repository "main/internal/pkg/album/repository/postgres"
	artist_repository "main/internal/pkg/artist/repository/postgres"
	common_middleware "main/internal/pkg/common/middleware"
	"main/internal/pkg/session"
	session_repository "main/internal/pkg/session/repository/postgresql"
	session_usecase "main/internal/pkg/session/usecase"
	track_delivery "main/internal/pkg/track/delivery/http"
	track_repository "main/internal/pkg/track/repository/postgresql"
	track_usecase "main/internal/pkg/track/usecase"
	user_delivery "main/internal/pkg/user/delivery/http"
	user_repository "main/internal/pkg/user/repository/postgresql"
	user_usecase "main/internal/pkg/user/usecase"
	"net/http"
)

const EnvPostgresQueryName = "POSTGRES_QUERY"
const ServerPort = ":8080"

func main() {
	db, err := init_db.InitPostgres(EnvPostgresQueryName)
	if err != nil {
		log.Fatalf("error database connecting %v", err)
	}
	defer db.Close()

	router := mux.NewRouter()
	logger := logger_init.LogRusInit()

	sessionRepo := session_repository.NewPostgres(db)
	userRepo := user_repository.NewPostgres(db)
	trackRepo := track_repository.NewPostgres(db)
	albumRepo := album_repository.NewPostgres(db)
	artistRepo := artist_repository.NewPostgres(db)
	logger.Infoln("Repositories initialized")

	sessionUseCase := session_usecase.NewDefault(sessionRepo)
	userUseCase := user_usecase.NewWithStatefulSessions(userRepo, sessionRepo)
	trackUseCase := track_usecase.NewDefault(trackRepo, &artistRepo, albumRepo)
	logger.Infoln("UseCases initialized")

	userHandler := user_delivery.NewHandler(&userUseCase)
	trackHandler := track_delivery.NewHandler(&trackUseCase, &sessionUseCase)
	logger.Infoln("Deliveries initialized")

	router = router_init.New(router, userHandler, trackHandler)

	routerCORS := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://82.146.45.164:8081"}),
		handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
		handlers.ExposedHeaders([]string{session.CookieName}),
		handlers.AllowCredentials(),
	)(router)

	log.Fatal(http.ListenAndServe(ServerPort, common_middleware.Logging(routerCORS, logger)))
}
