package recommendation_service_client

import (
	"context"
	"github.com/sirupsen/logrus"
	candidate "main/internal/microservices/candidate/proto"
	proto "main/internal/microservices/recommendation/proto"
	session_proto "main/internal/microservices/session/proto"
	grpc_track_server "main/internal/microservices/track/service/server"
	"main/internal/pkg/track"
)

type Client struct {
	recommendationManager proto.RecommendationServiceClient
	logger                *logrus.Logger
}

func NewClient(am proto.RecommendationServiceClient, logger *logrus.Logger) Client {
	return Client{
		recommendationManager: am,
		logger:                logger,
	}
}

func (c *Client) ClassifyCandidates(userId string, candidates []track.Response) ([]track.Response, error) {
	c.logger.Infoln("Recommendation Client ClassifyCandidates entered")

	candidatesClassifiedAndRanged, err := c.recommendationManager.ClassifyCandidates(context.Background(), &proto.CandidateToUserId{
		Tracks: &candidate.Candidates{Tracks: grpc_track_server.SerializeTracks(candidates)},
		UserId: &session_proto.UserId{UserId: userId},
	})
	if err != nil {
		c.logger.Infoln("Error while NN classification", err)
		return nil, err
	}

	return grpc_track_server.DeserializeTracks(candidatesClassifiedAndRanged.GetTracks()), nil
}
