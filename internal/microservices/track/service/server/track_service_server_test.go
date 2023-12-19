package grpc_track_server

import (
	"context"
	"github.com/golang/mock/gomock"
	google_proto "github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	pb "main/internal/microservices/track/proto"
	activity_mock "main/test/mocks/activity"
	album_mock "main/test/mocks/album"
	artist_mock "main/test/mocks/artist"
	track_mock "main/test/mocks/track"
	"testing"
)

func Test_TrackServiceServer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTrackRepo := track_mock.NewMockRepository(ctrl)
	mockArtistRepo := artist_mock.NewMockRepository(ctrl)
	mockAlbumRepo := album_mock.NewMockRepository(ctrl)
	mockProducerRepo := activity_mock.NewMockProducerRepository(ctrl)

	trackManager := TrackManager{
		repoTrack:  mockTrackRepo,
		repoArtist: mockArtistRepo,
		repoAlbum:  mockAlbumRepo,
		logger:     logrus.New(),
		queue:      mockProducerRepo,
	}

	const (
		trackId  uint64 = 1
		userId   string = "1e31379e-79ca-40d0-96db-6cb9e3e8f492"
		duration uint32 = 40
	)

	in := &pb.TrackToUserDur{
		TrackToUser: &pb.TrackToUserId{UserId: userId, TrackId: trackId},
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
		in := &pb.TrackToUserId{TrackId: trackId, UserId: userId}

		mockTrackRepo.EXPECT().CreateLike(in.GetUserId(), in.GetTrackId()).Return(nil)
		mockProducerRepo.EXPECT().PushLikeTrack(in.GetUserId(), in.GetTrackId()).Return(nil)

		result, err := trackManager.Like(context.Background(), in)
		assert.Nil(t, err)
		assert.Equal(t, &google_proto.Empty{}, result)
	})

	t.Run("IsLike", func(t *testing.T) {
		in := &pb.TrackToUserId{TrackId: trackId, UserId: userId}

		mockTrackRepo.EXPECT().CheckLike(in.GetUserId(), in.GetTrackId()).Return(true, nil)

		result, err := trackManager.IsLike(context.Background(), in)
		assert.Nil(t, err)
		assert.Equal(t, &pb.IsLikedTrack{IsLiked: true}, result)
	})

	t.Run("UnLike", func(t *testing.T) {
		in := &pb.TrackToUserId{TrackId: trackId, UserId: userId}

		mockTrackRepo.EXPECT().DeleteLike(in.GetUserId(), in.GetTrackId()).Return(nil)

		result, err := trackManager.Unlike(context.Background(), in)
		assert.Nil(t, err)
		assert.Equal(t, &google_proto.Empty{}, result)
	})
}
