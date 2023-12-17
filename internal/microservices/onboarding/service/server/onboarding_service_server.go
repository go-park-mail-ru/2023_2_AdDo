package grpc_onboarding_server

import (
	"context"
	google_proto "github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	artist_proto "main/internal/microservices/artist/proto"
	grpc_artist "main/internal/microservices/artist/service/client"
	grpc_artist_server "main/internal/microservices/artist/service/server"
	onboarding_proto "main/internal/microservices/onboarding/proto"
	onboarding_service_client "main/internal/microservices/onboarding/service/client"
	"main/internal/pkg/activity"
	"main/internal/pkg/onboarding"
)

type OnboardingManager struct {
	onboarding_proto.UnimplementedOnboardingServiceServer
	queue       activity.ProducerRepository
	logger      *logrus.Logger
	repoOnboard onboarding.Repository
}

func NewOnboardingManager(q activity.ProducerRepository, ro onboarding.Repository, logger *logrus.Logger) OnboardingManager {
	return OnboardingManager{
		logger:      logger,
		queue:       q,
		repoOnboard: ro,
	}
}

func (om *OnboardingManager) GetArtists(ctx context.Context, _ *google_proto.Empty) (*artist_proto.ArtistsBase, error) {
	om.logger.Infoln("Onboard micros Get Artists entered")

	result, err := om.repoOnboard.GetArtists()
	if err != nil {
		return nil, err
	}

	return grpc_artist_server.SerializeArtistsBase(result), nil
}

func (om *OnboardingManager) GetGenres(ctx context.Context, _ *google_proto.Empty) (*onboarding_proto.GenresBase, error) {
	om.logger.Infoln("Onboard micros Get Genres entered")

	result, err := om.repoOnboard.GetGenres()
	if err != nil {
		return nil, err
	}

	return onboarding_service_client.SerializeGenresBase(result), nil
}

func (om *OnboardingManager) SaveArtists(ctx context.Context, in *onboarding_proto.ArtistOnboarding) (*google_proto.Empty, error) {
	om.logger.Infoln("Onboard micros Save Artists entered")

	if err := om.repoOnboard.SetUserArtists(in.GetUserId(), grpc_artist.DeserializeArtistsBase(in.GetArtists())); err != nil {
		return nil, err
	}

	for _, a := range in.GetArtists().GetArtists() {
		if err := om.queue.PushLikeArtist(in.GetUserId(), a.GetId()); err != nil {
			return nil, err
		}
	}

	return &google_proto.Empty{}, nil
}

func (om *OnboardingManager) SaveGenres(ctx context.Context, in *onboarding_proto.GenreOnboarding) (*google_proto.Empty, error) {
	om.logger.Infoln("Onboard micros Save Genres entered")

	if err := om.repoOnboard.SetUserGenres(in.GetUserId(), onboarding_service_client.DeserializeGenresBase(in.GetGenres())); err != nil {
		return nil, err
	}

	for _, a := range in.GetGenres().GetGenres() {
		if err := om.queue.PushLikeGenre(in.GetUserId(), a.GetId()); err != nil {
			return nil, err
		}
	}

	return &google_proto.Empty{}, nil
}
