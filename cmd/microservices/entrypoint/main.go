package main

import (
	"main/init/middleware"
	router_init "main/init/router"
	csrf "main/internal/common/get_csrf"
	log "main/internal/common/logger"
	"main/internal/common/middleware/metrics"
	modify_playlist "main/internal/common/middleware/playlist_middleware/modify"
	read_playlist "main/internal/common/middleware/playlist_middleware/read"
	proto2 "main/internal/microservices/album/proto"
	grpc_album "main/internal/microservices/album/service/client"
	artist "main/internal/microservices/artist/proto"
	grpc_artist "main/internal/microservices/artist/service/client"
	daily_playlist "main/internal/microservices/daily-playlist/proto"
	grpc_daily_playlist "main/internal/microservices/daily-playlist/service/client"
	proto4 "main/internal/microservices/image/proto"
	grpc_image "main/internal/microservices/image/service/client"
	proto6 "main/internal/microservices/mailer/proto"
	grpc_mailer "main/internal/microservices/mailer/service/client"
	onboarding "main/internal/microservices/onboarding/proto"
	onboarding_service_client "main/internal/microservices/onboarding/service/client"
	proto3 "main/internal/microservices/playlist/proto"
	grpc_playlist "main/internal/microservices/playlist/service/client"
	session2 "main/internal/microservices/session/proto"
	grpc_session "main/internal/microservices/session/service/client"
	"main/internal/microservices/track/proto"
	grpc_track "main/internal/microservices/track/service/client"
	user_client "main/internal/microservices/user/proto"
	grpc_user "main/internal/microservices/user/service/client"
	proto5 "main/internal/microservices/wave/proto"
	grpc_wave_client "main/internal/microservices/wave/service/client"
	album_delivery "main/internal/pkg/album/delivery/http"
	artist_delivery "main/internal/pkg/artist/delivery/http"
	daily_playlist_delivery "main/internal/pkg/daily-playlist/delivery/http"
	onboarding_delivery "main/internal/pkg/onboarding/delivery"
	playlist_delivery "main/internal/pkg/playlist/delivery/http"
	track_delivery "main/internal/pkg/track/delivery/http"
	user_delivery "main/internal/pkg/user/delivery/http"
	websocket_wave "main/internal/pkg/wave/delivery/websocket"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"

	"net/http"

	"google.golang.org/grpc/credentials/insecure"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
)

const EnvPostgresQueryName = "DATABASE_URL"
const ServerPort = ":8080"

var loggerSingleton = log.Singleton{}

