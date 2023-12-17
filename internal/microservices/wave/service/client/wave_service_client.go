package grpc_wave_client

import (
	"context"
	"github.com/sirupsen/logrus"
	grpc_track_server "main/internal/microservices/track/service/server"
	"main/internal/microservices/wave/proto"
	"main/internal/pkg/track"
)

type Client struct {
	waveManager proto.WaveServiceClient
	logger      *logrus.Logger
}

func NewClient(wm proto.WaveServiceClient, logger *logrus.Logger) Client {
	return Client{waveManager: wm, logger: logger}
}

func (c *Client) GetMyWaveMusic(userId string, count uint32) ([]track.Response, error) {
	result, err := c.waveManager.GetMyWaveMusic(context.Background(), &proto.UserIdCount{
		UserId:     userId,
		TrackCount: count,
	})
	if err != nil {
		c.logger.Errorln("error while getting wave tracks", err)
		return make([]track.Response, 0), err
	}

	return grpc_track_server.DeserializeTracks(result), nil
}
