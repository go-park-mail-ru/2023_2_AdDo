package router_init

import (
	"github.com/gorilla/csrf"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	_ "main/api/openapi"
	album_delivery "main/internal/pkg/album/delivery/http"
	artist_delivery "main/internal/pkg/artist/delivery/http"
	common_handler "main/internal/pkg/common/handler"
	common_middleware "main/internal/pkg/common/middleware"
	"main/internal/pkg/session"
	track_delivery "main/internal/pkg/track/delivery/http"
	user_delivery "main/internal/pkg/user/delivery/http"
	"net/http"
)

//	@title			MusicOn API
//	@version		1.0
//	@description	Music web app

//	@termsOfService	http://swagger.io/terms/

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@SecurityDefinitions.apikey	cookieAuth
//	@in							header
//	@name						JSESSIONID

//	@SecurityDefinitions.apikey	csrfToken
//	@in							header
//	@name						X-Csrf-Token

// @host		musicon.space
// @BasePath	/api/v1
func New(userHandler user_delivery.UserHandler, trackHandler track_delivery.TrackHandler, artistHandler artist_delivery.ArtistHandler, albumHandler album_delivery.AlbumHandler, logger *logrus.Logger) http.Handler {
	router := mux.NewRouter()

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	router.Handle("/api/v1/sign_up", common_handler.Handler{H: userHandler.SignUp}).Methods("POST")
	router.Handle("/api/v1/login", common_handler.Handler{H: userHandler.Login}).Methods("POST")
	router.Handle("/api/v1/logout", common_handler.Handler{H: userHandler.LogOut}).Methods("POST")
	router.Handle("/api/v1/listen/{id}", common_handler.Handler{H: trackHandler.Listen}).Methods("POST")
	router.Handle("/api/v1/like/{id}", common_handler.Handler{H: trackHandler.Like}).Methods("POST")

	router.Handle("/api/v1/artist/{id}", common_handler.Handler{H: artistHandler.ArtistInfo}).Methods("GET")
	router.Handle("/api/v1/album/{id}", common_handler.Handler{H: albumHandler.AlbumTracks}).Methods("GET")

	router.Handle("/api/v1/auth", common_handler.Handler{H: userHandler.Auth}).Methods("GET")
	router.Handle("/api/v1/me", common_handler.Handler{H: userHandler.Me}).Methods("GET")
	router.Handle("/api/v1/feed", common_handler.Handler{H: albumHandler.Feed}).Methods("GET")
	router.Handle("/api/v1/new", common_handler.Handler{H: albumHandler.New}).Methods("GET")
	router.Handle("/api/v1/most_liked", common_handler.Handler{H: albumHandler.MostLiked}).Methods("GET")
	router.Handle("/api/v1/popular", common_handler.Handler{H: albumHandler.Popular}).Methods("GET")

	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://82.146.45.164:8081"}),
		handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "X-Csrf-Token"}),
		handlers.ExposedHeaders([]string{session.CookieName}),
		handlers.AllowCredentials(),
	)

	csrfMiddleware := csrf.Protect(
		session.CSRFKey,
		csrf.Secure(false),
		csrf.HttpOnly(false),
		csrf.MaxAge(session.TimeToLiveCSRF),
		csrf.RequestHeader("X-Csrf-Token"),
		csrf.CookieName("X-Csrf-Token"),
		csrf.FieldName("X-Csrf-Token"),
	)
	router.Use(corsMiddleware, csrfMiddleware)

	routerWithMiddleware := common_middleware.Logging(router, logger)
	routerWithMiddleware = common_middleware.PanicRecovery(routerWithMiddleware)

	return routerWithMiddleware
}
