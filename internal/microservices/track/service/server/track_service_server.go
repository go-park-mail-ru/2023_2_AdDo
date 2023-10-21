package grpc_track_server

import (
	"context"
	"github.com/sirupsen/logrus"
	pb "main/internal/microservices/track/proto"
	"main/internal/pkg/album"
	"main/internal/pkg/artist"
	"main/internal/pkg/track"
)

type TrackManager struct {
	repoTrack  track.Repository
	repoArtist artist.Repository
	repoAlbum  album.Repository
	logger     *logrus.Logger
	pb.UnimplementedTrackServiceServer
}

func NewTrackManager(repoTrack track.Repository, repoArtist artist.Repository, repoAlbum album.Repository, logger *logrus.Logger) TrackManager {
	return TrackManager{
		repoTrack:  repoTrack,
		repoArtist: repoArtist,
		repoAlbum:  repoAlbum,
		logger:     logger,
	}
}

func (tm *TrackManager) Listen(ctx context.Context, in *pb.TrackId) (*pb.Status, error) {
	tm.logger.Infoln("Track Micros Listen entered")

	err := tm.repoTrack.AddListen(in.GetTrackId())
	if err != nil {
		return nil, err
	}
	tm.logger.Infoln("listen for track ", in.GetTrackId(), " added")

	return &pb.Status{IsOk: true}, nil
}

func (tm *TrackManager) Like(ctx context.Context, in *pb.TrackToUserId) (*pb.Status, error) {
	tm.logger.Infoln("Track Micros Like entered")

	err := tm.repoTrack.CreateLike(in.GetUserId(), in.GetTrackId())
	if err != nil {
		return nil, err
	}
	tm.logger.Infoln("Like created for track ", in.GetTrackId(), " by user ", in.GetUserId())

	return &pb.Status{IsOk: true}, nil
}
