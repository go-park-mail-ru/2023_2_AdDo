package artist_usecase

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"main/internal/pkg/album"
	"main/internal/pkg/artist"
	"main/internal/pkg/track"
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

	useCase := Default{
		repoTrack:  mockTrackRepo,
		repoArtist: mockArtistRepo,
		repoAlbum:  mockAlbumRepo,
	}

	art := artist.Base{
		Id:     1,
		Name:   "Artist",
		Avatar: "Avatar",
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

	mockArtistRepo.EXPECT().Get(art.Id).Return(art, nil)
	mockAlbumRepo.EXPECT().GetByArtistId(art.Id).Return(albums, nil)
	mockTrackRepo.EXPECT().GetByArtist(art.Id).Return(tracks, nil)

	expectedArtist := artist.Response{
		Id:     1,
		Name:   "Artist",
		Avatar: "Avatar",
		Albums: albums,
		Tracks: tracks,
	}

	receivedArtist, err := useCase.GetArtistInfo(art.Id)

	assert.Equal(t, nil, err)
	assert.Equal(t, expectedArtist, receivedArtist)
}
