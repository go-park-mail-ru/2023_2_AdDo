package daily_playlist_service_server

import (
	"context"
	"github.com/sirupsen/logrus"
	daily_playlist_grpc "main/internal/microservices/daily-playlist/proto"
	session_proto "main/internal/microservices/session/proto"
	grpc_track_server "main/internal/microservices/track/service/server"
	daily_playlist "main/internal/pkg/daily-playlist"
	"main/internal/pkg/track"
)

type DailyPlaylistManager struct {
	logger            *logrus.Logger
	dailyPlaylistRepo daily_playlist.Repository
	repoTrack         track.Repository
	daily_playlist_grpc.UnimplementedDailyPlaylistServiceServer
}

func NewDailyManager(rt track.Repository, dpr daily_playlist.Repository, logger *logrus.Logger) DailyPlaylistManager {
	return DailyPlaylistManager{
		logger:            logger,
		dailyPlaylistRepo: dpr,
		repoTrack:         rt,
	}
}

func (dpm *DailyPlaylistManager) GetUserDailyPlaylist(_ context.Context, id *session_proto.UserId) (*daily_playlist_grpc.DailyPlaylist, error) {
	dpm.logger.Infoln("DailyPlaylistManager entered GetPlaylist")

	tracks, err := dpm.repoTrack.GetByUserForDailyPlaylist(id.GetUserId())
	if err != nil {
		dpm.logger.Errorln("Error GetPlaylist from repo", err)
		return nil, err
	}

	return &daily_playlist_grpc.DailyPlaylist{
		OwnerId: id.GetUserId(),
		Tracks:  grpc_track_server.SerializeTracks(tracks),
	}, nil
}
