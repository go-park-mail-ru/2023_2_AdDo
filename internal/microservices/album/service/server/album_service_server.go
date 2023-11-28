package grpc_album_server

import (
	"context"
	google_proto "github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	album_proto "main/internal/microservices/album/proto"
	session_proto "main/internal/microservices/session/proto"
	track_proto "main/internal/microservices/track/proto"
	grpc_track_server "main/internal/microservices/track/service/server"
	"main/internal/pkg/album"
	"main/internal/pkg/artist"
	"main/internal/pkg/track"
)

type AlbumManager struct {
	repoTrack  track.Repository
	repoArtist artist.Repository
	repoAlbum  album.Repository
	logger     *logrus.Logger
	album_proto.UnimplementedAlbumServiceServer
}

func NewAlbumManager(repoTrack track.Repository, repoArtist artist.Repository, repoAlbum album.Repository, logger *logrus.Logger) AlbumManager {
	return AlbumManager{
		repoTrack:  repoTrack,
		repoArtist: repoArtist,
		repoAlbum:  repoAlbum,
		logger:     logger,
	}
}

func SerializeAlbum(in album.Response) *album_proto.AlbumResponse {
	return &album_proto.AlbumResponse{
		Id:         in.Id,
		Name:       in.Name,
		Preview:    in.Preview,
		ArtistId:   in.ArtistId,
		ArtistName: in.ArtistName,
		Tracks:     grpc_track_server.SerializeTracks(in.Tracks),
	}
}

func SerializeAlbums(in []album.Response) *album_proto.AlbumsResponse {
	albums := make([]*album_proto.AlbumResponse, 0)
	for _, a := range in {
		albums = append(albums, SerializeAlbum(a))
	}
	return &album_proto.AlbumsResponse{Albums: albums}
}

func SerializeBase(in album.Base) *album_proto.AlbumBase {
	return &album_proto.AlbumBase{
		Id:      in.Id,
		Name:    in.Name,
		Preview: in.Preview,
	}
}

func SerializeAlbumsBase(in []album.Base) *album_proto.AlbumsBase {
	albums := make([]*album_proto.AlbumBase, 0)
	for _, base := range in {
		albums = append(albums, SerializeBase(base))
	}
	return &album_proto.AlbumsBase{Albums: albums}
}

func (am *AlbumManager) GetAlbum(ctx context.Context, in *album_proto.AlbumId) (*album_proto.AlbumResponse, error) {
	am.logger.Infoln("Album Micros GetAlbum entered")

	base, err := am.repoAlbum.Get(in.GetAlbumId())
	if err != nil {
		return nil, err
	}
	am.logger.Infoln("Got album Base")

	return am.formResponseAlbumWithAllTracks(base)
}

func (am *AlbumManager) GetAlbumByTrack(ctx context.Context, in *track_proto.TrackId) (*album_proto.AlbumResponse, error) {
	am.logger.Infoln("Album Micros GetAlbumByTrack entered")

	albumsBase, err := am.repoAlbum.GetByTrackId(in.GetTrackId())
	if err != nil {
		return nil, err
	}
	am.logger.Infoln("Got albums with required track")

	return am.formResponseAlbumWithAllTracks(albumsBase[0])
}

func (am *AlbumManager) formResponseAlbumWithAllTracks(albumBase album.Base) (*album_proto.AlbumResponse, error) {
	var result album.Response

	result.Id = albumBase.Id
	result.Name = albumBase.Name
	result.Preview = albumBase.Preview

	art, err := am.repoArtist.GetByAlbumId(albumBase.Id)
	if err != nil {
		return nil, err
	}
	am.logger.Infoln("Got artist by album id", art)

	result.ArtistId = art.Id
	result.ArtistName = art.Name

	tracks, err := am.repoTrack.GetByAlbum(albumBase.Id)
	if err != nil {
		return nil, err
	}
	am.logger.Infoln("Got tracks by album id", tracks)

	result.Tracks = tracks

	return SerializeAlbum(result), nil
}

