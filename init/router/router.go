package router_init

import (
	"github.com/gorilla/csrf"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	_ "main/api/openapi"
	"main/internal/common/handler"
	"main/internal/common/middleware/common"
	modify_playlist "main/internal/common/middleware/playlist_middleware/modify"
	read_playlist "main/internal/common/middleware/playlist_middleware/read"
	album_delivery "main/internal/pkg/album/delivery/http"
	artist_delivery "main/internal/pkg/artist/delivery/http"
	"main/internal/pkg/playlist"
	playlist_delivery "main/internal/pkg/playlist/delivery/http"
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
//	@in							cookie
//	@name						JSESSIONID

//	@SecurityDefinitions.apikey	csrfToken
//	@in							header
//	@name						X-Csrf-Token

//	@SecurityDefinitions.apikey	cookieCsrfToken
//	@in							cookie
//	@name						X-Csrf-Token

// @host		musicon.space
// @BasePath	/api/v1
func New(playlistUseCase playlist.UseCase, sessionUseCase session.UseCase, userHandler user_delivery.UserHandler, trackHandler track_delivery.TrackHandler, artistHandler artist_delivery.ArtistHandler, albumHandler album_delivery.AlbumHandler, playlistHandler playlist_delivery.Handler, logger *logrus.Logger) http.Handler {
	router := mux.NewRouter()

	playlistRouterModify := router.PathPrefix("").Subrouter()
	playlistRouterRead := router.PathPrefix("").Subrouter()

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	router.Handle("/api/v1/sign_up", common_handler.Handler{H: userHandler.SignUp}).Methods("POST")
	router.Handle("/api/v1/login", common_handler.Handler{H: userHandler.Login}).Methods("POST")
	router.Handle("/api/v1/update_info", common_handler.Handler{H: userHandler.UpdateUserInfo}).Methods("PUT")
	router.Handle("/api/v1/upload_avatar", common_handler.Handler{H: userHandler.UploadAvatar}).Methods("POST")
	router.Handle("/api/v1/remove_avatar", common_handler.Handler{H: userHandler.RemoveAvatar}).Methods("POST")
	router.Handle("/api/v1/auth", common_handler.Handler{H: userHandler.Auth}).Methods("GET")
	router.Handle("/api/v1/me", common_handler.Handler{H: userHandler.Me}).Methods("GET")
	router.Handle("/api/v1/logout", common_handler.Handler{H: userHandler.LogOut}).Methods("POST")

	router.Handle("/api/v1/listen", common_handler.Handler{H: trackHandler.Listen}).Methods("POST")
	router.Handle("/api/v1/track/{id}/like", common_handler.Handler{H: trackHandler.Like}).Methods("POST")
	router.Handle("/api/v1/track/{id}/is_like", common_handler.Handler{H: trackHandler.IsLike}).Methods("GET")
	router.Handle("/api/v1/track/{id}/unlike", common_handler.Handler{H: trackHandler.Unlike}).Methods("DELETE")
	router.Handle("/api/v1/collection/tracks", common_handler.Handler{H: trackHandler.GetUserTracks}).Methods("GET")

	router.Handle("/api/v1/feed", common_handler.Handler{H: albumHandler.Feed}).Methods("GET")
	router.Handle("/api/v1/new", common_handler.Handler{H: albumHandler.New}).Methods("GET")
	router.Handle("/api/v1/most_liked", common_handler.Handler{H: albumHandler.MostLiked}).Methods("GET")
	router.Handle("/api/v1/popular", common_handler.Handler{H: albumHandler.Popular}).Methods("GET")
	router.Handle("/api/v1/album/{id}/like", common_handler.Handler{H: albumHandler.Like}).Methods("POST")
	router.Handle("/api/v1/album/{id}/is_like", common_handler.Handler{H: albumHandler.IsLike}).Methods("GET")
	router.Handle("/api/v1/album/{id}/unlike", common_handler.Handler{H: albumHandler.Unlike}).Methods("DELETE")
	router.Handle("/api/v1/album/{id}", common_handler.Handler{H: albumHandler.AlbumTracks}).Methods("GET")

	router.Handle("/api/v1/artist/{id}/like", common_handler.Handler{H: artistHandler.Like}).Methods("POST")
	router.Handle("/api/v1/artist/{id}/is_like", common_handler.Handler{H: artistHandler.IsLike}).Methods("GET")
	router.Handle("/api/v1/artist/{id}/unlike", common_handler.Handler{H: artistHandler.Unlike}).Methods("DELETE")
	router.Handle("/api/v1/artist/{id}", common_handler.Handler{H: artistHandler.ArtistInfo}).Methods("GET")

	router.Handle("/api/v1/playlist", common_handler.Handler{H: playlistHandler.Create}).Methods("POST")

	playlistRouterModify.Handle("/api/v1/playlist/{id}", common_handler.Handler{H: playlistHandler.Delete}).Methods("DELETE")
	playlistRouterModify.Handle("/api/v1/playlist/{id}/add_track", common_handler.Handler{H: playlistHandler.AddTrack}).Methods("POST")
	playlistRouterModify.Handle("/api/v1/playlist/{id}/remove_track", common_handler.Handler{H: playlistHandler.RemoveTrack}).Methods("DELETE")
	playlistRouterModify.Handle("/api/v1/playlist/{id}/make_private", common_handler.Handler{H: playlistHandler.MakePrivate}).Methods("PUT")
	playlistRouterModify.Handle("/api/v1/playlist/{id}/make_public", common_handler.Handler{H: playlistHandler.MakePublic}).Methods("PUT")
	playlistRouterModify.Handle("/api/v1/playlist/{id}/update_preview", common_handler.Handler{H: playlistHandler.UpdatePreview}).Methods("POST")

	playlistRouterRead.Handle("/api/v1/playlist/{id}/like", common_handler.Handler{H: playlistHandler.Like}).Methods("POST")
	playlistRouterRead.Handle("/api/v1/playlist/{id}/is_like", common_handler.Handler{H: playlistHandler.IsLike}).Methods("GET")
	playlistRouterRead.Handle("/api/v1/playlist/{id}/unlike", common_handler.Handler{H: playlistHandler.Unlike}).Methods("DELETE")
	playlistRouterRead.Handle("/api/v1/playlist/{id}", common_handler.Handler{H: playlistHandler.Get}).Methods("GET")

	modifyPlaylistMiddleware := modify_playlist.NewMiddleware(playlistUseCase, sessionUseCase, logger)
	readPlaylistMiddleware := read_playlist.NewMiddleware(playlistUseCase, logger)

	playlistRouterModify.Use(modifyPlaylistMiddleware.ModifyPlaylistAccess)
	playlistRouterRead.Use(readPlaylistMiddleware.ReadPlaylistAccess)

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

	routerWithMiddleware := common.Logging(router, logger)
	routerWithMiddleware = common.PanicRecovery(routerWithMiddleware, logger)

	return routerWithMiddleware
}

/// TODO написать ручки на все коллекции: треки, артисты, плейлисты и альбомы, все коллекции возвращают имя владельца и его фото, тесты на них
/// TODO полнотекстовый поиск по артистам, трекам и альбомам среди имен, поиск работает в реальном времени
/// TODO easyjson для сериализации и десериализации