func main() {
	logger := loggerSingleton.GetLogger()
	logger.SetLevel(logrus.ErrorLevel)

	userConnection, err := grpc.Dial("user:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatalln("error connecting to user micros ", err)
	}

	sessionConnection, err := grpc.Dial("session:8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatalln("error connecting to session micros ", err)
	}

	trackConnection, err := grpc.Dial("track:8083", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatalln("error connecting to track micros ", err)
	}

	albumConnection, err := grpc.Dial("album:8084", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatalln("error connecting to album micros ", err)
	}

	playlistConnection, err := grpc.Dial("playlist:8085", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatalln("error connecting to playlist micros ", err)
	}

	artistConnection, err := grpc.Dial("artist:8086", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatalln("error connecting to artist micros ", err)
	}

	imageConnection, err := grpc.Dial("images:8087", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatalln("error connecting to images micros ", err)
	}

	dailyPlaylistConnection, err := grpc.Dial("daily-playlist:8090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatalln("error connecting to daily playlist micros ", err)
	}

	onboardingConnection, err := grpc.Dial("onboarding:8091", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatalln("error connecting to onboarding micros ", err)
	}

	waveConnection, err := grpc.Dial("wave:8092", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatalln("error connecting to wave micros ", err)
	}

	mailerConnection, err := grpc.Dial("mailer:8088", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatalln("error connecting to images micros ", err)
	}

	mailerAgent := grpc_mailer.NewClient(proto6.NewMailerServiceClient(mailerConnection), logger)
	imageAgent := grpc_image.NewClient(proto4.NewImageServiceClient(imageConnection), logger)
	userAgent := grpc_user.NewClient(user_client.NewUserServiceClient(userConnection), imageAgent, mailerAgent, logger)
	sessionAgent := grpc_session.NewClient(session2.NewSessionServiceClient(sessionConnection), logger)
	trackAgent := grpc_track.NewClient(proto.NewTrackServiceClient(trackConnection), logger)
	albumAgent := grpc_album.NewClient(proto2.NewAlbumServiceClient(albumConnection), logger)
	playlistAgent := grpc_playlist.NewClient(userAgent, proto3.NewPlaylistServiceClient(playlistConnection), imageAgent, logger)
	artistAgent := grpc_artist.NewClient(artist.NewArtistServiceClient(artistConnection), logger)
	onboardingAgent := onboarding_service_client.NewClient(onboarding.NewOnboardingServiceClient(onboardingConnection), logger)
	dailyPlaylistAgent := grpc_daily_playlist.NewClient(daily_playlist.NewDailyPlaylistServiceClient(dailyPlaylistConnection), logger)
	waveAgent := grpc_wave_client.NewClient(proto5.NewWaveServiceClient(waveConnection), logger)
	logger.Infoln("Clients to micros initialized")

	allowedOrigins := []string{
		"https://musicon.space",
	}

	ws := websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		for _, allowedOrigin := range allowedOrigins {
			if allowedOrigin == origin {
				return true
			}
		}
		return false
	}}

	albumHandler := album_delivery.NewHandler(&trackAgent, &albumAgent, &sessionAgent, logger)
	artistHandler := artist_delivery.NewHandler(&sessionAgent, &artistAgent, &trackAgent, logger)
	userHandler := user_delivery.NewHandler(&userAgent, &sessionAgent, logger)
	trackHandler := track_delivery.NewHandler(&trackAgent, &sessionAgent, logger)
	playlistHandler := playlist_delivery.NewHandler(&playlistAgent, &trackAgent, &sessionAgent, logger)
	onboardingHandler := onboarding_delivery.NewHandler(&sessionAgent, &onboardingAgent, logger)
	dailyPlaylistHandler := daily_playlist_delivery.NewHandler(&dailyPlaylistAgent, &sessionAgent, logger)
	waveHandler := websocket_wave.NewHandler(ws, &trackAgent, &sessionAgent, &waveAgent, logger)
	logger.Infoln("Deliveries initialized")

	modifyPlaylistMiddleware := modify_playlist.NewMiddleware(&playlistAgent, &sessionAgent, logger)
	readPlaylistMiddleware := read_playlist.NewMiddleware(&playlistAgent, logger)
	corsMiddleware := middleware.NewCors()
	//csrfMiddleware := middleware.NewCSRF()

	prometheusRegistry := prometheus.NewRegistry()
	metricsMiddleware := metrics.NewMiddleware(metrics.NewHandlers(), metrics.NewMetrics(prometheusRegistry))

	routerConfig := router_init.Config{
		Routes: []router_init.Route{
			router_init.NewRoute("/get_csrf", csrf.GetCSRF, http.MethodGet),
			router_init.NewRoute("/sign_up", userHandler.SignUp, http.MethodPost),
			router_init.NewRoute("/login", userHandler.Login, http.MethodPost),
			router_init.NewRoute("/update_info", userHandler.UpdateUserInfo, http.MethodPut),
			router_init.NewRoute("/upload_avatar", userHandler.UploadAvatar, http.MethodPost),
			router_init.NewRoute("/remove_avatar", userHandler.RemoveAvatar, http.MethodPost),
			router_init.NewRoute("/auth", userHandler.Auth, http.MethodGet),
			router_init.NewRoute("/auth/forgot_pasword", userHandler.ForgotPassword, http.MethodPost),
			router_init.NewRoute("/auth/reset_password/{reset_token}", userHandler.ResetPassword, http.MethodPost),
			router_init.NewRoute("/me", userHandler.Me, http.MethodGet),
			router_init.NewRoute("/logout", userHandler.LogOut, http.MethodPost),
			router_init.NewRoute("/artists", onboardingHandler.GetArtists, http.MethodGet),
			router_init.NewRoute("/artists", onboardingHandler.SaveArtists, http.MethodPost),
			router_init.NewRoute("/genres", onboardingHandler.GetGenres, http.MethodGet),
			router_init.NewRoute("/genres", onboardingHandler.SaveGenres, http.MethodPost),
			router_init.NewRoute("/daily", dailyPlaylistHandler.GetDailyPlaylist, http.MethodGet),
			//router_init.NewRoute("/wave", waveHandler.MyWave, http.MethodGet),

			router_init.NewRoute("/listen/{id}", trackHandler.Listen, http.MethodPost),
			router_init.NewRoute("/track/{id}", albumHandler.AlbumWithRequiredTrack, http.MethodGet),
			router_init.NewRoute("/track/{id}/like", trackHandler.Like, http.MethodPost),
			router_init.NewRoute("/track/{id}/is_like", trackHandler.IsLike, http.MethodGet),
			router_init.NewRoute("/track/{id}/unlike", trackHandler.Unlike, http.MethodDelete),
			router_init.NewRoute("/collection/tracks", trackHandler.GetUserTracks, http.MethodGet),
			router_init.NewRoute("/collection/albums", albumHandler.CollectionAlbum, http.MethodGet),
			router_init.NewRoute("/collection/artists", artistHandler.CollectionArtist, http.MethodGet),
			router_init.NewRoute("/collection/playlists", playlistHandler.CollectionPlaylist, http.MethodGet),

			router_init.NewRoute("/feed", albumHandler.Feed, http.MethodGet),
			router_init.NewRoute("/new", albumHandler.New, http.MethodGet),
			router_init.NewRoute("/most_liked", albumHandler.MostLiked, http.MethodGet),
			router_init.NewRoute("/popular", albumHandler.Popular, http.MethodGet),
			router_init.NewRoute("/album/{id}/like", albumHandler.Like, http.MethodPost),
			router_init.NewRoute("/album/{id}/is_like", albumHandler.IsLike, http.MethodGet),
			router_init.NewRoute("/album/{id}/unlike", albumHandler.Unlike, http.MethodDelete),
			router_init.NewRoute("/album/{id}", albumHandler.AlbumTracks, http.MethodGet),

			router_init.NewRoute("/artist/{id}/like", artistHandler.Like, http.MethodPost),
			router_init.NewRoute("/artist/{id}/is_like", artistHandler.IsLike, http.MethodGet),
			router_init.NewRoute("/artist/{id}/unlike", artistHandler.Unlike, http.MethodDelete),
			router_init.NewRoute("/artist/{id}", artistHandler.ArtistInfo, http.MethodGet),
			router_init.NewRoute("/search", artistHandler.FullSearch, http.MethodGet),
			router_init.NewRoute("/playlist", playlistHandler.Create, http.MethodPost),
			router_init.NewRoute("/my_playlists", playlistHandler.UserPlaylists, http.MethodGet),
		},
		Prefix: "/api/v1",
		Middlewares: []mux.MiddlewareFunc{
			corsMiddleware, metricsMiddleware.Collecting,
		},
		PrometheusRegistry: prometheusRegistry,

		SubRouterConfigs: []router_init.Config{
			router_init.Config{
				Routes: []router_init.Route{
					router_init.NewRoute("/playlist/{id}/like", playlistHandler.Like, http.MethodPost),
					router_init.NewRoute("/playlist/{id}/is_like", playlistHandler.IsLike, http.MethodGet),
					router_init.NewRoute("/playlist/{id}/unlike", playlistHandler.Unlike, http.MethodDelete),
					router_init.NewRoute("/playlist/{id}", playlistHandler.Get, http.MethodGet),
					router_init.NewRoute("/playlist/{id}/is_creator", playlistHandler.IsCreator, http.MethodGet),
				},
				Prefix: "",
				Middlewares: []mux.MiddlewareFunc{
					readPlaylistMiddleware.ReadPlaylistAccess,
				},
				SubRouterConfigs: nil,
			},

			router_init.Config{
				Routes: []router_init.Route{
					router_init.NewRoute("/playlist/{id}", playlistHandler.Delete, http.MethodDelete),
					router_init.NewRoute("/playlist/{id}/add_track", playlistHandler.AddTrack, http.MethodPost),
					router_init.NewRoute("/playlist/{id}/remove_track", playlistHandler.RemoveTrack, http.MethodDelete),
					router_init.NewRoute("/playlist/{id}/make_private", playlistHandler.MakePrivate, http.MethodPut),
					router_init.NewRoute("/playlist/{id}/make_public", playlistHandler.MakePublic, http.MethodPut),
					router_init.NewRoute("/playlist/{id}/update_preview", playlistHandler.UpdatePreview, http.MethodPost),
					router_init.NewRoute("/playlist/{id}/update_name", playlistHandler.UpdateName, http.MethodPost),
				},
				Prefix: "",
				Middlewares: []mux.MiddlewareFunc{
					modifyPlaylistMiddleware.ModifyPlaylistAccess,
				},
				SubRouterConfigs: nil,
			}},
	}
	router := router_init.New(routerConfig, logger, waveHandler)
	logger.Errorln("test error for deploying check")

	logger.Fatalln(http.ListenAndServe(ServerPort, router))
}
