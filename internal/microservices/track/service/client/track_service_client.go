package grpc_track

import (
	"context"
	"github.com/sirupsen/logrus"
	session_proto "main/internal/microservices/session/proto"
	track_proto "main/internal/microservices/track/proto"
	grpc_track_server "main/internal/microservices/track/service/server"
	"main/internal/pkg/track"
)

type Client struct {
	trackManager track_proto.TrackServiceClient
	logger       *logrus.Logger
}

func NewClient(tm track_proto.TrackServiceClient, logger *logrus.Logger) Client {
	return Client{trackManager: tm, logger: logger}
}

const MinTimeToListen = 40

func (c *Client) Listen(userId string, trackId uint64, dur uint32) error {
	c.logger.Infoln("Track Client Listen entered")

	if dur < MinTimeToListen {
		if _, err := c.trackManager.Skip(context.Background(), &track_proto.TrackToUserDur{TrackToUser: &track_proto.TrackToUserId{UserId: userId, TrackId: trackId}, Duration: dur}); err != nil {
			return err
		}
	} else {
		if _, err := c.trackManager.Listen(context.Background(), &track_proto.TrackToUserDur{TrackToUser: &track_proto.TrackToUserId{UserId: userId, TrackId: trackId}, Duration: dur}); err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) Like(userId string, trackId uint64) error {
	c.logger.Infoln("Track Client Like entered")

	if _, err := c.trackManager.Like(context.Background(), &track_proto.TrackToUserId{TrackId: trackId, UserId: userId}); err != nil {
		return err
	}

	return nil
}

func (c *Client) IsLike(userId string, trackId uint64) (bool, error) {
	c.logger.Infoln("Track Client IsLike entered")

	isLiked, err := c.trackManager.IsLike(context.Background(), &track_proto.TrackToUserId{TrackId: trackId, UserId: userId})
	if err != nil {
		return false, err
	}

	return isLiked.GetIsLiked(), nil
}

func (c *Client) Unlike(userId string, trackId uint64) error {
	c.logger.Infoln("Track Client Unlike entered")

	if _, err := c.trackManager.Unlike(context.Background(), &track_proto.TrackToUserId{TrackId: trackId, UserId: userId}); err != nil {
		return err
	}

	return nil
}

func (c *Client) GetUserLikedTracks(userId string) ([]track.Response, error) {
	c.logger.Infoln("Track Client GetUserLikedTracks entered")

	result, err := c.trackManager.GetUserLikedTracks(context.Background(), &session_proto.UserId{UserId: userId})
	if err != nil {
		c.logger.Errorln(err)
		return nil, err
	}

	return grpc_track_server.DeserializeTracks(result), nil
}

func (c *Client) LabelIsLikedTracks(userId string, tracks []track.Response) ([]track.Response, error) {
	c.logger.Infoln("Track Client Label Tracks entered")

	result, err := c.trackManager.LabelIsLikedForUser(context.Background(), &track_proto.UserToTracksForLabeling{
		Tracks: grpc_track_server.SerializeTracks(tracks),
		UserId: userId,
	})
	if err != nil {
		c.logger.Errorln(err)
		return nil, err
	}

	return grpc_track_server.DeserializeTracks(result), nil
}
