package grpc_track

import (
	"context"
	"github.com/sirupsen/logrus"
	pb "main/internal/microservices/track/proto"
	grpc_track_server "main/internal/microservices/track/service/server"
)

type Client struct {
	trackManager grpc_track_server.TrackManager
	logger       *logrus.Logger
}

func NewClient(tm grpc_track_server.TrackManager, logger *logrus.Logger) Client {
	return Client{trackManager: tm, logger: logger}
}

func (c *Client) Listen(trackId uint64) error {
	c.logger.Infoln("Track Client Listen entered")

	_, err := c.trackManager.Listen(context.Background(), &pb.TrackId{TrackId: trackId})
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Like(userId string, trackId uint64) error {
	c.logger.Infoln("Track Client Like entered")

	_, err := c.trackManager.Like(context.Background(), &pb.TrackToUserId{TrackId: trackId, UserId: userId})
	if err != nil {
		return err
	}

	return nil
}