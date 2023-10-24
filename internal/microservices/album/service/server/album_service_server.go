package grpc_album_server

import (
	"context"
	google_proto "github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	album_proto "main/internal/microservices/album/proto"
	track_proto "main/internal/microservices/track/proto"
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

func SerializeTrack(in track.Response) *track_proto.Track {
	return &track_proto.Track{
		Id:         in.Id,
		Name:       in.Name,
		Preview:    in.Preview,
		Content:    in.Content,
		ArtistName: in.ArtistName,
		Duration:   in.Duration,
		IsLiked:    in.IsLiked,
	}
}

func SerializeTracks(in []track.Response) *track_proto.TracksResponse {
	var result *track_proto.TracksResponse
	for _, t := range in {
		result.Tracks = append(result.Tracks, SerializeTrack(t))
	}
	return result
}

func SerializeAlbum(in album.Response) *album_proto.AlbumResponse {
	return &album_proto.AlbumResponse{
		Id:         in.Id,
		Name:       in.Name,
		Preview:    in.Preview,
		ArtistId:   in.ArtistId,
		ArtistName: in.ArtistName,
		Tracks:     SerializeTracks(in.Tracks),
	}
}

func SerializeAlbums(in []album.Response) *album_proto.AlbumsResponse {
	var result *album_proto.AlbumsResponse
	for _, a := range in {
		result.Albums = append(result.Albums, SerializeAlbum(a))
	}
	return result
}

func SerializeBase(in album.Base) *album_proto.AlbumBase {
	return &album_proto.AlbumBase{
		Id:      in.Id,
		Name:    in.Name,
		Preview: in.Preview,
	}
}

func SerializeAlbumsBase(in []album.Base) *album_proto.AlbumsBase {
	var result *album_proto.AlbumsBase

	for _, base := range in {
		result.Albums = append(result.Albums, SerializeBase(base))
	}
	return result
}

func (am *AlbumManager) GetAlbum(ctx context.Context, in *album_proto.AlbumId) (*album_proto.AlbumResponse, error) {
	am.logger.Infoln("Album Micros GetAlbum entered")

	var result album.Response

	base, err := am.repoAlbum.Get(in.GetAlbumId())
	if err != nil {
		return nil, err
	}
	am.logger.Infoln("Got album Base")

	result.Id = base.Id
	result.Name = base.Name
	result.Preview = base.Preview

	art, err := am.repoArtist.GetByAlbumId(in.GetAlbumId())
	if err != nil {
		return nil, err
	}
	am.logger.Infoln("Got Artist by album Id ", art)

	result.ArtistId = art.Id
	result.ArtistName = art.Name

	tracks, err := am.repoTrack.GetByAlbum(in.AlbumId)
	if err != nil {
		return nil, err
	}
	am.logger.Infoln("Got tracks by album id ", tracks)

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

	err := am.repoAlbum.CreateLike(in.GetUserId(), in.GetAlbumId())
	if err != nil {
		return nil, err
	}
	am.logger.Infoln("Like created")

	return &google_proto.Empty{}, nil
}
