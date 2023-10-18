package album_usecase

import (
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"main/internal/pkg/album"
	"main/internal/pkg/artist"
	"main/internal/pkg/track"
	album_mock "main/test/mocks/album"
	artist_mock "main/test/mocks/artist"
	track_mock "main/test/mocks/track"
	"testing"
)

func Test_formResponse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTrackRepo := track_mock.NewMockRepository(ctrl)
	mockArtistRepo := artist_mock.NewMockRepository(ctrl)
	mockAlbumRepo := album_mock.NewMockRepository(ctrl)

	useCase := Default{
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

	expectedAlbums := []album.Response{
		{
			Id:         1,
			Name:       "AlbumName1",
			Preview:    "Preview1",
			ArtistId:   1,
			ArtistName: "Artist1",
		},
		{
			Id:         2,
			Name:       "AlbumName2",
			Preview:    "Preview2",
			ArtistId:   2,
			ArtistName: "Artist2",
		},
	}

	receivedAlbums, err := useCase.formResponse(albums)

	assert.Equal(t, nil, err)
	assert.Equal(t, expectedAlbums, receivedAlbums)
}

func Test_GetAlbum(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTrackRepo := track_mock.NewMockRepository(ctrl)
	mockArtistRepo := artist_mock.NewMockRepository(ctrl)
	mockAlbumRepo := album_mock.NewMockRepository(ctrl)

	useCase := Default{
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
			Id:      1,
			Name:    "Track1",
			Preview: "Preview1",
			Content: "Content1",
		},
		{
			Id:      2,
			Name:    "Track2",
			Preview: "Preview2",
			Content: "Content2",
		},
	}

	mockAlbumRepo.EXPECT().Get(alb.Id).Return(alb, nil)
	mockArtistRepo.EXPECT().GetByAlbumId(alb.Id).Return(art, nil)
	mockTrackRepo.EXPECT().GetByAlbum(alb.Id).Return(tracks, nil)

	expectedAlbum := album.Response{
		Id:         1,
		Name:       "AlbumName",
		Preview:    "Preview",
		ArtistId:   1,
		ArtistName: "Artist",
		Tracks:     tracks,
	}

	receivedAlbum, err := useCase.GetAlbum(alb.Id)

	assert.Equal(t, nil, err)
	assert.Equal(t, expectedAlbum, receivedAlbum)
}
