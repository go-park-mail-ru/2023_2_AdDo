package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	init_db "main/init/database"
	router_init "main/init/router"
	album_repository "main/internal/pkg/album/repository/postgres"
	artist_repository "main/internal/pkg/artist/repository/postgres"
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
		log.Fatalf("error database connecting %w", err)
	}
	defer db.Close()

	router := mux.NewRouter()

	sessionRepo := session_repository.NewPostgres(db)
	userRepo := user_repository.NewPostgres(db)
	trackRepo := track_repository.NewPostgres(db)
	albumRepo := album_repository.NewPostgres(db)
	artistRepo := artist_repository.NewPostgres(db)
	log.Println("Repositories initialized")

	sessionUseCase := session_usecase.NewDefault(sessionRepo)
	userUseCase := user_usecase.NewWithStatefulSessions(userRepo, sessionRepo)
	trackUseCase := track_usecase.NewDefault(trackRepo, &artistRepo, albumRepo)
	log.Println("UseCases initialized")

	userHandler := user_delivery.NewHandler(&userUseCase)
	trackHandler := track_delivery.NewHandler(&trackUseCase, &sessionUseCase)
	log.Println("Deliveries initialized")

	router = router_init.New(router, userHandler, trackHandler)

	routerCORS := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
		handlers.ExposedHeaders([]string{session.CookieName}),
		handlers.AllowCredentials(),
	)(router)

	log.Fatal(http.ListenAndServe(ServerPort, routerCORS))
}
