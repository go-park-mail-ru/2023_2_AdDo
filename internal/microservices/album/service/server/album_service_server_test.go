package grpc_album_server

import (
	"context"
	"github.com/golang/mock/gomock"
	google_proto "github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"main/internal/microservices/album/proto"
	track_proto "main/internal/microservices/track/proto"
	"main/internal/pkg/album"
	"main/internal/pkg/artist"
	"main/internal/pkg/track"
	album_mock "main/test/mocks/album"
	artist_mock "main/test/mocks/artist"
	track_mock "main/test/mocks/track"
	"testing"
)

func Test_GetAlbum(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTrackRepo := track_mock.NewMockRepository(ctrl)
	mockArtistRepo := artist_mock.NewMockRepository(ctrl)
	mockAlbumRepo := album_mock.NewMockRepository(ctrl)

	albumManager := AlbumManager{
		repoTrack:  mockTrackRepo,
		repoArtist: mockArtistRepo,
		repoAlbum:  mockAlbumRepo,
		logger:     logrus.New(),
	}

	alb := album.Base{
		Id:      1,
		Name:    "AlbumName",
		Preview: "Preview",
	}

	art := artist.Base{
		Id:   1,
		Name: "Artist",
	}

	tracks := []track.Response{
		{
			Id:         1,
			Name:       "Track1",
			Preview:    "Preview1",
			Content:    "Content1",
			ArtistName: "Artist1",
			Duration:   100,
			IsLiked:    true,
		},
		{
			Id:         2,
			Name:       "Track2",
			Preview:    "Preview2",
			Content:    "Content2",
			ArtistName: "Artist2",
			Duration:   200,
			IsLiked:    false,
		},
	}

	const albumId uint64 = 1
	in := &proto.AlbumId{AlbumId: albumId}

	mockAlbumRepo.EXPECT().Get(in.GetAlbumId()).Return(alb, nil)
	mockArtistRepo.EXPECT().GetByAlbumId(in.GetAlbumId()).Return(art, nil)
	mockTrackRepo.EXPECT().GetByAlbum(in.GetAlbumId()).Return(tracks, nil)

	expectedSerializedAlbum := &proto.AlbumResponse{
		Id:         1,
		Name:       "AlbumName",
		Preview:    "Preview",
		ArtistId:   1,
		ArtistName: "Artist",
		Tracks: &track_proto.TracksResponse{
			Tracks: []*track_proto.Track{
				{
					Id:         1,
					Name:       "Track1",
					Preview:    "Preview1",
					Content:    "Content1",
					ArtistName: "Artist1",
					Duration:   100,
					IsLiked:    true,
				},
				{
					Id:         2,
					Name:       "Track2",
					Preview:    "Preview2",
					Content:    "Content2",
					ArtistName: "Artist2",
					Duration:   200,
					IsLiked:    false,
				},
			}},
	}

	receivedAlbum, err := albumManager.GetAlbum(context.Background(), in)

	assert.Equal(t, nil, err)
	assert.Equal(t, expectedSerializedAlbum, receivedAlbum)
}

func Test_formResponse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTrackRepo := track_mock.NewMockRepository(ctrl)
	mockArtistRepo := artist_mock.NewMockRepository(ctrl)
	mockAlbumRepo := album_mock.NewMockRepository(ctrl)

	albumManager := AlbumManager{
		repoTrack:  mockTrackRepo,
		repoArtist: mockArtistRepo,
		repoAlbum:  mockAlbumRepo,
		logger:     logrus.New(),
	}

	albums := []album.Base{
		{
			Id:      1,
			Name:    "AlbumName1",
			Preview: "Preview1",
		},
		{
			Id:      2,
			Name:    "AlbumName2",
			Preview: "Preview2",
		},
	}

	artist1 := artist.Base{
		Id:   1,
		Name: "Artist1",
	}

	artist2 := artist.Base{
		Id:   2,
		Name: "Artist2",
	}

	mockArtistRepo.EXPECT().GetByAlbumId(albums[0].Id).Return(artist1, nil)
	mockArtistRepo.EXPECT().GetByAlbumId(albums[1].Id).Return(artist2, nil)

	expectedSerializedAlbums := &proto.AlbumsResponse{
		Albums: []*proto.AlbumResponse{
			{
				Id:         1,
				Name:       "AlbumName1",
				Preview:    "Preview1",
				ArtistId:   1,
				ArtistName: "Artist1",
				Tracks: &track_proto.TracksResponse{
					Tracks: []*track_proto.Track{},
				},
			},
			{
				Id:         2,
				Name:       "AlbumName2",
				Preview:    "Preview2",
				ArtistId:   2,
				ArtistName: "Artist2",
				Tracks: &track_proto.TracksResponse{
					Tracks: []*track_proto.Track{},
				},
			},
		},
	}

	receivedAlbums, err := albumManager.formResponse(albums)

	assert.Equal(t, nil, err)
	assert.Equal(t, expectedSerializedAlbums, receivedAlbums)
}

func Test_Like(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTrackRepo := track_mock.NewMockRepository(ctrl)
	mockArtistRepo := artist_mock.NewMockRepository(ctrl)
	mockAlbumRepo := album_mock.NewMockRepository(ctrl)

	albumManager := AlbumManager{
		repoTrack:  mockTrackRepo,
		repoArtist: mockArtistRepo,
		repoAlbum:  mockAlbumRepo,
		logger:     logrus.New(),
	}

	in := &proto.AlbumToUserId{UserId: "user", AlbumId: 2}

	t.Run("Like", func(t *testing.T) {
		mockAlbumRepo.EXPECT().CreateLike(in.GetUserId(), in.GetAlbumId()).Return(nil)

		result, err := albumManager.Like(context.Background(), in)
		assert.Nil(t, err)
		assert.Equal(t, &google_proto.Empty{}, result)
	})

	t.Run("IsLike", func(t *testing.T) {
		mockAlbumRepo.EXPECT().CheckLike(in.GetUserId(), in.GetAlbumId()).Return(true, nil)

		result, err := albumManager.IsLike(context.Background(), in)
		assert.Nil(t, err)
		assert.Equal(t, &proto.IsLikedAlbum{IsLiked: true}, result)
	})

	t.Run("UnLike", func(t *testing.T) {
		mockAlbumRepo.EXPECT().DeleteLike(in.GetUserId(), in.GetAlbumId()).Return(nil)

		result, err := albumManager.Unlike(context.Background(), in)
		assert.Nil(t, err)
		assert.Equal(t, &google_proto.Empty{}, result)
	})
}
