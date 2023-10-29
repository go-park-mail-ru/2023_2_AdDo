package grpc_artist_server

import (
	"context"
	"github.com/sirupsen/logrus"
	grpc_album_server "main/internal/microservices/album/service/server"
	artist_proto "main/internal/microservices/artist/proto"
	"main/internal/pkg/album"
	"main/internal/pkg/artist"
	"main/internal/pkg/track"
)

type ArtistManager struct {
	repoArtist artist.Repository
	repoTrack  track.Repository
	repoAlbum  album.Repository
	logger     *logrus.Logger
	artist_proto.UnimplementedArtistServiceServer
}

func NewArtistManager(repoArtist artist.Repository, repoTrack track.Repository, repoAlbum album.Repository, logger *logrus.Logger) ArtistManager {
	return ArtistManager{
		repoArtist: repoArtist,
		repoTrack:  repoTrack,
		repoAlbum:  repoAlbum,
		logger:     logger,
	}
}

func SerializeArtist(in artist.Response) *artist_proto.Artist {
	return &artist_proto.Artist{
		Id:     in.Id,
		Name:   in.Name,
		Avatar: in.Avatar,
		Albums: grpc_album_server.SerializeAlbumsBase(in.Albums),
		Tracks: grpc_album_server.SerializeTracks(in.Tracks),
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

	for index, _ := range tracks {
		tracks[index].ArtistName = artistBase.Name
	}

	result.Tracks = tracks

	return SerializeArtist(result), nil
}
