package candidate_service_client

import (
	"context"
	"github.com/sirupsen/logrus"
	proto "main/internal/microservices/candidate/proto"
	session_proto "main/internal/microservices/session/proto"
	grpc_track "main/internal/microservices/track/service/client"
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

func (c *Client) GetCandidateForUser(userId string) ([]track.Response, error) {
	c.logger.Infoln("Get Candidate for user CandidateClient entered")

	candidates, err := c.candidateManager.GetCandidatesForUser(context.Background(), &session_proto.UserId{UserId: userId})
	if err != nil {
		c.logger.Errorln("Getting candidate finished with error", err)
		return nil, err
	}

	return grpc_track.DeserializeTracks(candidates.GetCandidates()), nil
}
