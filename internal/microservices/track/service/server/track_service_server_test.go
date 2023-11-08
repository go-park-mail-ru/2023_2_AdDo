package grpc_track_server

import (
	"context"
	"github.com/golang/mock/gomock"
	google_proto "github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	pb "main/internal/microservices/track/proto"
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

	trackManager := TrackManager{
		repoTrack:  mockTrackRepo,
		repoArtist: mockArtistRepo,
		repoAlbum:  mockAlbumRepo,
		logger:     logrus.New(),
	}

	t.Run("Listen", func(t *testing.T) {
		in := &pb.TrackId{TrackId: 1}
		mockTrackRepo.EXPECT().AddListen(in.GetTrackId()).Return(nil)

		result, err := trackManager.Listen(context.Background(), in)
		assert.Nil(t, err)
		assert.Equal(t, &google_proto.Empty{}, result)
	})

	in := &pb.TrackToUserId{TrackId: 1, UserId: "user"}

	t.Run("Like", func(t *testing.T) {
		mockTrackRepo.EXPECT().CreateLike(in.GetUserId(), in.GetTrackId()).Return(nil)

		result, err := trackManager.Like(context.Background(), in)
		assert.Nil(t, err)
		assert.Equal(t, &google_proto.Empty{}, result)
	})

	t.Run("IsLike", func(t *testing.T) {
		mockTrackRepo.EXPECT().CheckLike(in.GetUserId(), in.GetTrackId()).Return(true, nil)

		result, err := trackManager.IsLike(context.Background(), in)
		assert.Nil(t, err)
		assert.Equal(t, &pb.IsLikedTrack{IsLiked: true}, result)
	})

	t.Run("UnLike", func(t *testing.T) {
		mockTrackRepo.EXPECT().DeleteLike(in.GetUserId(), in.GetTrackId()).Return(nil)

		result, err := trackManager.Unlike(context.Background(), in)
		assert.Nil(t, err)
		assert.Equal(t, &google_proto.Empty{}, result)
	})
}
