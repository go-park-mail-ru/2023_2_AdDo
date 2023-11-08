package album_delivery

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
	"main/internal/pkg/session"
	"main/internal/pkg/track"
	album_mock "main/test/mocks/album"
	session_mock "main/test/mocks/session"
	track_mock "main/test/mocks/track"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestReceivingAlbums(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTrackUseCase := track_mock.NewMockUseCase(ctrl)
	mockAlbumUseCase := album_mock.NewMockUseCase(ctrl)
	mockSessionUseCase := session_mock.NewMockUseCase(ctrl)

	handler := &AlbumHandler{
		trackUseCase:   mockTrackUseCase,
		albumUseCase:   mockAlbumUseCase,
		sessionUseCase: mockSessionUseCase,
		logger:         logrus.New(),
	}

	expectedAlbums := []album.Response{
		{
			Id:         1,
			Name:       "Album 1",
			Preview:    "Preview 1",
			ArtistId:   1,
			ArtistName: "Artist 1",
			Tracks:     []track.Response{},
		},
		{
			Id:         2,
			Name:       "Album 2",
			Preview:    "Preview 2",
			ArtistId:   2,
			ArtistName: "Artist 2",
			Tracks:     []track.Response{},
		},
	}

	t.Run("Feed Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/feed", nil)
		w := httptest.NewRecorder()

		mockAlbumUseCase.EXPECT().GetRandom().Return(expectedAlbums, nil)
		err := handler.Feed(w, req)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, w.Code)

		var receivedAlbums []album.Response
		err = json.NewDecoder(w.Body).Decode(&receivedAlbums)
		assert.Nil(t, err)
		assert.Equal(t, expectedAlbums, receivedAlbums)
	})

	t.Run("Feed Fail", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/feed", nil)
		w := httptest.NewRecorder()

		mockAlbumUseCase.EXPECT().GetRandom().Return(nil, errors.New("error while getting albums"))
		err := handler.Feed(w, req)

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusInternalServerError, err.(common_handler.StatusError).Code)
	})

	t.Run("New Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/new", nil)
		w := httptest.NewRecorder()

		mockAlbumUseCase.EXPECT().GetNew().Return(expectedAlbums, nil)
		err := handler.New(w, req)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, w.Code)

		var receivedAlbums []album.Response
		err = json.NewDecoder(w.Body).Decode(&receivedAlbums)
		assert.Nil(t, err)
		assert.Equal(t, expectedAlbums, receivedAlbums)
	})

	t.Run("MostLiked Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/most_liked", nil)
		w := httptest.NewRecorder()

		mockAlbumUseCase.EXPECT().GetMostLiked().Return(expectedAlbums, nil)
		err := handler.MostLiked(w, req)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, w.Code)

		var receivedAlbums []album.Response
		err = json.NewDecoder(w.Body).Decode(&receivedAlbums)
		assert.Nil(t, err)
		assert.Equal(t, expectedAlbums, receivedAlbums)
	})

	t.Run("Popular Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/popular", nil)
		w := httptest.NewRecorder()

		mockAlbumUseCase.EXPECT().GetPopular().Return(expectedAlbums, nil)
		err := handler.Popular(w, req)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, w.Code)

		var receivedAlbums []album.Response
		err = json.NewDecoder(w.Body).Decode(&receivedAlbums)
		assert.Nil(t, err)
		assert.Equal(t, expectedAlbums, receivedAlbums)
	})
}

func TestAlbumTracks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTrackUseCase := track_mock.NewMockUseCase(ctrl)
	mockAlbumUseCase := album_mock.NewMockUseCase(ctrl)
	mockSessionUseCase := session_mock.NewMockUseCase(ctrl)

	handler := &AlbumHandler{
		trackUseCase:   mockTrackUseCase,
		albumUseCase:   mockAlbumUseCase,
		sessionUseCase: mockSessionUseCase,
		logger:         logrus.New(),
	}

	t.Run("NoPathParameter", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/album", nil)
		w := httptest.NewRecorder()

		err := handler.AlbumTracks(w, req)
		assert.NotNil(t, err)
		assert.Equal(t, http.StatusBadRequest, err.(common_handler.StatusError).Code)
	})

	t.Run("WrongPathParameter", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/album", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "***"})
		w := httptest.NewRecorder()

		err := handler.AlbumTracks(w, req)
		assert.NotNil(t, err)
		assert.Equal(t, http.StatusBadRequest, err.(common_handler.StatusError).Code)
	})

	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/album", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "1"})
		w := httptest.NewRecorder()

		expectedAlbum := album.Response{
			Id:         1,
			Name:       "Album 1",
			Preview:    "Preview 1",
			ArtistId:   1,
			ArtistName: "Artist 1",
			Tracks:     []track.Response{},
		}

		mockAlbumUseCase.EXPECT().GetAlbum(uint64(1)).Return(expectedAlbum, nil)
		err := handler.AlbumTracks(w, req)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, w.Code)

		var receivedAlbum album.Response
		err = json.NewDecoder(w.Body).Decode(&receivedAlbum)
		assert.Nil(t, err)
		assert.Equal(t, expectedAlbum, receivedAlbum)
	})
}

