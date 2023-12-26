package grpc_recommendation_server

import (
	"context"
	"github.com/sirupsen/logrus"
	candidate "main/internal/microservices/candidate/proto"
	"main/internal/microservices/recommendation/proto"
	candidate2 "main/internal/pkg/candidate"
)

type RecommendationManager struct {
	proto.UnimplementedRecommendationServiceServer
	logger           *logrus.Logger
	candidateUseCase candidate2.UseCase
}

func NewRecommendationManager(l *logrus.Logger) RecommendationManager {
	return RecommendationManager{
		logger: l,
	}
}

func (rm *RecommendationManager) ClassifyCandidates(ctx context.Context, id *proto.CandidateToUserId) (*candidate.Candidates, error) {
	// tracks := grpc_track_server.DeserializeTracks(id.GetTracks().GetTracks())
	//vecs := rm.candidateUseCase.GetVecsForTracks(tracks)
	//
	//session, err := tensorflow.NewSession(graph, nil)
	//if err != nil {
	//	fmt.Println("Ошибка при создании сессии TensorFlow:", err)
	//	return
	//}
	return id.GetTracks(), nil

}
