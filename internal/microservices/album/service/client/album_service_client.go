package grpc_album

import (
	"context"
	"github.com/sirupsen/logrus"
	"main/internal/microservices/album/proto"
	pb "main/internal/microservices/track/proto"
	"main/internal/pkg/album"
	"main/internal/pkg/track"
)

type Client struct {
	albumManager proto.AlbumServiceClient
	logger       *logrus.Logger
}

func DeserializeTrack(in *pb.Track) track.Response {
	return track.Response{
		Id:         in.GetId(),
		Name:       in.GetName(),
		Preview:    in.GetPreview(),
		Content:    in.GetContent(),
		ArtistName: in.GetArtistName(),
		Duration:   in.GetDuration(),
		IsLiked:    in.GetIsLiked(),
	}
}

func DeserializeTracks(in *pb.TracksResponse) []track.Response {
	result := make([]track.Response, 0)
	for _, t := range in.GetTracks() {
		result = append(result, DeserializeTrack(t))
	}
	return result
}

func DeserializeAlbum(in *proto.Album) album.Response {
	return album.Response{
		Id:         in.GetId(),
		Name:       in.GetName(),
		Preview:    in.GetPreview(),
		ArtistId:   in.GetArtistId(),
		ArtistName: in.GetArtistName(),
		Tracks:     DeserializeTracks(in.GetTracks()),
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

	result, err := c.albumManager.GetRandom(context.Background(), &pb.Status{IsOk: true})
	if err != nil {
		return nil, err
	}

	return DeserializeAlbumsResponse(result), nil
}

func (c *Client) GetMostLiked() ([]album.Response, error) {
	c.logger.Infoln("Client to Album Micros GetMostLiked entered")

	result, err := c.albumManager.GetMostLiked(context.Background(), &pb.Status{IsOk: true})
	if err != nil {
		return nil, err
	}

	return DeserializeAlbumsResponse(result), nil
}

func (c *Client) GetPopular() ([]album.Response, error) {
	c.logger.Infoln("Client to Album Micros GetPopular entered")

	result, err := c.albumManager.GetPopular(context.Background(), &pb.Status{IsOk: true})
	if err != nil {
		return nil, err
	}

	return DeserializeAlbumsResponse(result), nil
}

func (c *Client) GetNew() ([]album.Response, error) {
	c.logger.Infoln("Client to Album Micros GetNew entered")

	result, err := c.albumManager.GetNew(context.Background(), &pb.Status{IsOk: true})
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

	_, err := c.albumManager.Like(context.Background(), &proto.AlbumToUserId{UserId: userId, AlbumId: albumId})
	if err != nil {
		return err
	}

	return nil
}
