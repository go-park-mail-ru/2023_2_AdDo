package grpc_track_server

import (
	"context"
	google_proto "github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	session_proto "main/internal/microservices/session/proto"
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
	tracks := make([]*track_proto.Track, 0)
	for _, t := range in {
		tracks = append(tracks, SerializeTrack(t))
	}
	return &track_proto.TracksResponse{Tracks: tracks}
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

func (tm *TrackManager) IsLike(ctx context.Context, in *track_proto.TrackToUserId) (*track_proto.IsLikedTrack, error) {
	tm.logger.Infoln("Track Micros Like entered")

	isLiked, err := tm.repoTrack.CheckLike(in.GetUserId(), in.GetTrackId())
	if err != nil {
		return nil, err
	}
	tm.logger.Infoln("Like checked for track ", in.GetTrackId(), " by user ", in.GetUserId())

	return &track_proto.IsLikedTrack{IsLiked: isLiked}, nil
}

func (tm *TrackManager) Unlike(ctx context.Context, in *track_proto.TrackToUserId) (*google_proto.Empty, error) {
	tm.logger.Infoln("Track Micros Like entered")

	err := tm.repoTrack.DeleteLike(in.GetUserId(), in.GetTrackId())
	if err != nil {
		return nil, err
	}
	tm.logger.Infoln("Like deleted for track ", in.GetTrackId(), " by user ", in.GetUserId())

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

func (tm *TrackManager) GetUserLikedTracks(ctx context.Context, in *session_proto.UserId) (*track_proto.TracksResponse, error) {
	tm.logger.Infoln("Track Micros GetUserLikedTracks entered")

	result, err := tm.repoTrack.GetByUser(in.GetUserId())
	if err != nil {
		tm.logger.Errorln(err)
		return nil, err
	}

	return SerializeTracks(result), nil
}
