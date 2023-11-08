package grpc_playlist_server

import (
	"context"
	google_proto "github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	image_proto "main/internal/microservices/image/proto"
	playlist_proto "main/internal/microservices/playlist/proto"
	session_proto "main/internal/microservices/session/proto"
	grpc_track_server "main/internal/microservices/track/service/server"
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
		Id:          in.Id,
		Name:        in.Name,
		Preview:     in.Preview,
		CreatorId:   in.AuthorId,
		CreatorName: in.AuthorName,
		IsYours:     in.IsYours,
		Tracks:      grpc_track_server.SerializeTracks(in.Tracks),
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
	playlists := make([]*playlist_proto.PlaylistBase, 0)
	for _, base := range in {
		playlists = append(playlists, SerializePlaylistBase(base))
	}
	return &playlist_proto.PlaylistsBase{Playlists: playlists}
}

func (pm *PlaylistManager) Create(ctx context.Context, in *playlist_proto.PlaylistBase) (*playlist_proto.PlaylistResponse, error) {
	pm.logger.Infoln("Playlist Service Create Method entered")

	result, err := pm.repoPlaylist.Create(ctx, DeserializePlaylist(in))
	if err != nil {
		return nil, err
	}
	pm.logger.Infoln("Playlist Created")

	return SerializePlaylistResponse(result), nil
}

