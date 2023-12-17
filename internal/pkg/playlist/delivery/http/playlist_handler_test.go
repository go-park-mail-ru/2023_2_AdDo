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
	"strconv"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestActionsOnPlaylist(t *testing.T) {
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
		sessionId           = "sessionID"
		userId              = "4dfe819e-3112-4e3d-900e-6ea9db5c41fe"
		playlistId   uint64 = 5
		playlistName        = "Playlist"
		trackId      uint64 = 10
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
			AuthorId: userId,
		}

		playlistResponse := playlist.Response{
			Id:       playlistId,
			Name:     "Новый плейлист",
			AuthorId: userId,
			Preview:  "",
		}

		req := httptest.NewRequest(http.MethodPost, "/playlist", nil)
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockSessionUseCase.EXPECT().GetUserId(sessionId).Return(userId, nil)
		mockPlaylistUseCase.EXPECT().Create(playlistBase).Return(playlistResponse, nil)

		err := handler.Create(w, req)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, w.Code)

		var receivedPlaylist playlist.Response
		err = json.NewDecoder(w.Body).Decode(&receivedPlaylist)
		assert.Nil(t, err)
		assert.Equal(t, receivedPlaylist, playlistResponse)
	})

	t.Run("Get", func(t *testing.T) {
		playlistResponse := playlist.Response{
			Id:       playlistId,
			Name:     "Playlist",
			AuthorId: userId,
			Preview:  "preview",

			Tracks:   make([]track.Response, 0),
		}

		req := httptest.NewRequest(http.MethodGet, "/playlist", nil)
		req = mux.SetURLVars(req, map[string]string{"id": strconv.FormatUint(playlistId, 10)})
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()


		mockSessionUseCase.EXPECT().CheckSession(sessionId).Return(true, nil)
		mockPlaylistUseCase.EXPECT().Get(playlistId).Return(playlistResponse, nil)

		err := handler.Get(w, req)
		assert.Nil(t, err)

		var received playlist.Response
		err = json.NewDecoder(w.Body).Decode(&received)
		assert.Nil(t, err)
		assert.Equal(t, playlistResponse, received)
	})

	requestBody, err := json.Marshal(track.Id{Id: trackId})
	assert.Nil(t, err)

	t.Run("AddTrack", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/playlist/add_track", bytes.NewBuffer(requestBody))
		req = mux.SetURLVars(req, map[string]string{"id": strconv.FormatUint(playlistId, 10)})
		w := httptest.NewRecorder()

		mockPlaylistUseCase.EXPECT().AddTrack(playlistId, trackId).Return(nil)

		err = handler.AddTrack(w, req)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})

	t.Run("RemoveTrack", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/playlist/remove_track", bytes.NewBuffer(requestBody))
		req = mux.SetURLVars(req, map[string]string{"id": strconv.FormatUint(playlistId, 10)})
		w := httptest.NewRecorder()

		mockPlaylistUseCase.EXPECT().RemoveTrack(playlistId, trackId).Return(nil)

		err = handler.RemoveTrack(w, req)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})

	t.Run("UpdateName", func(t *testing.T) {
		requestBody, err = json.Marshal(playlist.Name{Name: playlistName})
		assert.Nil(t, err)

		req := httptest.NewRequest(http.MethodPost, "/playlist/update_name", bytes.NewBuffer(requestBody))
		req = mux.SetURLVars(req, map[string]string{"id": strconv.FormatUint(playlistId, 10)})
		w := httptest.NewRecorder()

		mockPlaylistUseCase.EXPECT().UpdateName(playlistId, playlistName).Return(nil)

		err = handler.UpdateName(w, req)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})

	t.Run("Delete", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/playlist", nil)
		req = mux.SetURLVars(req, map[string]string{"id": strconv.FormatUint(playlistId, 10)})
		w := httptest.NewRecorder()

		mockPlaylistUseCase.EXPECT().DeleteById(playlistId).Return(nil)

		err = handler.Delete(w, req)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})

	t.Run("MakePublic", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/playlist/make_public", nil)
		req = mux.SetURLVars(req, map[string]string{"id": strconv.FormatUint(playlistId, 10)})
		w := httptest.NewRecorder()

		mockPlaylistUseCase.EXPECT().MakePublic(playlistId).Return(nil)

		err = handler.MakePublic(w, req)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})

	t.Run("MakePrivate", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/playlist/make_private", nil)
		req = mux.SetURLVars(req, map[string]string{"id": strconv.FormatUint(playlistId, 10)})
		w := httptest.NewRecorder()

		mockPlaylistUseCase.EXPECT().MakePrivate(playlistId).Return(nil)

		err = handler.MakePrivate(w, req)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNoContent, w.Code)
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
		playlistId uint64 = 1
		sessionId         = "sessionID"
		userId            = "4dfe819e-3112-4e3d-900e-6ea9db5c41fe"
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
		req := httptest.NewRequest(http.MethodPost, "/playlist/like", nil)
		w := httptest.NewRecorder()
		err := handler.Like(w, req)
		assert.Equal(t, http.StatusBadRequest, err.(common_handler.StatusError).Code)
	})

	t.Run("Like GetCookieError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/playlist/like", nil)
		req = mux.SetURLVars(req, map[string]string{"id": strconv.FormatUint(playlistId, 10)})
		w := httptest.NewRecorder()

		err := handler.Like(w, req)
		assert.Equal(t, http.StatusUnauthorized, err.(common_handler.StatusError).Code)
	})

	t.Run("Like Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/playlist/like", nil)
		req = mux.SetURLVars(req, map[string]string{"id": strconv.FormatUint(playlistId, 10)})
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockSessionUseCase.EXPECT().GetUserId(sessionId).Return(userId, nil)
		mockPlaylistUseCase.EXPECT().Like(userId, playlistId).Return(nil)

		err := handler.Like(w, req)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})

	t.Run("IsLike GetUserIdByCookieError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/playlist/is_like", nil)
		req = mux.SetURLVars(req, map[string]string{"id": strconv.FormatUint(playlistId, 10)})
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockSessionUseCase.EXPECT().GetUserId(sessionId).Return("", errors.New("error while getting user"))

		err := handler.IsLike(w, req)
		assert.Equal(t, http.StatusUnauthorized, err.(common_handler.StatusError).Code)
	})

	t.Run("IsLike LikeCheckError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/playlist/is_like", nil)
		req = mux.SetURLVars(req, map[string]string{"id": strconv.FormatUint(playlistId, 10)})
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockSessionUseCase.EXPECT().GetUserId(sessionId).Return(userId, nil)
		mockPlaylistUseCase.EXPECT().IsLike(userId, playlistId).Return(false, errors.New("error while checking like"))

		err := handler.IsLike(w, req)
		assert.Equal(t, http.StatusNotFound, err.(common_handler.StatusError).Code)
	})

	t.Run("IsLike Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/playlist/is_like", nil)
		req = mux.SetURLVars(req, map[string]string{"id": strconv.FormatUint(playlistId, 10)})
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockSessionUseCase.EXPECT().GetUserId(sessionId).Return(userId, nil)
		mockPlaylistUseCase.EXPECT().IsLike(userId, playlistId).Return(isLiked, nil)

		err := handler.IsLike(w, req)
		assert.Nil(t, err)

		var received response.IsLiked
		err = json.NewDecoder(w.Body).Decode(&received)
		assert.Nil(t, err)
		assert.Equal(t, response.IsLiked{IsLiked: isLiked}, received)
	})

	t.Run("Unlike Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/playlist/unlike", nil)
		req = mux.SetURLVars(req, map[string]string{"id": strconv.FormatUint(playlistId, 10)})
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockSessionUseCase.EXPECT().GetUserId(sessionId).Return(userId, nil)
		mockPlaylistUseCase.EXPECT().Unlike(userId, playlistId).Return(nil)

		err := handler.Unlike(w, req)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})
}
