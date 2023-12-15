package candidate_service_client

import (
	"context"
	"github.com/sirupsen/logrus"
	proto "main/internal/microservices/candidate/proto"
	session_proto "main/internal/microservices/session/proto"
	grpc_track_server "main/internal/microservices/track/service/server"
	"main/internal/pkg/track"
)

type Client struct {
	candidateManager proto.CandidateServiceClient
	logger           *logrus.Logger
}

func NewClient(am proto.CandidateServiceClient, logger *logrus.Logger) Client {
	return Client{
		candidateManager: am,
		logger:           logger,
	}
}

func (c *Client) GetCandidateForDaily(userId string) ([]track.Response, error) {
	c.logger.Infoln("Get Candidate for Daily Playlist entered")

	candidates, err := c.candidateManager.GetCandidatesForDailyPlaylist(context.Background(), &session_proto.UserId{UserId: userId})
	if err != nil {
		c.logger.Errorln("Getting candidate finished with error", err)
		return nil, err
	}

	return grpc_track_server.DeserializeTracks(candidates.GetTracks()), nil
}

func (c *Client) GetCandidateForWave(userId string) ([]track.Response, error) {
	c.logger.Infoln("Get Candidate for Wave entered")

	candidates, err := c.candidateManager.GetCandidatesForWave(context.Background(), &session_proto.UserId{UserId: userId})
	if err != nil {
		c.logger.Errorln("Getting candidate finished with error", err)
		return nil, err
	}

	return grpc_track_server.DeserializeTracks(candidates.GetTracks()), nil
}
