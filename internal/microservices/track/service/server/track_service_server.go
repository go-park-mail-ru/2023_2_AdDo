package grpc_track_server

import (
	"context"
	google_proto "github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	session_proto "main/internal/microservices/session/proto"
	track_proto "main/internal/microservices/track/proto"
	"main/internal/pkg/activity"
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
	queue activity.ProducerRepository
}

func NewTrackManager(q activity.ProducerRepository, repoTrack track.Repository, repoArtist artist.Repository, repoAlbum album.Repository, logger *logrus.Logger) TrackManager {
	return TrackManager{
		repoTrack:  repoTrack,
		repoArtist: repoArtist,
		repoAlbum:  repoAlbum,
		logger:     logger,
		queue:      q,
	}
}

func SerializeTrack(in track.Response) *track_proto.Track {
	return &track_proto.Track{
		Id:         in.Id,
		Name:       in.Name,
		Preview:    in.Preview,
		Content:    in.Content,
		ArtistId:   in.ArtistId,
		ArtistName: in.ArtistName,
		Duration:   in.Duration,
		IsLiked:    in.IsLiked,
		Lyrics:     in.Lyrics,
	}
}

func SerializeTracks(in []track.Response) *track_proto.TracksResponse {
	tracks := make([]*track_proto.Track, 0)
	for _, t := range in {
		tracks = append(tracks, SerializeTrack(t))
	}
	return &track_proto.TracksResponse{Tracks: tracks}
}

func DeserializeTrack(in *track_proto.Track) track.Response {
	return track.Response{
		Id:         in.GetId(),
		Name:       in.GetName(),
		Preview:    in.GetPreview(),
		Content:    in.GetContent(),
		ArtistId:   in.GetArtistId(),
		ArtistName: in.GetArtistName(),
		Duration:   in.GetDuration(),
		IsLiked:    in.GetIsLiked(),
		Lyrics:     in.GetLyrics(),
	}
}

func DeserializeTracks(in *track_proto.TracksResponse) []track.Response {
	result := make([]track.Response, 0)
	for _, t := range in.GetTracks() {
		result = append(result, DeserializeTrack(t))
	}
	return result
}

func (tm *TrackManager) Listen(ctx context.Context, in *track_proto.TrackToUserDur) (*google_proto.Empty, error) {
	tm.logger.Infoln("Track Micros Listen entered")

	if err := tm.repoTrack.AddListen(in.GetTrackToUser().GetTrackId()); err != nil {
		return nil, err
	}
	tm.logger.Infoln("listen for track ", in.GetTrackToUser().GetTrackId(), " added")

	if err := tm.repoTrack.CreateListen(in.GetTrackToUser().GetUserId(), in.GetTrackToUser().GetTrackId(), in.GetDuration()); err != nil {
		return nil, err
	}
	tm.logger.Infoln("created listen track")

	if err := tm.queue.PushListenTrack(in.GetTrackToUser().GetUserId(), in.GetTrackToUser().GetTrackId(), in.GetDuration()); err != nil {
		return nil, err
	}
	tm.logger.Infoln("pushed skip track")

	return &google_proto.Empty{}, nil
}

func (tm *TrackManager) Skip(ctx context.Context, in *track_proto.TrackToUserDur) (*google_proto.Empty, error) {
	tm.logger.Infoln("Track Micros Skip entered")

	if err := tm.repoTrack.CreateSkip(in.GetTrackToUser().GetUserId(), in.GetTrackToUser().GetTrackId(), in.GetDuration()); err != nil {
		return nil, err
	}
	tm.logger.Infoln("created skip track")

	if err := tm.queue.PushSkipTrack(in.GetTrackToUser().GetUserId(), in.GetTrackToUser().GetTrackId(), in.GetDuration()); err != nil {
		return nil, err
	}
	tm.logger.Infoln("pushed skip track")
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

	if err := tm.repoTrack.DeleteLike(in.GetUserId(), in.GetTrackId()); err != nil {
		return nil, err
	}
	tm.logger.Infoln("Like deleted for track ", in.GetTrackId(), " by user ", in.GetUserId())

	return &google_proto.Empty{}, nil
}

func (tm *TrackManager) Like(ctx context.Context, in *track_proto.TrackToUserId) (*google_proto.Empty, error) {
	tm.logger.Infoln("Track Micros Like entered")

	if err := tm.repoTrack.CreateLike(in.GetUserId(), in.GetTrackId()); err != nil {
		return nil, err
	}
	tm.logger.Infoln("Like created for track ", in.GetTrackId(), " by user ", in.GetUserId())

	if err := tm.queue.PushLikeTrack(in.GetUserId(), in.GetTrackId()); err != nil {
		return nil, err
	}
	tm.logger.Infoln("Like pushed to q ", in.GetTrackId(), " by user ", in.GetUserId())

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

func (tm *TrackManager) LabelIsLikedForUser(ctx context.Context, in *track_proto.UserToTracksForLabeling) (*track_proto.TracksResponse, error) {
	tm.logger.Infoln("Track Micros Label Tracks entered")

	result, err := tm.repoTrack.LabelIsLikedTracks(in.GetUserId(), DeserializeTracks(in.GetTracks()))
	if err != nil {
		tm.logger.Errorln(err)
		return nil, err
	}

	return SerializeTracks(result), nil
}
