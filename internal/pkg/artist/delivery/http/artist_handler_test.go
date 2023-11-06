package artist_delivery

import (
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	common_handler "main/internal/common/handler"
	"main/internal/common/response"
	"main/internal/pkg/album"
	"main/internal/pkg/artist"
	"main/internal/pkg/session"
	"main/internal/pkg/track"
	artist_mock "main/test/mocks/artist"
	session_mock "main/test/mocks/session"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
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

func TestLike(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockArtistUseCase := artist_mock.NewMockUseCase(ctrl)
	mockSessionUseCase := session_mock.NewMockUseCase(ctrl)

	handler := &ArtistHandler{
		ArtistUseCase:  mockArtistUseCase,
		SessionUseCase: mockSessionUseCase,
		logger:         logrus.New(),
	}

	const (
		artistId           = "1"
		artistIdInt uint64 = 1
		sessionId          = "sessionID"
		userId             = "qwer-qwer-qwer"
		isLiked            = true
	)

	cookie := http.Cookie{
		Name:     session.CookieName,
		Value:    sessionId,
		Expires:  time.Now().Add(session.TimeToLiveCookie),
		Secure:   true,
		HttpOnly: true,
	}

	t.Run("Like GetPathParamError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/artist/like", nil)
		w := httptest.NewRecorder()

		err := handler.Like(w, req)
		assert.Equal(t, http.StatusBadRequest, err.(common_handler.StatusError).Code)
	})

	t.Run("Like GetCookieError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/artist/like", nil)
		req = mux.SetURLVars(req, map[string]string{"id": artistId})
		w := httptest.NewRecorder()

		err := handler.Like(w, req)
		assert.Equal(t, http.StatusUnauthorized, err.(common_handler.StatusError).Code)
	})

	t.Run("Like Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/artist/like", nil)
		req = mux.SetURLVars(req, map[string]string{"id": artistId})
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockSessionUseCase.EXPECT().GetUserId(sessionId).Return(userId, nil)
		mockArtistUseCase.EXPECT().Like(userId, artistIdInt).Return(nil)

		err := handler.Like(w, req)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})

	t.Run("IsLike GetUserIdByCookieError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/artist/is_like", nil)
		req = mux.SetURLVars(req, map[string]string{"id": artistId})
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockSessionUseCase.EXPECT().GetUserId(sessionId).Return("", errors.New("error while getting user"))

		err := handler.IsLike(w, req)
		assert.Equal(t, http.StatusUnauthorized, err.(common_handler.StatusError).Code)
	})

	t.Run("IsLike LikeCheckError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/artist/is_like", nil)
		req = mux.SetURLVars(req, map[string]string{"id": artistId})
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockSessionUseCase.EXPECT().GetUserId(sessionId).Return(userId, nil)
		mockArtistUseCase.EXPECT().IsLike(userId, artistIdInt).Return(false, errors.New("error while checking like"))

		err := handler.IsLike(w, req)
		assert.Equal(t, http.StatusNotFound, err.(common_handler.StatusError).Code)
	})

	t.Run("IsLike Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/artist/is_like", nil)
		req = mux.SetURLVars(req, map[string]string{"id": artistId})
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockSessionUseCase.EXPECT().GetUserId(sessionId).Return(userId, nil)
		mockArtistUseCase.EXPECT().IsLike(userId, artistIdInt).Return(isLiked, nil)

		err := handler.IsLike(w, req)
		assert.Nil(t, err)

		var received response.IsLiked
		err = json.NewDecoder(w.Body).Decode(&received)
		assert.Nil(t, err)
		assert.Equal(t, response.IsLiked{IsLiked: isLiked}, received)
	})

	t.Run("Unlike Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/artist/unlike", nil)
		req = mux.SetURLVars(req, map[string]string{"id": artistId})
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockSessionUseCase.EXPECT().GetUserId(sessionId).Return(userId, nil)
		mockArtistUseCase.EXPECT().Unlike(userId, artistIdInt).Return(nil)

		err := handler.Unlike(w, req)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})
}
