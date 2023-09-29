package main

import (
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	init_db "main/init/database"
	router_init "main/init/router"
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
	log.Println("Repositories initialized")

	sessionUsecase := session_usecase.NewDefault(sessionRepo)
	userUsecase := user_usecase.NewWithStatefulSessions(userRepo, sessionRepo)
	trackUsecase := track_usecase.NewDefault(trackRepo)
	log.Println("UseCases initialized")

	userHandler := user_delivery.NewHandler(&userUsecase)
	trackHandler := track_delivery.NewHandler(&trackUsecase, &sessionUsecase)
	log.Println("Deliveries initialized")

	router = router_init.New(router, userHandler, trackHandler)

	log.Fatal(http.ListenAndServe(ServerPort, router))
}
