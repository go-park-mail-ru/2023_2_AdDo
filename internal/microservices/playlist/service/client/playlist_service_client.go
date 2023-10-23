package grpc_playlist

import (
	"context"
	"github.com/sirupsen/logrus"
	grpc_album "main/internal/microservices/album/service/client"
	playlist_proto "main/internal/microservices/playlist/proto"
	grpc_playlist_server "main/internal/microservices/playlist/service/server"
	"main/internal/pkg/playlist"
)

type Client struct {
	playlistManager playlist_proto.PlaylistServiceClient
	logger          *logrus.Logger
}

func NewClient(pm playlist_proto.PlaylistServiceClient, logger *logrus.Logger) Client {
	return Client{playlistManager: pm, logger: logger}
}

func DeserializePlaylistResponse(in *playlist_proto.Response) playlist.Response {
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

	result, err := c.playlistManager.GetUserPlaylists(context.Background(), &playlist_proto.UserId{Id: userId})
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

func (c *Client) UpdatePreview(playlistId uint64, image string) error {
	c.logger.Infoln("Playlist client  entered")

	_, err := c.playlistManager.UpdatePreview(context.Background(), &playlist_proto.PlaylistIdToImageUrl{Id: playlistId, Image: image})
	if err != nil {
		return err
	}

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
