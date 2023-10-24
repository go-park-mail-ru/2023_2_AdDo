package grpc_track_server

import (
	"context"
	google_proto "github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	track_proto "main/internal/microservices/track/proto"
	"main/internal/pkg/album"
	"main/internal/pkg/artist"
	"main/internal/pkg/track"
)

type TrackManager struct {
	repoTrack  track.Repository
	repoArtist artist.Repository
	repoAlbum  album.Repository
	logger     *logrus.Logger
	track_proto.UnimplementedTrackServiceServer
}

func NewTrackManager(repoTrack track.Repository, repoArtist artist.Repository, repoAlbum album.Repository, logger *logrus.Logger) TrackManager {
	return TrackManager{
		repoTrack:  repoTrack,
		repoArtist: repoArtist,
		repoAlbum:  repoAlbum,
		logger:     logger,
	}
}

func (tm *TrackManager) Listen(ctx context.Context, in *track_proto.TrackId) (*google_proto.Empty, error) {
	tm.logger.Infoln("Track Micros Listen entered")

	err := tm.repoTrack.AddListen(in.GetTrackId())
	if err != nil {
		return nil, err
	}
	tm.logger.Infoln("listen for track ", in.GetTrackId(), " added")

	return &google_proto.Empty{}, nil
}

func (tm *TrackManager) Like(ctx context.Context, in *track_proto.TrackToUserId) (*google_proto.Empty, error) {
	tm.logger.Infoln("Track Micros Like entered")

	err := tm.repoTrack.CreateLike(in.GetUserId(), in.GetTrackId())
	if err != nil {
		return nil, err
	}
	tm.logger.Infoln("Like created for track ", in.GetTrackId(), " by user ", in.GetUserId())

	return &google_proto.Empty{}, nil
}
