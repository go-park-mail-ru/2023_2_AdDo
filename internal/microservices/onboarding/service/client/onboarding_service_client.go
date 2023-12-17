package onboarding_service_client

import (
	"context"
	google_proto "github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	grpc_artist "main/internal/microservices/artist/service/client"
	grpc_artist_server "main/internal/microservices/artist/service/server"
	onboarding_proto "main/internal/microservices/onboarding/proto"
	"main/internal/pkg/artist"
	"main/internal/pkg/onboarding"
)

type Client struct {
	onboardingManager onboarding_proto.OnboardingServiceClient
	logger            *logrus.Logger
}

func NewClient(pm onboarding_proto.OnboardingServiceClient, logger *logrus.Logger) Client {
	return Client{onboardingManager: pm, logger: logger}
}

func DeserializeGenreBase(base *onboarding_proto.GenreBase) onboarding.GenreBase {
	return onboarding.GenreBase{
		Id:      base.GetId(),
		Name:    base.GetName(),
		Preview: base.GetPreview(),
	}
}

func SerializeGenreBase(in onboarding.GenreBase) *onboarding_proto.GenreBase {
	return &onboarding_proto.GenreBase{
		Id:      in.Id,
		Name:    in.Name,
		Preview: in.Preview,
	}
}

func DeserializeGenresBase(base *onboarding_proto.GenresBase) []onboarding.GenreBase {
	result := make([]onboarding.GenreBase, 0)
	for _, b := range base.GetGenres() {
		result = append(result, DeserializeGenreBase(b))
	}

	return result
}

func SerializeGenresBase(in []onboarding.GenreBase) *onboarding_proto.GenresBase {
	result := make([]*onboarding_proto.GenreBase, 0)
	for _, b := range in {
		result = append(result, SerializeGenreBase(b))
	}
	return &onboarding_proto.GenresBase{Genres: result}
}

func (c *Client) GetGenres() ([]onboarding.GenreBase, error) {
	c.logger.Infoln("Onboard Client Get Genres entered")

	result, err := c.onboardingManager.GetGenres(context.Background(), &google_proto.Empty{})
	if err != nil {
		return nil, err
	}

	return DeserializeGenresBase(result), nil
}

func (c *Client) GetArtists() ([]artist.Base, error) {
	c.logger.Infoln("Onboard Client Get Artists entered")

	result, err := c.onboardingManager.GetArtists(context.Background(), &google_proto.Empty{})
	if err != nil {
		return nil, err
	}

	return grpc_artist.DeserializeArtistsBase(result), nil
}

func (c *Client) SetUserGenres(o onboarding.GenreOnboarding) error {
	c.logger.Infoln("Onboard Client Set Genres entered")

	_, err := c.onboardingManager.SaveGenres(context.Background(), &onboarding_proto.GenreOnboarding{
		Genres: SerializeGenresBase(o.Genres),
		UserId: o.UserId,
	})
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) SetUserArtists(o onboarding.ArtistOnboarding) error {
	c.logger.Infoln("Onboard Client Set Artists entered")

	_, err := c.onboardingManager.SaveArtists(context.Background(), &onboarding_proto.ArtistOnboarding{
		Artists: grpc_artist_server.SerializeArtistsBase(o.Artists),
		UserId:  o.UserId,
	})
	if err != nil {
		return err
	}

	return nil
}
