package grpc_artist_server

import (
	"context"
	google_proto "github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	grpc_album_server "main/internal/microservices/album/service/server"
	artist_proto "main/internal/microservices/artist/proto"
	grpc_playlist_server "main/internal/microservices/playlist/service/server"
	session_proto "main/internal/microservices/session/proto"
	grpc_track_server "main/internal/microservices/track/service/server"
	"main/internal/pkg/activity"
	"main/internal/pkg/album"
	"main/internal/pkg/artist"
	"main/internal/pkg/playlist"
	"main/internal/pkg/track"
)

type ArtistManager struct {
	repoArtist   artist.Repository
	repoTrack    track.Repository
	repoAlbum    album.Repository
	repoPlaylist playlist.Repository
	logger       *logrus.Logger
	artist_proto.UnimplementedArtistServiceServer
	queue activity.ProducerRepository
}

func NewArtistManager(q activity.ProducerRepository, repoPlaylist playlist.Repository, repoArtist artist.Repository, repoTrack track.Repository, repoAlbum album.Repository, logger *logrus.Logger) ArtistManager {
	return ArtistManager{
		repoArtist:   repoArtist,
		repoTrack:    repoTrack,
		repoAlbum:    repoAlbum,
		repoPlaylist: repoPlaylist,
		logger:       logger,
		queue:        q,
	}
}

func SerializeArtist(in artist.Response) *artist_proto.Artist {
	return &artist_proto.Artist{
		Id:     in.Id,
		Name:   in.Name,
		Avatar: in.Avatar,
		Albums: grpc_album_server.SerializeAlbumsBase(in.Albums),
		Tracks: grpc_track_server.SerializeTracks(in.Tracks),
	}
}

func SerializeArtistBase(in artist.Base) *artist_proto.ArtistBase {
	return &artist_proto.ArtistBase{
		Id:     in.Id,
		Name:   in.Name,
		Avatar: in.Avatar,
	}
}

func SerializeArtistsBase(in []artist.Base) *artist_proto.ArtistsBase {
	result := make([]*artist_proto.ArtistBase, 0)
	for _, base := range in {
		result = append(result, SerializeArtistBase(base))
	}
	return &artist_proto.ArtistsBase{
		Artists: result,
	}
}

func SerializeSearchResponse(in artist.SearchResponse) *artist_proto.SearchResponse {
	return &artist_proto.SearchResponse{
		Tracks:    grpc_track_server.SerializeTracks(in.Tracks),
		Albums:    grpc_album_server.SerializeAlbumsBase(in.Albums),
		Playlists: grpc_playlist_server.SerializePlaylistsBase(in.Playlists),
		Artists:   SerializeArtistsBase(in.Artists),
	}
}

func (am *ArtistManager) GetArtistInfo(ctx context.Context, in *artist_proto.ArtistId) (*artist_proto.Artist, error) {
	am.logger.Infoln("Artist Micros GetArtistInfo entered")

	var result artist.Response

	artistBase, err := am.repoArtist.Get(in.GetArtistId())
	if err != nil {
		return nil, err
	}
	am.logger.Infoln("Got artist base")

	result.Id = artistBase.Id
	result.Name = artistBase.Name
	result.Avatar = artistBase.Avatar

	albums, err := am.repoAlbum.GetByArtistId(in.GetArtistId())
	if err != nil {
		return nil, err
	}
	am.logger.Infoln("Got artist albums")

	result.Albums = albums

	tracks, err := am.repoTrack.GetByArtist(in.GetArtistId())
	if err != nil {
		return nil, err
	}
	am.logger.Infoln("Got artist tracks")

	result.Tracks = tracks

	return SerializeArtist(result), nil
}

func (am *ArtistManager) Like(ctx context.Context, in *artist_proto.ArtistToUserId) (*google_proto.Empty, error) {
	am.logger.Infoln("Artist Micros Like entered")

	if err := am.repoArtist.CreateLike(in.GetUserId(), in.GetArtistId()); err != nil {
		return nil, err
	}
	am.logger.Infoln("Like created")

	if err := am.queue.PushLikeArtist(in.GetUserId(), in.GetArtistId()); err != nil {
		return nil, err
	}
	am.logger.Infoln("Like Pushed to queue")

	return &google_proto.Empty{}, nil
}

func (am *ArtistManager) IsLike(ctx context.Context, in *artist_proto.ArtistToUserId) (*artist_proto.IsLikedArtist, error) {
	am.logger.Infoln("Artist Micros IsLike entered")

	isLiked, err := am.repoArtist.CheckLike(in.GetUserId(), in.GetArtistId())
	if err != nil {
		return nil, err
	}
	am.logger.Infoln("Like checked")

	return &artist_proto.IsLikedArtist{IsLiked: isLiked}, nil
}

func (am *ArtistManager) Unlike(ctx context.Context, in *artist_proto.ArtistToUserId) (*google_proto.Empty, error) {
	am.logger.Infoln("Artist Micros Like entered")

	if err := am.repoArtist.DeleteLike(in.GetUserId(), in.GetArtistId()); err != nil {
		return nil, err
	}
	am.logger.Infoln("Like created")

	return &google_proto.Empty{}, nil
}

func (am *ArtistManager) FullSearch(ctx context.Context, in *artist_proto.Query) (*artist_proto.SearchResponse, error) {
	am.logger.Infoln("Artist Micros FullSearch entered")

	tracks, err := am.repoTrack.Search(in.GetQuery())
	if err != nil {
		return nil, err
	}

	albums, err := am.repoAlbum.Search(in.GetQuery())
	if err != nil {
		return nil, err
	}

	playlists, err := am.repoPlaylist.Search(ctx, in.GetQuery())
	if err != nil {
		return nil, err
	}

	artists, err := am.repoArtist.Search(in.GetQuery())
	if err != nil {
		return nil, err
	}

	result := artist.SearchResponse{
		Playlists: playlists,
		Albums:    albums,
		Tracks:    tracks,
		Artists:   artists,
	}
	return SerializeSearchResponse(result), nil
}

func (am *ArtistManager) CollectionArtist(ctx context.Context, in *session_proto.UserId) (*artist_proto.ArtistsBase, error) {
	am.logger.Infoln("Artist Micros Collection Artist entered")

	artists, err := am.repoArtist.GetByUserId(in.GetUserId())
	if err != nil {
		return nil, err
	}

	return SerializeArtistsBase(artists), nil
}
