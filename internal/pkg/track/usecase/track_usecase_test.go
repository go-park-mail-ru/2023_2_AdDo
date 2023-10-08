package track_usecase

import (
	"errors"
	"main/internal/pkg/album"
	"main/internal/pkg/artist"
	"main/internal/pkg/track"
	album_mock "main/test/mocks/album"
	artist_mock "main/test/mocks/artist"
	track_mock "main/test/mocks/track"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTrackRepo := track_mock.NewMockRepository(ctrl)
	mockArtistRepo := artist_mock.NewMockRepository(ctrl)
	mockAlbumRepo := album_mock.NewMockRepository(ctrl)

	useCase := Default{
		repoTrack:  mockTrackRepo,
		repoArtist: mockArtistRepo,
		repoAlbum:  mockAlbumRepo,
	}

	t.Run("Success", func(t *testing.T) {
		expectedTracks := []track.Response{
			{
				Id:     1,
				Name:   "Track 1",
				Artist: []artist.Response{{Name: "Artist 1"}},
				Album:  []album.Response{{Name: "Album 1"}},
			},
			{
				Id:     2,
				Name:   "Track 2",
				Artist: []artist.Response{{Name: "Artist 2"}},
				Album:  []album.Response{{Name: "Album 2"}},
			},
		}

		mockTrackRepo.EXPECT().GetAll().Return(expectedTracks, nil)
		mockArtistRepo.EXPECT().GetByTrackId(uint64(1)).Return([]artist.Response{{Name: "Artist 1"}}, nil)
		mockAlbumRepo.EXPECT().GetByTrackId(uint64(1)).Return([]album.Response{{Name: "Album 1"}}, nil)
		mockArtistRepo.EXPECT().GetByTrackId(uint64(2)).Return([]artist.Response{{Name: "Artist 2"}}, nil)
		mockAlbumRepo.EXPECT().GetByTrackId(uint64(2)).Return([]album.Response{{Name: "Album 2"}}, nil)

		tracks, err := useCase.GetAll()

		assert.NoError(t, err)
		assert.Equal(t, expectedTracks, tracks)
	})

	t.Run("ErrorNoTracks", func(t *testing.T) {
		mockTrackRepo.EXPECT().GetAll().Return(nil, track.ErrNoTracks)

		tracks, err := useCase.GetAll()

		assert.ErrorIs(t, err, track.ErrNoTracks)
		assert.Nil(t, tracks)
	})

	t.Run("ErrorGettingArtists", func(t *testing.T) {
		mockTrackRepo.EXPECT().GetAll().Return([]track.Response{{Id: 1}}, nil)
		mockArtistRepo.EXPECT().GetByTrackId(uint64(1)).Return(nil, errors.New("error getting artists"))

		tracks, err := useCase.GetAll()

		assert.Error(t, err)
		assert.Nil(t, tracks)
	})

	t.Run("ErrorGettingAlbums", func(t *testing.T) {
		mockTrackRepo.EXPECT().GetAll().Return([]track.Response{{Id: 1}}, nil)
		mockArtistRepo.EXPECT().GetByTrackId(uint64(1)).Return([]artist.Response{artist.Response{Name: "Artist 1"}}, nil)
		mockAlbumRepo.EXPECT().GetByTrackId(uint64(1)).Return(nil, errors.New("error getting albums"))

		tracks, err := useCase.GetAll()

		assert.Error(t, err)
		assert.Nil(t, tracks)
	})
}

