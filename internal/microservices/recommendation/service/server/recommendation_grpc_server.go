package grpc_recommendation_server

import (
	"context"
	"github.com/sirupsen/logrus"
	candidate "main/internal/microservices/candidate/proto"
	"main/internal/microservices/recommendation/proto"
)

type RecommendationManager struct {
	proto.UnimplementedRecommendationServiceServer
	logger *logrus.Logger
}

func NewRecommendationManager(l *logrus.Logger) RecommendationManager {
	return RecommendationManager{logger: l}
}

func (rm *RecommendationManager) ClassifyCandidates(ctx context.Context, id *proto.CandidateToUserId) (*candidate.Candidates, error) {
	// add rnn model and range these tracks
	return id.GetTracks(), nil
}