func TestLike(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTrackUseCase := track_mock.NewMockUseCase(ctrl)
	mockAlbumUseCase := album_mock.NewMockUseCase(ctrl)
	mockSessionUseCase := session_mock.NewMockUseCase(ctrl)

	handler := &AlbumHandler{
		trackUseCase:   mockTrackUseCase,
		albumUseCase:   mockAlbumUseCase,
		sessionUseCase: mockSessionUseCase,
		logger:         logrus.New(),
	}

	const (
		albumId           = "1"
		albumIdInt uint64 = 1
		sessionId         = "sessionID"
		userId            = "qwer-qwer-qwer"
		isLiked           = true
	)

	cookie := http.Cookie{
		Name:     session.CookieName,
		Value:    sessionId,
		Expires:  time.Now().Add(session.TimeToLiveCookie),
		Secure:   true,
		HttpOnly: true,
	}

	t.Run("Like GetPathParamError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/album/like", nil)
		w := httptest.NewRecorder()

		err := handler.Like(w, req)
		assert.Equal(t, http.StatusBadRequest, err.(common_handler.StatusError).Code)
	})

	t.Run("Like GetCookieError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/album/like", nil)
		req = mux.SetURLVars(req, map[string]string{"id": albumId})
		w := httptest.NewRecorder()

		err := handler.Like(w, req)
		assert.Equal(t, http.StatusUnauthorized, err.(common_handler.StatusError).Code)
	})

	t.Run("Like Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/album/like", nil)
		req = mux.SetURLVars(req, map[string]string{"id": albumId})
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockSessionUseCase.EXPECT().GetUserId(sessionId).Return(userId, nil)
		mockAlbumUseCase.EXPECT().Like(userId, albumIdInt).Return(nil)

		err := handler.Like(w, req)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})

	t.Run("IsLike GetPathParamError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/album/is_like", nil)
		w := httptest.NewRecorder()

		err := handler.IsLike(w, req)
		assert.Equal(t, http.StatusBadRequest, err.(common_handler.StatusError).Code)
	})

	t.Run("IsLike GetUserIdByCookieError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/album/is_like", nil)
		req = mux.SetURLVars(req, map[string]string{"id": albumId})
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockSessionUseCase.EXPECT().GetUserId(sessionId).Return("", errors.New("error while getting user"))

		err := handler.IsLike(w, req)
		assert.Equal(t, http.StatusUnauthorized, err.(common_handler.StatusError).Code)
	})

	t.Run("IsLike Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/album/is_like", nil)
		req = mux.SetURLVars(req, map[string]string{"id": albumId})
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockSessionUseCase.EXPECT().GetUserId(sessionId).Return(userId, nil)
		mockAlbumUseCase.EXPECT().IsLike(userId, albumIdInt).Return(isLiked, nil)

		err := handler.IsLike(w, req)
		assert.Nil(t, err)

		var received response.IsLiked
		err = json.NewDecoder(w.Body).Decode(&received)
		assert.Nil(t, err)
		assert.Equal(t, response.IsLiked{IsLiked: isLiked}, received)
	})

	t.Run("Unlike Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/album/unlike", nil)
		req = mux.SetURLVars(req, map[string]string{"id": albumId})
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockSessionUseCase.EXPECT().GetUserId(sessionId).Return(userId, nil)
		mockAlbumUseCase.EXPECT().Unlike(userId, albumIdInt).Return(nil)

		err := handler.Unlike(w, req)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})
}