func (am *AlbumManager) GetRandom(ctx context.Context, status *google_proto.Empty) (*album_proto.AlbumsResponse, error) {
	am.logger.Infoln("Album Micros GetRandom entered")

	albums, err := am.repoAlbum.GetRandom(album.LimitForMainPage)
	if err != nil {
		return nil, err
	}
	am.logger.Infoln("Got random albums")

	return am.formResponse(albums)
}

func (am *AlbumManager) GetMostLiked(ctx context.Context, status *google_proto.Empty) (*album_proto.AlbumsResponse, error) {
	am.logger.Infoln("Album Micros GetMostLiked entered")

	albums, err := am.repoAlbum.GetByLikeCount(album.LimitForMainPage)
	if err != nil {
		return nil, err
	}
	am.logger.Infoln("Got album")

	return am.formResponse(albums)
}

func (am *AlbumManager) GetByUserId(ctx context.Context, id *session_proto.UserId) (*album_proto.AlbumsBase, error) {
	am.logger.Infoln("Album Micros GetByUserId entered")

	albums, err := am.repoAlbum.GetByUserId(id.GetUserId())
	if err != nil {
		return nil, err
	}
	am.logger.Infoln("Got albums by user id")

	return SerializeAlbumsBase(albums), nil
}

func (am *AlbumManager) GetPopular(ctx context.Context, status *google_proto.Empty) (*album_proto.AlbumsResponse, error) {
	am.logger.Infoln("Album Micros GetPopular entered")

	albums, err := am.repoAlbum.GetByListenCount(album.LimitForMainPage)
	if err != nil {
		return nil, err
	}
	am.logger.Infoln("Got albums by Listen count")

	return am.formResponse(albums)
}

func (am *AlbumManager) GetNew(ctx context.Context, status *google_proto.Empty) (*album_proto.AlbumsResponse, error) {
	am.logger.Infoln("Album Micros GetNew entered")

	albums, err := am.repoAlbum.GetByReleaseDate(album.LimitForMainPage)
	if err != nil {
		return nil, err
	}
	am.logger.Infoln("Got new albums")

	return am.formResponse(albums)
}

func (am *AlbumManager) formResponse(albumBase []album.Base) (*album_proto.AlbumsResponse, error) {
	am.logger.Infoln("Album Micros fromResponse entered")

	result := make([]album.Response, 0)
	for _, base := range albumBase {
		var a album.Response
		art, err := am.repoArtist.GetByAlbumId(base.Id)
		if err != nil {
			return nil, err
		}
		am.logger.Infoln("artist founded")

		a.Id = base.Id
		a.Name = base.Name
		a.Preview = base.Preview
		a.ArtistId = art.Id
		a.ArtistName = art.Name

		result = append(result, a)
	}
	am.logger.Infoln("response formed")

	return SerializeAlbums(result), nil
}

func (am *AlbumManager) Like(ctx context.Context, in *album_proto.AlbumToUserId) (*google_proto.Empty, error) {
	am.logger.Infoln("Album Micros Like entered")

	if err := am.repoAlbum.CreateLike(in.GetUserId(), in.GetAlbumId()); err != nil {
		return nil, err
	}
	am.logger.Infoln("Like created")

	return &google_proto.Empty{}, nil
}

func (am *AlbumManager) IsLike(ctx context.Context, in *album_proto.AlbumToUserId) (*album_proto.IsLikedAlbum, error) {
	am.logger.Infoln("Album Micros Like entered")

	result, err := am.repoAlbum.CheckLike(in.GetUserId(), in.GetAlbumId())
	if err != nil {
		return nil, err
	}
	am.logger.Infoln("Like created")

	return &album_proto.IsLikedAlbum{IsLiked: result}, nil
}

func (am *AlbumManager) Unlike(ctx context.Context, in *album_proto.AlbumToUserId) (*google_proto.Empty, error) {
	am.logger.Infoln("Album Micros Like entered")

	if err := am.repoAlbum.DeleteLike(in.GetUserId(), in.GetAlbumId()); err != nil {
		return nil, err
	}
	am.logger.Infoln("Like created")

	return &google_proto.Empty{}, nil
}
