package grpc_playlist_server

import (
	"context"
	google_proto "github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	grpc_album_server "main/internal/microservices/album/service/server"
	playlist_proto "main/internal/microservices/playlist/proto"
	session_proto "main/internal/microservices/session/proto"
	"main/internal/pkg/playlist"
	"main/internal/pkg/track"
)

type PlaylistManager struct {
	repoPlaylist playlist.Repository
	repoTracks   track.Repository
	logger       *logrus.Logger
	playlist_proto.UnimplementedPlaylistServiceServer
}

func NewPlaylistManager(rp playlist.Repository, rt track.Repository, logger *logrus.Logger) PlaylistManager {
	return PlaylistManager{
		repoPlaylist: rp,
		repoTracks:   rt,
		logger:       logger,
	}
}

func DeserializePlaylist(in *playlist_proto.PlaylistBase) playlist.Base {
	return playlist.Base{
		Id:       in.GetId(),
		Name:     in.GetName(),
		AuthorId: in.GetCreatorId(),
		Preview:  in.GetPreview(),
	}
}

func SerializePlaylistResponse(in playlist.Response) *playlist_proto.PlaylistResponse {
	return &playlist_proto.PlaylistResponse{
		Id:        in.Id,
		Name:      in.Name,
		Preview:   in.Preview,
		CreatorId: in.AuthorId,
		Tracks:    grpc_album_server.SerializeTracks(in.Tracks),
	}
}

func SerializePlaylistBase(in playlist.Base) *playlist_proto.PlaylistBase {
	return &playlist_proto.PlaylistBase{
		Id:        in.Id,
		Name:      in.Name,
		CreatorId: in.AuthorId,
		Preview:   in.Preview,
	}
}

func SerializePlaylistsBase(in []playlist.Base) *playlist_proto.PlaylistsBase {
	var result *playlist_proto.PlaylistsBase

	for _, base := range in {
		result.Playlists = append(result.Playlists, SerializePlaylistBase(base))
	}
	return result
}

func (pm *PlaylistManager) Create(ctx context.Context, in *playlist_proto.PlaylistBase) (*google_proto.Empty, error) {
	pm.logger.Infoln("Playlist Service Create Method entered")

	err := pm.repoPlaylist.Create(ctx, DeserializePlaylist(in))
	if err != nil {
		return nil, err
	}
	pm.logger.Infoln("Playlist Created")

	return &google_proto.Empty{}, nil
}

func (pm *PlaylistManager) Get(ctx context.Context, in *playlist_proto.PlaylistId) (*playlist_proto.PlaylistResponse, error) {
	pm.logger.Infoln("Playlist Service Get Method entered")

	result, err := pm.repoPlaylist.Get(ctx, in.GetId())
	if err != nil {
		return nil, err
	}

	tracks, err := pm.repoTracks.GetByPlaylist(in.GetId())
	if err != nil {
		return nil, err
	}

	return SerializePlaylistResponse(playlist.Response{
		Id:       result.Id,
		Name:     result.Name,
		AuthorId: result.AuthorId,
		Preview:  "",
		Tracks:   tracks,
	}), nil
}

func (pm *PlaylistManager) GetUserPlaylists(ctx context.Context, in *session_proto.UserId) (*playlist_proto.PlaylistsBase, error) {
	pm.logger.Infoln("Playlist Service GetUserPlaylists Method entered")

	result, err := pm.repoPlaylist.GetByCreatorId(ctx, in.GetUserId())
	if err != nil {
		return nil, err
	}

	return SerializePlaylistsBase(result), nil
}

func (pm *PlaylistManager) AddTrack(ctx context.Context, in *playlist_proto.PlaylistToTrackId) (*google_proto.Empty, error) {
	pm.logger.Infoln("Playlist Service AddTrack Method entered")

	err := pm.repoPlaylist.AddTrack(ctx, in.GetPlaylistId(), in.GetTrackId())
	if err != nil {
		return nil, err
	}
	pm.logger.Infoln("Track Added")

	return &google_proto.Empty{}, nil
}

func (pm *PlaylistManager) UpdatePreview(ctx context.Context, in *playlist_proto.PlaylistIdToImageUrl) (*google_proto.Empty, error) {
	pm.logger.Infoln("Playlist Service UpdatePreview Method entered")

	err := pm.repoPlaylist.UpdateImage(ctx, in.GetId(), in.GetImage())
	if err != nil {
		return nil, err
	}
	pm.logger.Infoln("Photo Updated")

	return &google_proto.Empty{}, nil
}

func (pm *PlaylistManager) DeleteById(ctx context.Context, in *playlist_proto.PlaylistId) (*google_proto.Empty, error) {
	pm.logger.Infoln("Playlist Service DeleteById Method entered")

	err := pm.repoPlaylist.Delete(ctx, in.GetId())
	if err != nil {
		return nil, err
	}
	pm.logger.Infoln("Playlist Deleted")

	return &google_proto.Empty{}, nil
}
