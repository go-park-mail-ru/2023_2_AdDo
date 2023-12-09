package grpc_daily_playlist

import (
	"context"
	"github.com/sirupsen/logrus"
	daily_playlist_proto "main/internal/microservices/daily-playlist/proto"
	session_proto "main/internal/microservices/session/proto"
	grpc_track "main/internal/microservices/track/service/client"
	daily_playlist "main/internal/pkg/daily-playlist"
)

type Client struct {
	dailyManager daily_playlist_proto.DailyPlaylistServiceClient
	logger       *logrus.Logger
}

func NewClient(pm daily_playlist_proto.DailyPlaylistServiceClient, logger *logrus.Logger) Client {
	return Client{dailyManager: pm, logger: logger}
}

func (c *Client) GetUserDailyPlaylist(userId string) (daily_playlist.Response, error) {
	c.logger.Infoln("entered daily playlist client")

	result, err := c.dailyManager.GetUserDailyPlaylist(context.Background(), &session_proto.UserId{UserId: userId})
	if err != nil {
		return daily_playlist.Response{}, err
	}

	return daily_playlist.Response{OwnerId: result.GetOwnerId(),
		Id:     result.GetId(),
		Tracks: grpc_track.DeserializeTracks(result.GetTracks())}, nil
}
