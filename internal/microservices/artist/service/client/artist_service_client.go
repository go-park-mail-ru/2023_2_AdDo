package grpc_artist

import (
	"context"
	"github.com/sirupsen/logrus"
	album_proto "main/internal/microservices/album/proto"
	proto "main/internal/microservices/artist/proto"
	grpc_playlist "main/internal/microservices/playlist/service/client"
	session_proto "main/internal/microservices/session/proto"
	grpc_track_server "main/internal/microservices/track/service/server"
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

func DeserializeAlbumBase(in *album_proto.AlbumBase) album.Base {
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
		Tracks: grpc_track_server.DeserializeTracks(in.GetTracks()),
	}
}

func DeserializeArtistBase(in *proto.ArtistBase) artist.Base {
	return artist.Base{
		Id:     in.GetId(),
		Name:   in.GetName(),
		Avatar: in.GetAvatar(),
	}
}

func DeserializeArtistsBase(in *proto.ArtistsBase) []artist.Base {
	result := make([]artist.Base, 0)
	for _, base := range in.GetArtists() {
		result = append(result, DeserializeArtistBase(base))
	}
	return result
}

func DeserializeSearchResponse(in *proto.SearchResponse) artist.SearchResponse {
	return artist.SearchResponse{
		Artists:   DeserializeArtistsBase(in.GetArtists()),
		Albums:    DeserializeAlbumsBase(in.GetAlbums()),
		Playlists: grpc_playlist.DeserializePlaylistsBase(in.GetPlaylists()),
		Tracks:    grpc_track_server.DeserializeTracks(in.GetTracks()),
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

func (c *Client) Like(userId string, artistId uint64) error {
	c.logger.Infoln("Client for artist micros")

	if _, err := c.artistManager.Like(context.Background(), &proto.ArtistToUserId{UserId: userId, ArtistId: artistId}); err != nil {
		return err
	}

	return nil
}

func (c *Client) IsLike(userId string, artistId uint64) (bool, error) {
	c.logger.Infoln("Client for artist micros")

	isLiked, err := c.artistManager.IsLike(context.Background(), &proto.ArtistToUserId{UserId: userId, ArtistId: artistId})
	if err != nil {
		return false, err
	}

	return isLiked.GetIsLiked(), nil
}

func (c *Client) Unlike(userId string, artistId uint64) error {
	c.logger.Infoln("Client for artist micros")

	if _, err := c.artistManager.Unlike(context.Background(), &proto.ArtistToUserId{UserId: userId, ArtistId: artistId}); err != nil {
		return err
	}

	return nil
}

func (c *Client) FullSearch(query string) (artist.SearchResponse, error) {
	c.logger.Infoln("Client for artist micros Full Search")

	result, err := c.artistManager.FullSearch(context.Background(), &proto.Query{Query: query})
	if err != nil {
		return artist.SearchResponse{}, err
	}
	return DeserializeSearchResponse(result), nil
}

func (c *Client) GetUserArtists(userId string) (artist.Artists, error) {
	c.logger.Infoln("Client for artist micros GetUserArtists")

	result, err := c.artistManager.CollectionArtist(context.Background(), &session_proto.UserId{UserId: userId})
	if err != nil {
		return artist.Artists{}, err
	}

	return artist.Artists{Artists: DeserializeArtistsBase(result)}, nil
}