func TestGetPopular(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTrackRepo := track_mock.NewMockRepository(ctrl)
	mockArtistRepo := artist_mock.NewMockRepository(ctrl)
	mockAlbumRepo := album_mock.NewMockRepository(ctrl)

	useCase := Default{
		repoTrack:  mockTrackRepo,
		repoArtist: mockArtistRepo,
		repoAlbum:  mockAlbumRepo,
	}

	tracks := []track.Response{
		{
			Id:        1,
			Name:      "Track 1",
			PlayCount: 100,
		},
		{
			Id:        2,
			Name:      "Track 2",
			PlayCount: 20,
		},
		{
			Id:        3,
			Name:      "Track 3",
			PlayCount: 10,
		},
	}
	var limit uint32 = 3
	mockTrackRepo.EXPECT().GetPopular(limit).Return(tracks, nil)
	mockArtistRepo.EXPECT().GetByTrackId(uint64(1)).Return([]artist.Response{{Name: "Artist 1"}}, nil)
	mockAlbumRepo.EXPECT().GetByTrackId(uint64(1)).Return([]album.Response{{Name: "Album 1"}}, nil)
	mockArtistRepo.EXPECT().GetByTrackId(uint64(2)).Return([]artist.Response{{Name: "Artist 2"}}, nil)
	mockAlbumRepo.EXPECT().GetByTrackId(uint64(2)).Return([]album.Response{{Name: "Album 2"}}, nil)
	mockArtistRepo.EXPECT().GetByTrackId(uint64(3)).Return([]artist.Response{{Name: "Artist 3"}}, nil)
	mockAlbumRepo.EXPECT().GetByTrackId(uint64(3)).Return([]album.Response{{Name: "Album 3"}}, nil)

	expectedTracks := []track.Response{
		{
			Id:        1,
			Name:      "Track 1",
			Artist:    []artist.Response{{Name: "Artist 1"}},
			Album:     []album.Response{{Name: "Album 1"}},
			PlayCount: 100,
		},
		{
			Id:        2,
			Name:      "Track 2",
			Artist:    []artist.Response{{Name: "Artist 2"}},
			Album:     []album.Response{{Name: "Album 2"}},
			PlayCount: 20,
		},
		{
			Id:        3,
			Name:      "Track 3",
			Artist:    []artist.Response{{Name: "Artist 3"}},
			Album:     []album.Response{{Name: "Album 3"}},
			PlayCount: 10,
		},
	}

	popularTracks, err := useCase.GetPopular(limit)
	assert.Nil(t, err)
	assert.Equal(t, expectedTracks, popularTracks)
}

func TestGetByPlaylist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTrackRepo := track_mock.NewMockRepository(ctrl)
	mockArtistRepo := artist_mock.NewMockRepository(ctrl)
	mockAlbumRepo := album_mock.NewMockRepository(ctrl)

	useCase := Default{
		repoTrack:  mockTrackRepo,
		repoArtist: mockArtistRepo,
		repoAlbum:  mockAlbumRepo,
	}

	playlistId := uint64(1)
	trackIds := []uint64{1, 2}

	tracks := []track.Response{
		{
			Id:        1,
			Name:      "Track 1",
			Preview:   "Preview 1",
			Content:   "Content 1",
			PlayCount: 10,
		},
		{
			Id:        2,
			Name:      "Track 2",
			Preview:   "Preview 2",
			Content:   "Content 2",
			PlayCount: 20,
		},
	}

	mockTrackRepo.EXPECT().GetTrackIdsByPlaylist(playlistId).Return(trackIds, nil)

	mockTrackRepo.EXPECT().GetByTrackId(trackIds[0]).Return(tracks[0], nil)
	mockTrackRepo.EXPECT().GetByTrackId(trackIds[1]).Return(tracks[1], nil)

	mockArtistRepo.EXPECT().GetByTrackId(trackIds[0]).Return([]artist.Response{{Name: "Artist 1"}}, nil)
	mockArtistRepo.EXPECT().GetByTrackId(trackIds[1]).Return([]artist.Response{{Name: "Artist 2"}}, nil)
	mockAlbumRepo.EXPECT().GetByTrackId(trackIds[0]).Return([]album.Response{{Name: "Album 1"}}, nil)
	mockAlbumRepo.EXPECT().GetByTrackId(trackIds[1]).Return([]album.Response{{Name: "Album 2"}}, nil)

	expected := []track.Response{
		{
			Id:        1,
			Name:      "Track 1",
			Artist:    []artist.Response{{Name: "Artist 1"}},
			Album:     []album.Response{{Name: "Album 1"}},
			Preview:   "Preview 1",
			Content:   "Content 1",
			PlayCount: 10,
		},
		{
			Id:        2,
			Name:      "Track 2",
			Artist:    []artist.Response{{Name: "Artist 2"}},
			Album:     []album.Response{{Name: "Album 2"}},
			Preview:   "Preview 2",
			Content:   "Content 2",
			PlayCount: 20,
		},
	}

	actual, err := useCase.GetByPlaylist(playlistId)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
