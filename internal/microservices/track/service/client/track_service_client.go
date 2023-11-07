package grpc_track

import (
	"context"
	"github.com/sirupsen/logrus"
	session_proto "main/internal/microservices/session/proto"
	track_proto "main/internal/microservices/track/proto"
	"main/internal/pkg/track"
)

type Client struct {
	trackManager track_proto.TrackServiceClient
	logger       *logrus.Logger
}

func NewClient(tm track_proto.TrackServiceClient, logger *logrus.Logger) Client {
	return Client{trackManager: tm, logger: logger}
}

func DeserializeTrack(in *track_proto.Track) track.Response {
	return track.Response{
		Id:         in.GetId(),
		Name:       in.GetName(),
		Preview:    in.GetPreview(),
		Content:    in.GetContent(),
		ArtistId:   in.GetArtistId(),
		ArtistName: in.GetArtistName(),
		Duration:   in.GetDuration(),
		IsLiked:    in.GetIsLiked(),
	}
}

func DeserializeTracks(in *track_proto.TracksResponse) []track.Response {
	result := make([]track.Response, 0)
	for _, t := range in.GetTracks() {
		result = append(result, DeserializeTrack(t))
	}
	return result
}

func (c *Client) Listen(trackId uint64) error {
	c.logger.Infoln("Track Client Listen entered")

	if _, err := c.trackManager.Listen(context.Background(), &track_proto.TrackId{TrackId: trackId}); err != nil {
		return err
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

	return DeserializeTracks(result), nil
}
