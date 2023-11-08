package track_usecase

import (
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	album_mock "main/test/mocks/album"
	artist_mock "main/test/mocks/artist"
	track_mock "main/test/mocks/track"
	"testing"
)

func Test_UseCase(t *testing.T) {
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

	t.Run("Listen", func(t *testing.T) {
		const trackId uint64 = 1
		mockTrackRepo.EXPECT().AddListen(trackId).Return(nil)

		err := useCase.Listen(trackId)
		assert.Nil(t, err)
	})

	t.Run("Like", func(t *testing.T) {
		const userId = "1"
		const trackId uint64 = 2
		mockTrackRepo.EXPECT().CreateLike(userId, trackId).Return(nil)

		err := useCase.Like(userId, trackId)
		assert.Nil(t, err)
	})
}
