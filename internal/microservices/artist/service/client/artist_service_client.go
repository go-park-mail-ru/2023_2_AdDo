package grpc_artist

import (
	"context"
	"github.com/sirupsen/logrus"
	album_proto "main/internal/microservices/album/proto"
	grpc_album "main/internal/microservices/album/service/client"
	proto "main/internal/microservices/artist/proto"
	"main/internal/pkg/album"
	"main/internal/pkg/artist"
)

type Client struct {
	artistManager proto.ArtistServiceClient
	logger        *logrus.Logger
}

func NewClient(am proto.ArtistServiceClient, logger *logrus.Logger) Client {
	return Client{
		artistManager: am,
		logger:        logger,
	}
}

func DeserializeAlbumBase(in *album_proto.Base) album.Base {
	return album.Base{
		Id:      in.GetId(),
		Name:    in.GetName(),
		Preview: in.GetPreview(),
	}
}

func DeserializeAlbumsBase(in *album_proto.AlbumsBase) []album.Base {
	result := make([]album.Base, 0)
	for _, base := range in.GetAlbums() {
		result = append(result, DeserializeAlbumBase(base))
	}

	return result
}

func DeserializeArtist(in *proto.Artist) artist.Response {
	return artist.Response{
		Id:     in.GetId(),
		Name:   in.GetName(),
		Avatar: in.GetAvatar(),
		Albums: DeserializeAlbumsBase(in.GetAlbums()),
		Tracks: grpc_album.DeserializeTracks(in.GetTracks()),
	}
}

func (c *Client) GetArtistInfo(artistId uint64) (artist.Response, error) {
	c.logger.Infoln("Client for artist micros")

	result, err := c.artistManager.GetArtistInfo(context.Background(), &proto.ArtistId{ArtistId: artistId})
	if err != nil {
		return artist.Response{}, err
	}

	return DeserializeArtist(result), nil
}