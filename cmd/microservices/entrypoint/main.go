package main

import (
	"main/init/middleware"
	router_init "main/init/router"
	csrf "main/internal/common/get_csrf"
	log "main/internal/common/logger"
	modify_playlist "main/internal/common/middleware/playlist_middleware/modify"
	read_playlist "main/internal/common/middleware/playlist_middleware/read"
	check_vote "main/internal/common/middleware/survey_middleware"
	proto2 "main/internal/microservices/album/proto"
	grpc_album "main/internal/microservices/album/service/client"
	artist "main/internal/microservices/artist/proto"
	grpc_artist "main/internal/microservices/artist/service/client"
	proto4 "main/internal/microservices/image/proto"
	grpc_image "main/internal/microservices/image/service/client"
	proto3 "main/internal/microservices/playlist/proto"
	grpc_playlist "main/internal/microservices/playlist/service/client"
	session2 "main/internal/microservices/session/proto"
	grpc_session "main/internal/microservices/session/service/client"
	proto5 "main/internal/microservices/survey/proto"
	"main/internal/microservices/track/proto"
	grpc_track "main/internal/microservices/track/service/client"
	user_client "main/internal/microservices/user/proto"
	grpc_user "main/internal/microservices/user/service/client"
	album_delivery "main/internal/pkg/album/delivery/http"
	artist_delivery "main/internal/pkg/artist/delivery/http"
	playlist_delivery "main/internal/pkg/playlist/delivery/http"
	survey_delivery "main/internal/pkg/survey/delivery"
	track_delivery "main/internal/pkg/track/delivery/http"
	user_delivery "main/internal/pkg/user/delivery/http"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const EnvPostgresQueryName = "DATABASE_URL"
const ServerPort = ":8080"

var loggerSingleton = log.Singleton{}

func main() {
	logger := loggerSingleton.GetLogger()

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

	surveyConnection, err := grpc.Dial("images:8088", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatalln("error connecting to survey micros ", err)
	}

	imageAgent := grpc_image.NewClient(proto4.NewImageServiceClient(imageConnection), logger)
	userAgent := grpc_user.NewClient(user_client.NewUserServiceClient(userConnection), imageAgent, logger)
	sessionAgent := grpc_session.NewClient(session2.NewSessionServiceClient(sessionConnection), logger)
	trackAgent := grpc_track.NewClient(proto.NewTrackServiceClient(trackConnection), logger)
	albumAgent := grpc_album.NewClient(proto2.NewAlbumServiceClient(albumConnection), logger)
	playlistAgent := grpc_playlist.NewClient(userAgent, proto3.NewPlaylistServiceClient(playlistConnection), imageAgent, logger)
	artistAgent := grpc_artist.NewClient(artist.NewArtistServiceClient(artistConnection), logger)

	surveyAgent := grpc_survey.NewClient(proto5.NewSurveyServiceClient(surveyConnection), logger)
	logger.Infoln("Clients to micros initialized")

	albumHandler := album_delivery.NewHandler(&trackAgent, &albumAgent, &sessionAgent, logger)
	artistHandler := artist_delivery.NewHandler(&sessionAgent, &artistAgent, logger)
	userHandler := user_delivery.NewHandler(&userAgent, &sessionAgent, logger)
	trackHandler := track_delivery.NewHandler(&trackAgent, &sessionAgent, logger)
	playlistHandler := playlist_delivery.NewHandler(&playlistAgent, &sessionAgent, logger)
	surveyHandler := survey_delivery.NewHandler(&surveyAgent, &sessionAgent, logger)
	logger.Infoln("Deliveries initialized")

	modifyPlaylistMiddleware := modify_playlist.NewMiddleware(&playlistAgent, &sessionAgent, logger)
	readPlaylistMiddleware := read_playlist.NewMiddleware(&playlistAgent, logger)
	checkVoteMiddleware := check_vote.NewMiddleware(&surveyAgent, &sessionAgent, logger)
	corsMiddleware := middleware.NewCors()
	csrfMiddleware := middleware.NewCSRF()

	routerConfig := router_init.Config{
		Routes: []router_init.Route{
			router_init.NewRoute("/get_csrf", csrf.GetCSRF, http.MethodGet),
			router_init.NewRoute("/sign_up", userHandler.SignUp, http.MethodPost),
			router_init.NewRoute("/login", userHandler.Login, http.MethodPost),
			router_init.NewRoute("/update_info", userHandler.UpdateUserInfo, http.MethodPut),
			router_init.NewRoute("/upload_avatar", userHandler.UploadAvatar, http.MethodPost),
			router_init.NewRoute("/remove_avatar", userHandler.RemoveAvatar, http.MethodPost),
			router_init.NewRoute("/auth", userHandler.Auth, http.MethodGet),
			router_init.NewRoute("/me", userHandler.Me, http.MethodGet),
			router_init.NewRoute("/logout", userHandler.LogOut, http.MethodPost),

			router_init.NewRoute("/listen", trackHandler.Listen, http.MethodPost),
			router_init.NewRoute("/track/{id}/like", trackHandler.Like, http.MethodPost),
			router_init.NewRoute("/track/{id}/is_like", trackHandler.IsLike, http.MethodGet),
			router_init.NewRoute("/track/{id}/unlike", trackHandler.Unlike, http.MethodDelete),
			router_init.NewRoute("/collection/tracks", trackHandler.GetUserTracks, http.MethodGet),

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
			router_init.NewRoute("/playlist", playlistHandler.Create, http.MethodPost),
			router_init.NewRoute("/survey/{id}", surveyHandler.GetSurvey, http.MethodGet),
			router_init.NewRoute("/survey/{id}/is_submit", surveyHandler.IsSubmit, http.MethodGet),
			router_init.NewRoute("/survey/{id}/get_stat", surveyHandler.GetStat, http.MethodGet),
		},
		Prefix: "/api/v1",
		Middlewares: []mux.MiddlewareFunc{
			csrfMiddleware, corsMiddleware,
		},

		SubRouterConfigs: []router_init.Config{
			router_init.Config{
				Routes: []router_init.Route{
					router_init.NewRoute("/playlist/{id}/like", playlistHandler.Like, http.MethodPost),
					router_init.NewRoute("/playlist/{id}/is_like", playlistHandler.IsLike, http.MethodGet),
					router_init.NewRoute("/playlist/{id}/unlike", playlistHandler.Unlike, http.MethodDelete),
					router_init.NewRoute("/playlist/{id}", playlistHandler.Get, http.MethodGet),
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
				},
				Prefix: "",
				Middlewares: []mux.MiddlewareFunc{
					modifyPlaylistMiddleware.ModifyPlaylistAccess,
				},
				SubRouterConfigs: nil,
			},

			router_init.Config{
				Routes: []router_init.Route{
					router_init.NewRoute("/survey/submit/{id}", surveyHandler.Submit, http.MethodPost),
				},
				Prefix: "",
				Middlewares: []mux.MiddlewareFunc{
					checkVoteMiddleware.CheckUserVote,
				},
				SubRouterConfigs: nil,
			},
		},
	}
	router := router_init.New(routerConfig, logger)

	logger.Fatalln(http.ListenAndServe(ServerPort, router))
}