func (pm *PlaylistManager) Get(ctx context.Context, in *playlist_proto.PlaylistToUserId) (*playlist_proto.PlaylistResponse, error) {
	pm.logger.Infoln("Playlist Service Get Method entered")

	result, err := pm.repoPlaylist.Get(ctx, in.GetPlaylistId())
	if err != nil {
		return nil, err
	}

	tracks, err := pm.repoTracks.GetByPlaylist(in.GetPlaylistId())
	if err != nil {
		return nil, err
	}

	isCreator, err := pm.repoPlaylist.IsCreator(context.Background(), in.GetUserId(), in.GetPlaylistId())
	if err != nil {
		return nil, err
	}

	return SerializePlaylistResponse(playlist.Response{
		Id:       result.Id,
		Name:     result.Name,
		AuthorId: result.AuthorId,
		Preview:  result.Preview,
		Tracks:   tracks,
		IsYours:  isCreator,
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

	if err := pm.repoPlaylist.AddTrack(ctx, in.GetPlaylistId(), in.GetTrackId()); err != nil {
		return nil, err
	}
	pm.logger.Infoln("Track Added")

	return &google_proto.Empty{}, nil
}

func (pm *PlaylistManager) RemoveTrack(ctx context.Context, in *playlist_proto.PlaylistToTrackId) (*google_proto.Empty, error) {
	pm.logger.Infoln("Playlist Service RemoveTrack Method entered")

	if err := pm.repoPlaylist.RemoveTrack(ctx, in.GetPlaylistId(), in.GetTrackId()); err != nil {
		return nil, err
	}
	pm.logger.Infoln("TrackRemoved")

	return &google_proto.Empty{}, nil
}

func (pm *PlaylistManager) UpdatePreview(ctx context.Context, in *playlist_proto.PlaylistIdToImageUrl) (*google_proto.Empty, error) {
	pm.logger.Infoln("Playlist Service UpdatePreview Method entered")

	if err := pm.repoPlaylist.UpdateImage(ctx, in.GetId(), in.GetUrl().GetUrl()); err != nil {
		return nil, err
	}
	pm.logger.Infoln("Photo Updated")

	return &google_proto.Empty{}, nil
}

func (pm *PlaylistManager) RemovePreview(ctx context.Context, in *playlist_proto.PlaylistId) (*image_proto.ImageUrl, error) {
	pm.logger.Infoln("Playlist Service RemovePreview Method entered")

	avatarPath, err := pm.repoPlaylist.RemovePreviewPath(ctx, in.GetId())
	if err != nil {
		return nil, err
	}
	pm.logger.Infoln("Preview removed")

	return &image_proto.ImageUrl{Url: avatarPath}, nil
}

func (pm *PlaylistManager) DeleteById(ctx context.Context, in *playlist_proto.PlaylistId) (*google_proto.Empty, error) {
	pm.logger.Infoln("Playlist Service DeleteById Method entered")

	if err := pm.repoPlaylist.Delete(ctx, in.GetId()); err != nil {
		return nil, err
	}
	pm.logger.Infoln("Playlist Deleted")

	return &google_proto.Empty{}, nil
}

func (pm *PlaylistManager) Like(ctx context.Context, in *playlist_proto.PlaylistToUserId) (*google_proto.Empty, error) {
	pm.logger.Infoln("Album Micros Like entered")

	if err := pm.repoPlaylist.CreateLike(context.Background(), in.GetUserId(), in.GetPlaylistId()); err != nil {
		return nil, err
	}
	pm.logger.Infoln("Like created")

	return &google_proto.Empty{}, nil
}

func (pm *PlaylistManager) IsLike(ctx context.Context, in *playlist_proto.PlaylistToUserId) (*playlist_proto.IsLikedPlaylist, error) {
	pm.logger.Infoln("Playlist Micros Like entered")

	isLiked, err := pm.repoPlaylist.CheckLike(context.Background(), in.GetUserId(), in.GetPlaylistId())
	if err != nil {
		return nil, err
	}
	pm.logger.Infoln("Like checked")

	return &playlist_proto.IsLikedPlaylist{IsLiked: isLiked}, nil
}

func (pm *PlaylistManager) Unlike(ctx context.Context, in *playlist_proto.PlaylistToUserId) (*google_proto.Empty, error) {
	pm.logger.Infoln("Playlist Micros Like entered")

	if err := pm.repoPlaylist.DeleteLike(context.Background(), in.GetUserId(), in.GetPlaylistId()); err != nil {
		return nil, err
	}
	pm.logger.Infoln("Like deleted")

	return &google_proto.Empty{}, nil
}

func (pm *PlaylistManager) HasModifyAccess(ctx context.Context, in *playlist_proto.PlaylistToUserId) (*playlist_proto.HasAccess, error) {
	pm.logger.Infoln("Album Micros HasModAccess entered")

	isCreator, err := pm.repoPlaylist.IsCreator(ctx, in.GetUserId(), in.GetPlaylistId())
	if err != nil {
		return &playlist_proto.HasAccess{IsAccess: false}, nil
	}

	return &playlist_proto.HasAccess{IsAccess: isCreator}, nil
}

func (pm *PlaylistManager) HasReadAccess(ctx context.Context, in *playlist_proto.PlaylistId) (*playlist_proto.HasAccess, error) {
	pm.logger.Infoln("Album Micros HasReadAccess entered")

	isPrivate, err := pm.repoPlaylist.IsPrivate(ctx, in.GetId())
	if err != nil {
		return &playlist_proto.HasAccess{IsAccess: false}, nil
	}

	return &playlist_proto.HasAccess{IsAccess: !isPrivate}, nil
}

func (pm *PlaylistManager) MakePrivate(ctx context.Context, in *playlist_proto.PlaylistId) (*google_proto.Empty, error) {
	pm.logger.Infoln("Album Micros MakePrivate entered")

	if err := pm.repoPlaylist.MakePrivate(ctx, in.GetId()); err != nil {
		return nil, nil
	}

	return &google_proto.Empty{}, nil
}

func (pm *PlaylistManager) MakePublic(ctx context.Context, in *playlist_proto.PlaylistId) (*google_proto.Empty, error) {
	pm.logger.Infoln("Album Micros MakePublic entered")

	if err := pm.repoPlaylist.MakePublic(ctx, in.GetId()); err != nil {
		return nil, nil
	}

	return &google_proto.Empty{}, nil
}
