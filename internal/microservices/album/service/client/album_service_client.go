package grpc_album

import (
	"context"
	google_proto "github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"main/internal/microservices/album/proto"
	grpc_track "main/internal/microservices/track/service/client"
	"main/internal/pkg/album"
)

type Client struct {
	albumManager proto.AlbumServiceClient
	logger       *logrus.Logger
}

func DeserializeAlbum(in *proto.AlbumResponse) album.Response {
	return album.Response{
		Id:         in.GetId(),
		Name:       in.GetName(),
		Preview:    in.GetPreview(),
		ArtistId:   in.GetArtistId(),
		ArtistName: in.GetArtistName(),
		Tracks:     grpc_track.DeserializeTracks(in.GetTracks()),
	}
}

func DeserializeAlbumsResponse(in *proto.AlbumsResponse) []album.Response {
	result := make([]album.Response, 0)
	for _, a := range in.GetAlbums() {
		result = append(result, DeserializeAlbum(a))
	}
	return result
}

func NewClient(albumManager proto.AlbumServiceClient, logger *logrus.Logger) Client {
	return Client{albumManager: albumManager, logger: logger}
}

func (c *Client) GetRandom() ([]album.Response, error) {
	c.logger.Infoln("Client to Album Micros GetRandom entered")

	result, err := c.albumManager.GetRandom(context.Background(), &google_proto.Empty{})
	if err != nil {
		return nil, err
	}

	return DeserializeAlbumsResponse(result), nil
}

func (c *Client) GetMostLiked() ([]album.Response, error) {
	c.logger.Infoln("Client to Album Micros GetMostLiked entered")

	result, err := c.albumManager.GetMostLiked(context.Background(), &google_proto.Empty{})
	if err != nil {
		return nil, err
	}

	return DeserializeAlbumsResponse(result), nil
}

func (c *Client) GetPopular() ([]album.Response, error) {
	c.logger.Infoln("Client to Album Micros GetPopular entered")

	result, err := c.albumManager.GetPopular(context.Background(), &google_proto.Empty{})
	if err != nil {
		return nil, err
	}

	return DeserializeAlbumsResponse(result), nil
}

func (c *Client) GetNew() ([]album.Response, error) {
	c.logger.Infoln("Client to Album Micros GetNew entered")

	result, err := c.albumManager.GetNew(context.Background(), &google_proto.Empty{})
	if err != nil {
		return nil, err
	}

	return DeserializeAlbumsResponse(result), nil
}

func (c *Client) GetAlbum(albumId uint64) (album.Response, error) {
	c.logger.Infoln("Client to Album Micros GetAlbum entered")

	result, err := c.albumManager.GetAlbum(context.Background(), &proto.AlbumId{AlbumId: albumId})
	if err != nil {
		return album.Response{}, err
	}

	return DeserializeAlbum(result), nil
}

func (c *Client) Like(userId string, albumId uint64) error {
	c.logger.Infoln("Client to Album Micros Like entered")

	if _, err := c.albumManager.Like(context.Background(), &proto.AlbumToUserId{UserId: userId, AlbumId: albumId}); err != nil {
		return err
	}

	return nil
}

func (c *Client) IsLike(userId string, albumId uint64) (bool, error) {
	c.logger.Infoln("Client to Album Micros Like entered")

	isLiked, err := c.albumManager.IsLike(context.Background(), &proto.AlbumToUserId{UserId: userId, AlbumId: albumId})
	if err != nil {
		return false, err
	}

	return isLiked.GetIsLiked(), nil
}

func (c *Client) Unlike(userId string, albumId uint64) error {
	c.logger.Infoln("Client to Album Micros Like entered")

	if _, err := c.albumManager.Unlike(context.Background(), &proto.AlbumToUserId{UserId: userId, AlbumId: albumId}); err != nil {
		return err
	}

	return nil
}
