package grpc_artist_server

import (
	"context"
	"github.com/golang/mock/gomock"
	google_proto "github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	album_proto "main/internal/microservices/album/proto"
	proto "main/internal/microservices/artist/proto"
	track_proto "main/internal/microservices/track/proto"
	"main/internal/pkg/album"
	"main/internal/pkg/artist"
	"main/internal/pkg/track"
	activity_mock "main/test/mocks/activity"
	album_mock "main/test/mocks/album"
	artist_mock "main/test/mocks/artist"
	track_mock "main/test/mocks/track"
	"testing"
)

func Test_GetArtistInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTrackRepo := track_mock.NewMockRepository(ctrl)
	mockArtistRepo := artist_mock.NewMockRepository(ctrl)
	mockAlbumRepo := album_mock.NewMockRepository(ctrl)

	artistManager := ArtistManager{
		repoTrack:  mockTrackRepo,
		repoArtist: mockArtistRepo,
		repoAlbum:  mockAlbumRepo,
		logger:     logrus.New(),
	}

	art := artist.Base{
		Id:     1,
		Name:   "Artist",
		Avatar: "Avatar",
	}

	albums := []album.Base{
		{
			Id:      1,
			Name:    "AlbumName",
			Preview: "Preview",
		},
	}

	tracks := []track.Response{
		{
			Id:         1,
			Name:       "Track1",
			Preview:    "Preview1",
			Content:    "Content1",
			ArtistName: "Artist",
			Duration:   "1:40",
			IsLiked:    true,
		},
	}

	in := &proto.ArtistId{ArtistId: 1}
	mockArtistRepo.EXPECT().Get(in.GetArtistId()).Return(art, nil)
	mockAlbumRepo.EXPECT().GetByArtistId(in.GetArtistId()).Return(albums, nil)
	mockTrackRepo.EXPECT().GetByArtist(in.GetArtistId()).Return(tracks, nil)

	expectedSerializedArtist := &proto.Artist{
		Id:     1,
		Name:   "Artist",
		Avatar: "Avatar",
		Albums: &album_proto.AlbumsBase{
			Albums: []*album_proto.AlbumBase{
				{
					Id:      1,
					Name:    "AlbumName",
					Preview: "Preview",
				},
			},
		},
		Tracks: &track_proto.TracksResponse{
			Tracks: []*track_proto.Track{
				{
					Id:         1,
					Name:       "Track1",
					Preview:    "Preview1",
					Content:    "Content1",
					ArtistName: "Artist",
					Duration:   "1:40",
					IsLiked:    true,
				},
			},
		},
	}

	receivedArtist, err := artistManager.GetArtistInfo(context.Background(), in)

	assert.Equal(t, nil, err)
	assert.Equal(t, expectedSerializedArtist, receivedArtist)
}

func Test_Like(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTrackRepo := track_mock.NewMockRepository(ctrl)
	mockArtistRepo := artist_mock.NewMockRepository(ctrl)
	mockAlbumRepo := album_mock.NewMockRepository(ctrl)
	mockProducerRepo := activity_mock.NewMockProducerRepository(ctrl)

	artistManager := ArtistManager{
		repoTrack:  mockTrackRepo,
		repoArtist: mockArtistRepo,
		repoAlbum:  mockAlbumRepo,
		logger:     logrus.New(),
		queue:      mockProducerRepo,
	}

	in := &proto.ArtistToUserId{UserId: "user", ArtistId: 2}

	t.Run("Like", func(t *testing.T) {
		mockArtistRepo.EXPECT().CreateLike(in.GetUserId(), in.GetArtistId()).Return(nil)
		mockProducerRepo.EXPECT().PushLikeArtist(in.GetUserId(), in.GetArtistId()).Return(nil)

		result, err := artistManager.Like(context.Background(), in)
		assert.Nil(t, err)
		assert.Equal(t, &google_proto.Empty{}, result)
	})

	t.Run("IsLike", func(t *testing.T) {
		mockArtistRepo.EXPECT().CheckLike(in.GetUserId(), in.GetArtistId()).Return(true, nil)

		result, err := artistManager.IsLike(context.Background(), in)
		assert.Nil(t, err)
		assert.Equal(t, &proto.IsLikedArtist{IsLiked: true}, result)
	})

	t.Run("UnLike", func(t *testing.T) {
		mockArtistRepo.EXPECT().DeleteLike(in.GetUserId(), in.GetArtistId()).Return(nil)

		result, err := artistManager.Unlike(context.Background(), in)
		assert.Nil(t, err)
		assert.Equal(t, &google_proto.Empty{}, result)
	})
}
