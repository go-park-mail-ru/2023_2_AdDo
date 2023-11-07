package playlist_delivery

import (
	"bytes"
	"encoding/json"
	"errors"
	common_handler "main/internal/common/handler"
	"main/internal/common/response"
	"main/internal/pkg/playlist"
	"main/internal/pkg/session"
	"main/internal/pkg/track"
	playlist_mock "main/test/mocks/playlist"
	session_mock "main/test/mocks/session"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestCreateAndGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPlaylistUseCase := playlist_mock.NewMockUseCase(ctrl)
	mockSessionUseCase := session_mock.NewMockUseCase(ctrl)

	handler := &Handler{
		playlistUseCase: mockPlaylistUseCase,
		sessionUseCase:  mockSessionUseCase,
		logger:          logrus.New(),
	}

	const (
		sessionId            = "sessionID"
		userId               = "qwer-qwer-qwer"
		playlistId           = "100"
		playlistIdInt uint64 = 100
	)

	cookie := http.Cookie{
		Name:     session.CookieName,
		Value:    sessionId,
		Expires:  time.Now().Add(session.TimeToLiveCookie),
		Secure:   true,
		HttpOnly: true,
	}

	t.Run("Create", func(t *testing.T) {
		playlistBase := playlist.Base{
			// Id:       playlistIdInt,
			// Name:     "Playlist",
			AuthorId: userId,
			// Preview:  "preview",
		}
		// playlistResponse := playlist.Response{
		// 	Id:       playlistIdInt,
		// 	Name:     "Playlist",
		// 	AuthorId: userId,
		// 	Preview:  "preview",
		// }
		playlistResponse := playlist.Response{}
		requestBody, err := json.Marshal(playlistBase)
		assert.Nil(t, err)

		req := httptest.NewRequest(http.MethodPost, "/playlist", bytes.NewBuffer(requestBody))
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockSessionUseCase.EXPECT().GetUserId(sessionId).Return(userId, nil)
		mockPlaylistUseCase.EXPECT().Create(playlistBase).Return(playlistResponse, nil)

		err = handler.Create(w, req)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Get", func(t *testing.T) {
		playlistResponse := playlist.Response{
			Id:       playlistIdInt,
			Name:     "Playlist",
			AuthorId: userId,
			Preview:  "preview",
			IsYours:  true,
			Tracks:   make([]track.Response, 0),
		}

		req := httptest.NewRequest(http.MethodPost, "/playlist", nil)
		req = mux.SetURLVars(req, map[string]string{"id": playlistId})
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockSessionUseCase.EXPECT().GetUserId(sessionId).Return(userId, nil)
		mockPlaylistUseCase.EXPECT().Get(userId, playlistIdInt).Return(playlistResponse, nil)

		err := handler.Get(w, req)
		assert.Nil(t, err)

		var received playlist.Response
		err = json.NewDecoder(w.Body).Decode(&received)
		assert.Nil(t, err)
		assert.Equal(t, playlistResponse, received)
	})
}

func TestLike(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPlaylistUseCase := playlist_mock.NewMockUseCase(ctrl)
	mockSessionUseCase := session_mock.NewMockUseCase(ctrl)

	handler := &Handler{
		playlistUseCase: mockPlaylistUseCase,
		sessionUseCase:  mockSessionUseCase,
		logger:          logrus.New(),
	}

	const (
		playlistId           = "1"
		playlistIdInt uint64 = 1
		sessionId            = "sessionID"
		userId               = "qwer-qwer-qwer"
		isLiked              = true
	)

	cookie := http.Cookie{
		Name:     session.CookieName,
		Value:    sessionId,
		Expires:  time.Now().Add(session.TimeToLiveCookie),
		Secure:   true,
		HttpOnly: true,
	}

	t.Run("Like GetPathParamError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/playlist/like", nil)
		w := httptest.NewRecorder()
		err := handler.Like(w, req)
		assert.Equal(t, http.StatusBadRequest, err.(common_handler.StatusError).Code)
	})

	t.Run("Like GetCookieError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/playlist/like", nil)
		req = mux.SetURLVars(req, map[string]string{"id": playlistId})
		w := httptest.NewRecorder()

		err := handler.Like(w, req)
		assert.Equal(t, http.StatusUnauthorized, err.(common_handler.StatusError).Code)
	})

	t.Run("Like Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/playlist/like", nil)
		req = mux.SetURLVars(req, map[string]string{"id": playlistId})
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockSessionUseCase.EXPECT().GetUserId(sessionId).Return(userId, nil)
		mockPlaylistUseCase.EXPECT().Like(userId, playlistIdInt).Return(nil)

		err := handler.Like(w, req)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})

	t.Run("IsLike GetUserIdByCookieError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/playlist/is_like", nil)
		req = mux.SetURLVars(req, map[string]string{"id": playlistId})
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockSessionUseCase.EXPECT().GetUserId(sessionId).Return("", errors.New("error while getting user"))

		err := handler.IsLike(w, req)
		assert.Equal(t, http.StatusUnauthorized, err.(common_handler.StatusError).Code)
	})

	t.Run("IsLike LikeCheckError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/playlist/is_like", nil)
		req = mux.SetURLVars(req, map[string]string{"id": playlistId})
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockSessionUseCase.EXPECT().GetUserId(sessionId).Return(userId, nil)
		mockPlaylistUseCase.EXPECT().IsLike(userId, playlistIdInt).Return(false, errors.New("error while checking like"))

		err := handler.IsLike(w, req)
		assert.Equal(t, http.StatusNotFound, err.(common_handler.StatusError).Code)
	})

	t.Run("IsLike Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/playlist/is_like", nil)
		req = mux.SetURLVars(req, map[string]string{"id": playlistId})
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockSessionUseCase.EXPECT().GetUserId(sessionId).Return(userId, nil)
		mockPlaylistUseCase.EXPECT().IsLike(userId, playlistIdInt).Return(isLiked, nil)

		err := handler.IsLike(w, req)
		assert.Nil(t, err)

		var received response.IsLiked
		err = json.NewDecoder(w.Body).Decode(&received)
		assert.Nil(t, err)
		assert.Equal(t, response.IsLiked{IsLiked: isLiked}, received)
	})

	t.Run("Unlike Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/playlist/unlike", nil)
		req = mux.SetURLVars(req, map[string]string{"id": playlistId})
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockSessionUseCase.EXPECT().GetUserId(sessionId).Return(userId, nil)
		mockPlaylistUseCase.EXPECT().Unlike(userId, playlistIdInt).Return(nil)

		err := handler.Unlike(w, req)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})
}
