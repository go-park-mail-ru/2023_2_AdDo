package grpc_track_server

import (
	"context"
	"github.com/golang/mock/gomock"
	google_proto "github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	session_proto "main/internal/microservices/session/proto"
	track_proto "main/internal/microservices/track/proto"
	"main/internal/pkg/track"
	activity_mock "main/test/mocks/activity"
	track_mock "main/test/mocks/track"
	"testing"
)

func Test_TrackServiceServer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTrackRepo := track_mock.NewMockRepository(ctrl)
	mockProducerRepo := activity_mock.NewMockProducerRepository(ctrl)

	trackManager := TrackManager{
		repoTrack: mockTrackRepo,
		queue:     mockProducerRepo,
		logger:    logrus.New(),
	}

	const (
		trackId  uint64 = 1
		userId   string = "1e31379e-79ca-40d0-96db-6cb9e3e8f492"
		duration uint32 = 40
	)

	in := &track_proto.TrackToUserDur{
		TrackToUser: &track_proto.TrackToUserId{UserId: userId, TrackId: trackId},
		Duration:    duration,
	}

	t.Run("Listen", func(t *testing.T) {
		mockTrackRepo.EXPECT().AddListen(in.GetTrackToUser().GetTrackId()).Return(nil)
		mockTrackRepo.EXPECT().
			CreateListen(in.GetTrackToUser().GetUserId(), in.GetTrackToUser().GetTrackId(), in.GetDuration()).
			Return(nil)

		mockProducerRepo.EXPECT().
			PushListenTrack(in.GetTrackToUser().GetUserId(), in.GetTrackToUser().GetTrackId(), in.GetDuration()).
			Return(nil)

		result, err := trackManager.Listen(context.Background(), in)
		assert.Nil(t, err)
		assert.Equal(t, &google_proto.Empty{}, result)
	})

	t.Run("Skip", func(t *testing.T) {
		mockTrackRepo.EXPECT().
			CreateSkip(in.GetTrackToUser().GetUserId(), in.GetTrackToUser().GetTrackId(), in.GetDuration()).
			Return(nil)

		mockProducerRepo.EXPECT().
			PushSkipTrack(in.GetTrackToUser().GetUserId(), in.GetTrackToUser().GetTrackId(), in.GetDuration()).
			Return(nil)

		result, err := trackManager.Skip(context.Background(), in)
		assert.Nil(t, err)
		assert.Equal(t, &google_proto.Empty{}, result)
	})

	t.Run("Like", func(t *testing.T) {
		in := &track_proto.TrackToUserId{TrackId: trackId, UserId: userId}

		mockTrackRepo.EXPECT().CreateLike(in.GetUserId(), in.GetTrackId()).Return(nil)
		mockProducerRepo.EXPECT().PushLikeTrack(in.GetUserId(), in.GetTrackId()).Return(nil)

		result, err := trackManager.Like(context.Background(), in)
		assert.Nil(t, err)
		assert.Equal(t, &google_proto.Empty{}, result)
	})

	t.Run("IsLike", func(t *testing.T) {
		in := &track_proto.TrackToUserId{TrackId: trackId, UserId: userId}

		mockTrackRepo.EXPECT().CheckLike(in.GetUserId(), in.GetTrackId()).Return(true, nil)

		result, err := trackManager.IsLike(context.Background(), in)
		assert.Nil(t, err)
		assert.Equal(t, &track_proto.IsLikedTrack{IsLiked: true}, result)
	})

	t.Run("UnLike", func(t *testing.T) {
		in := &track_proto.TrackToUserId{TrackId: trackId, UserId: userId}

		mockTrackRepo.EXPECT().DeleteLike(in.GetUserId(), in.GetTrackId()).Return(nil)

		result, err := trackManager.Unlike(context.Background(), in)
		assert.Nil(t, err)
		assert.Equal(t, &google_proto.Empty{}, result)
	})
}

func Test_GettingTracks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTrackRepo := track_mock.NewMockRepository(ctrl)
	mockProducerRepo := activity_mock.NewMockProducerRepository(ctrl)

	trackManager := TrackManager{
		repoTrack: mockTrackRepo,
		logger:    logrus.New(),
		queue:     mockProducerRepo,
	}

	const userId = "1e31379e-79ca-40d0-96db-6cb9e3e8f492"

	tracks := []track.Response{
		{
			Id:       1,
			Name:     "Track1",
			Content:  "Content1",
			Duration: "1:40",
		},
		{
			Id:       2,
			Name:     "Track2",
			Content:  "Content2",
			Duration: "2:40",
		}}

	t.Run("GetUserLikedTracks", func(t *testing.T) {
		in := &session_proto.UserId{UserId: userId}

		mockTrackRepo.EXPECT().GetByUser(in.GetUserId()).Return(tracks, nil)

		result, err := trackManager.GetUserLikedTracks(context.Background(), in)
		assert.Nil(t, err)
		assert.Equal(t, tracks, DeserializeTracks(result))
	})

	t.Run("LabelIsLikedForUser", func(t *testing.T) {
		in := track_proto.UserToTracksForLabeling{
			UserId: userId,
			Tracks: SerializeTracks(tracks),
		}

		tracks[0].IsLiked = true
		tracks[1].IsLiked = true
		mockTrackRepo.EXPECT().LabelIsLikedTracks(in.GetUserId(), DeserializeTracks(in.GetTracks())).Return(tracks, nil)

		result, err := trackManager.LabelIsLikedForUser(context.Background(), &in)
		assert.Nil(t, err)

		deserializedResult := DeserializeTracks(result)
		assert.Equal(t, tracks, deserializedResult)
		assert.Equal(t, true, deserializedResult[0].IsLiked)
		assert.Equal(t, true, deserializedResult[1].IsLiked)
	})
}
