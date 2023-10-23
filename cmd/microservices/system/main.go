package main

import (
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	logger_init "main/init/logger"
	router_init "main/init/router"
	proto2 "main/internal/microservices/album/proto"
	grpc_album "main/internal/microservices/album/service/client"
	artist "main/internal/microservices/artist/proto"
	grpc_artist "main/internal/microservices/artist/service/client"
	proto3 "main/internal/microservices/playlist/proto"
	grpc_playlist "main/internal/microservices/playlist/service/client"
	session2 "main/internal/microservices/session/proto"
	grpc_session "main/internal/microservices/session/service/client"
	"main/internal/microservices/track/proto"
	grpc_track "main/internal/microservices/track/service/client"
	user_client "main/internal/microservices/user/proto"
	grpc_user "main/internal/microservices/user/service/client"
	album_delivery "main/internal/pkg/album/delivery/http"
	artist_delivery "main/internal/pkg/artist/delivery/http"
	track_delivery "main/internal/pkg/track/delivery/http"
	user_delivery "main/internal/pkg/user/delivery/http"
	"net/http"
)

const EnvPostgresQueryName = "DATABASE_URL"
const ServerPort = ":8080"

func main() {
	logger := logger_init.LogRusInit()

	userConnection, err := grpc.Dial("127.0.0.1:8081", grpc.EmptyDialOption{})
	if err != nil {
		logger.Fatalln("error connecting to user micro ", err)
	}

	sessionConnection, err := grpc.Dial("127.0.0.1:8082", grpc.EmptyDialOption{})
	if err != nil {
		logger.Fatalln("error connecting to user micro ", err)
	}

	trackConnection, err := grpc.Dial("127.0.0.1:8083", grpc.EmptyDialOption{})
	if err != nil {
		logger.Fatalln("error connecting to user micro ", err)
	}

	albumConnection, err := grpc.Dial("127.0.0.1:8084", grpc.EmptyDialOption{})
	if err != nil {
		logger.Fatalln("error connecting to user micro ", err)
	}

	playlistConnection, err := grpc.Dial("127.0.0.1:8085", grpc.EmptyDialOption{})
	if err != nil {
		logger.Fatalln("error connecting to user micro ", err)
	}

	artistConnection, err := grpc.Dial("127.0.0.1:8086", grpc.EmptyDialOption{})
	if err != nil {
		logger.Fatalln("error connecting to user micro ", err)
	}

	userAgent := grpc_user.NewClient(user_client.NewUserServiceClient(userConnection), logger)
	sessionAgent := grpc_session.NewClient(session2.NewSessionServiceClient(sessionConnection), logger)
	trackAgent := grpc_track.NewClient(proto.NewTrackServiceClient(trackConnection), logger)
	albumAgent := grpc_album.NewClient(proto2.NewAlbumServiceClient(albumConnection), logger)
	_ = grpc_playlist.NewClient(proto3.NewPlaylistServiceClient(playlistConnection), logger)
	artistAgent := grpc_artist.NewClient(artist.NewArtistServiceClient(artistConnection), logger)

	logger.Infoln("Clients to micros initialized")

	albumHandler := album_delivery.NewHandler(&trackAgent, &albumAgent, &sessionAgent, logger)
	artistHandler := artist_delivery.NewHandler(&artistAgent, logger)
	userHandler := user_delivery.NewHandler(&userAgent, logger)
	trackHandler := track_delivery.NewHandler(&trackAgent, &sessionAgent, logger)
	logger.Infoln("Deliveries initialized")

	router := router_init.New(userHandler, trackHandler, artistHandler, albumHandler, logger)

	logger.Fatalln(http.ListenAndServe(ServerPort, router))
}
