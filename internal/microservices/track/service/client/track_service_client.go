package grpc_track

import (
	"context"
	"github.com/sirupsen/logrus"
	pb "main/internal/microservices/track/proto"
)

type Client struct {
	trackManager pb.TrackServiceClient
	logger       *logrus.Logger
}

func NewClient(tm pb.TrackServiceClient, logger *logrus.Logger) Client {
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

func (c *Client) IsLike(userId string, trackId uint64) (bool, error) {
	c.logger.Infoln("Track Client Like entered")

	isLiked, err := c.trackManager.IsLike(context.Background(), &pb.TrackToUserId{TrackId: trackId, UserId: userId})
	if err != nil {
		return false, err
	}

	return isLiked.GetIsLiked(), nil
}

func (c *Client) Unlike(userId string, trackId uint64) error {
	c.logger.Infoln("Track Client Unlike entered")

	_, err := c.trackManager.Unlike(context.Background(), &pb.TrackToUserId{TrackId: trackId, UserId: userId})
	if err != nil {
		return err
	}

	return nil
}
