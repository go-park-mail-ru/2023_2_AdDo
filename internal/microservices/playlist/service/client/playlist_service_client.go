package grpc_playlist

import (
	"context"
	"github.com/sirupsen/logrus"
	"io"
	grpc_album "main/internal/microservices/album/service/client"
	image_proto "main/internal/microservices/image/proto"
	grpc_image "main/internal/microservices/image/service/client"
	playlist_proto "main/internal/microservices/playlist/proto"
	grpc_playlist_server "main/internal/microservices/playlist/service/server"
	session_proto "main/internal/microservices/session/proto"
	"main/internal/pkg/playlist"
)

type Client struct {
	playlistManager playlist_proto.PlaylistServiceClient
	imageClient     grpc_image.Client
	logger          *logrus.Logger
}

func NewClient(pm playlist_proto.PlaylistServiceClient, client grpc_image.Client, logger *logrus.Logger) Client {
	return Client{playlistManager: pm, logger: logger, imageClient: client}
}

func DeserializePlaylistResponse(in *playlist_proto.PlaylistResponse) playlist.Response {
	return playlist.Response{
		Id:       in.GetId(),
		Name:     in.GetName(),
		AuthorId: in.GetCreatorId(),
		Preview:  in.GetPreview(),
		Tracks:   grpc_album.DeserializeTracks(in.GetTracks()),
	}
}

func DeserializePlaylistsBase(in *playlist_proto.PlaylistsBase) []playlist.Base {
	result := make([]playlist.Base, 0)
	for _, base := range in.GetPlaylists() {
		result = append(result, grpc_playlist_server.DeserializePlaylist(base))
	}
	return result
}

func (c *Client) Create(base playlist.Base) error {
	c.logger.Infoln("Playlist client  entered")

	_, err := c.playlistManager.Create(context.Background(), grpc_playlist_server.SerializePlaylistBase(base))
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Get(playlistId uint64) (playlist.Response, error) {
	c.logger.Infoln("Playlist client  entered")

	result, err := c.playlistManager.Get(context.Background(), &playlist_proto.PlaylistId{Id: playlistId})
	if err != nil {
		return playlist.Response{}, err
	}

	return DeserializePlaylistResponse(result), nil
}

func (c *Client) GetUserPlaylists(userId string) ([]playlist.Base, error) {
	c.logger.Infoln("Playlist client  entered")

	result, err := c.playlistManager.GetUserPlaylists(context.Background(), &session_proto.UserId{UserId: userId})
	if err != nil {
		return nil, err
	}

	return DeserializePlaylistsBase(result), nil
}

func (c *Client) AddTrack(playlistId, trackId uint64) error {
	c.logger.Infoln("Playlist client  entered")

	_, err := c.playlistManager.AddTrack(context.Background(), &playlist_proto.PlaylistToTrackId{PlaylistId: playlistId, TrackId: trackId})
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) RemoveTrack(playlistId, trackId uint64) error {
	c.logger.Infoln("Playlist client RemoveTrack  entered")

	_, err := c.playlistManager.RemoveTrack(context.Background(), &playlist_proto.PlaylistToTrackId{PlaylistId: playlistId, TrackId: trackId})
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) UpdatePreview(playlistId uint64, src io.Reader, size int64) error {
	c.logger.Infoln("Playlist client  entered")

	url, err := c.imageClient.UploadPlaylistImage(src, size)
	if err != nil {
		return err
	}
	c.logger.Infoln("Image Uploaded")

	_, err = c.playlistManager.UpdatePreview(context.Background(), &playlist_proto.PlaylistIdToImageUrl{Id: playlistId, Url: &image_proto.ImageUrl{Url: url}})
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) RemovePreview(playlistId uint64) error {
	c.logger.Infoln("Playlist client  entered")

	url, err := c.playlistManager.RemovePreview(context.Background(), &playlist_proto.PlaylistId{Id: playlistId})
	if err != nil {
		return err
	}
	c.logger.Infoln("Path deleted from db")

	err = c.imageClient.RemoveImage(url.GetUrl())
	if err != nil {
		return err
	}
	c.logger.Infoln("Image deleted from sss")

	return nil
}

func (c *Client) DeleteById(playlistId uint64) error {
	c.logger.Infoln("Playlist client  entered")

	_, err := c.playlistManager.DeleteById(context.Background(), &playlist_proto.PlaylistId{Id: playlistId})
	if err != nil {
		return err
	}

	return nil
}
