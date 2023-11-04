package grpc_playlist

import (
	"context"
	"github.com/sirupsen/logrus"
	"io"
	image_proto "main/internal/microservices/image/proto"
	grpc_image "main/internal/microservices/image/service/client"
	playlist_proto "main/internal/microservices/playlist/proto"
	grpc_playlist_server "main/internal/microservices/playlist/service/server"
	session_proto "main/internal/microservices/session/proto"
	grpc_track "main/internal/microservices/track/service/client"
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
		IsYours:  in.GetIsYours(),
		AuthorId: in.GetCreatorId(),
		Preview:  in.GetPreview(),
		Tracks:   grpc_track.DeserializeTracks(in.GetTracks()),
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

func (c *Client) Get(userId string, playlistId uint64) (playlist.Response, error) {
	c.logger.Infoln("Playlist client Get entered")

	result, err := c.playlistManager.Get(context.Background(), &playlist_proto.PlaylistToUserId{UserId: userId, PlaylistId: playlistId})
	if err != nil {
		return playlist.Response{}, err
	}

	return DeserializePlaylistResponse(result), nil
}

func (c *Client) GetUserPlaylists(userId string) ([]playlist.Base, error) {
	c.logger.Infoln("Playlist client GetUserPlaylists entered")

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

	_ = c.RemovePreview(playlistId)
	c.logger.Infoln("Last preview and image removed")

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

func (c *Client) Like(userId string, playlistId uint64) error {
	c.logger.Infoln("Client to Playlist Micros Like entered")

	_, err := c.playlistManager.Like(context.Background(), &playlist_proto.PlaylistToUserId{UserId: userId, PlaylistId: playlistId})
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) IsLike(userId string, playlistId uint64) (bool, error) {
	c.logger.Infoln("Client to Playlist Micros CheckLike entered")

	isLiked, err := c.playlistManager.IsLike(context.Background(), &playlist_proto.PlaylistToUserId{UserId: userId, PlaylistId: playlistId})
	if err != nil {
		return false, err
	}

	return isLiked.GetIsLiked(), nil
}

func (c *Client) Unlike(userId string, playlistId uint64) error {
	c.logger.Infoln("Client to Playlist Micros Unlike entered")

	_, err := c.playlistManager.Unlike(context.Background(), &playlist_proto.PlaylistToUserId{UserId: userId, PlaylistId: playlistId})
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) HasModifyAccess(userId string, playlistId uint64) (bool, error) {
	c.logger.Infoln("Client to Playlist Micros HasModifyAccess entered")

	result, err := c.playlistManager.HasModifyAccess(context.Background(), &playlist_proto.PlaylistToUserId{UserId: userId, PlaylistId: playlistId})
	if err != nil {
		return false, err
	}

	return result.GetIsAccess(), nil
}

func (c *Client) HasReadAccess(playlistId uint64) (bool, error) {
	c.logger.Infoln("Client to Playlist Micros HasReadAccess entered")

	result, err := c.playlistManager.HasReadAccess(context.Background(), &playlist_proto.PlaylistId{Id: playlistId})
	if err != nil {
		return false, err
	}

	return result.GetIsAccess(), nil
}

func (c *Client) MakePrivate(playlistId uint64) error {
	c.logger.Infoln("Client to Playlist Micros MakePrivate entered")

	_, err := c.playlistManager.MakePrivate(context.Background(), &playlist_proto.PlaylistId{Id: playlistId})
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) MakePublic(playlistId uint64) error {
	c.logger.Infoln("Client to Playlist Micros MakePublic entered")

	_, err := c.playlistManager.MakePublic(context.Background(), &playlist_proto.PlaylistId{Id: playlistId})
	if err != nil {
		return err
	}

	return nil
}
