package artist_delivery

import (
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	common_handler "main/internal/common/handler"
	"main/internal/pkg/album"
	"main/internal/pkg/artist"
	"main/internal/pkg/track"
	artist_mock "main/test/mocks/artist"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestArtistInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockArtistUseCase := artist_mock.NewMockUseCase(ctrl)

	handler := &ArtistHandler{
		ArtistUseCase: mockArtistUseCase,
		logger:        logrus.New(),
	}

	t.Run("WrongPathParameter", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/artist", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "aaaa"})
		w := httptest.NewRecorder()

		err := handler.ArtistInfo(w, req)
		assert.NotNil(t, err)
		assert.Equal(t, http.StatusBadRequest, err.(common_handler.StatusError).Code)
	})

	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/artist", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "1"})
		w := httptest.NewRecorder()

		expectedArtist := artist.Response{
			Id:     1,
			Name:   "Artist",
			Avatar: "Avatar",
			Albums: []album.Base{
				{
					Id:      1,
					Name:    "Album",
					Preview: "Preview",
				},
			},
			Tracks: []track.Response{
				{
					Id:      1,
					Name:    "Track 1",
					Preview: "Preview 1",
					Content: "Content 1",
				},
				{
					Id:      2,
					Name:    "Track 2",
					Preview: "Preview 2",
					Content: "Content 2",
				},
			},
		}

		mockArtistUseCase.EXPECT().GetArtistInfo(uint64(1)).Return(expectedArtist, nil)
		err := handler.ArtistInfo(w, req)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, w.Code)

		var receivedArtist artist.Response
		err = json.NewDecoder(w.Body).Decode(&receivedArtist)
		assert.Nil(t, err)
		assert.Equal(t, expectedArtist, receivedArtist)
	})
}
