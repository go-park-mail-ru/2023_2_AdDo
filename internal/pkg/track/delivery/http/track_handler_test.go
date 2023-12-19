package track_delivery

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	common_handler "main/internal/common/handler"
	"main/internal/common/response"
	"main/internal/pkg/session"
	"main/internal/pkg/track"
	session_mock "main/test/mocks/session"
	track_mock "main/test/mocks/track"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

func TestListen(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTrackUseCase := track_mock.NewMockUseCase(ctrl)
	mockSessionUseCase := session_mock.NewMockUseCase(ctrl)

	handler := &TrackHandler{
		trackUseCase:   mockTrackUseCase,
		sessionUseCase: mockSessionUseCase,
		logger:         logrus.New(),
	}

	const (
		trackId   uint64 = 1
		sessionId        = "sessionID"
		userId           = "4e94ca97-f6b2-4fa4-9509-16fc645e2e37"
		duration  int    = 40
	)

	cookie := http.Cookie{
		Name:     session.CookieName,
		Value:    sessionId,
		Expires:  time.Now().Add(session.TimeToLiveCookie),
		Secure:   true,
		HttpOnly: true,
	}

	t.Run("BadRequestError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/listen", nil)
		w := httptest.NewRecorder()

		err := handler.Listen(w, req)

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusBadRequest, err.(common_handler.StatusError).Code)
	})

	t.Run("UnauthorizedError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/listen", nil)
		req = mux.SetURLVars(req, map[string]string{"id": strconv.FormatUint(trackId, 10)})
		w := httptest.NewRecorder()

		err := handler.Listen(w, req)

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusUnauthorized, err.(common_handler.StatusError).Code)
	})

	t.Run("Success", func(t *testing.T) {
		requestBody, err := json.Marshal(Duration{Duration: duration})

		req := httptest.NewRequest(http.MethodPost, "/listen", bytes.NewBuffer(requestBody))
		req = mux.SetURLVars(req, map[string]string{"id": strconv.FormatUint(trackId, 10)})
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockSessionUseCase.EXPECT().GetUserId(sessionId).Return(userId, nil)
		mockTrackUseCase.EXPECT().Listen(userId, trackId, uint32(duration)).Return(nil)
		err = handler.Listen(w, req)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})
}

func TestLike(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTrackUseCase := track_mock.NewMockUseCase(ctrl)
	mockSessionUseCase := session_mock.NewMockUseCase(ctrl)

	handler := &TrackHandler{
		trackUseCase:   mockTrackUseCase,
		sessionUseCase: mockSessionUseCase,
		logger:         logrus.New(),
	}

	const (
		trackId   uint64 = 1
		sessionId        = "sessionID"
		userId           = "4e94ca97-f6b2-4fa4-9509-16fc645e2e37"
		isLiked          = true
	)

	cookie := http.Cookie{
		Name:     session.CookieName,
		Value:    sessionId,
		Expires:  time.Now().Add(session.TimeToLiveCookie),
		Secure:   true,
		HttpOnly: true,
	}

	t.Run("Like GetPathParamError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/track/like", nil)
		w := httptest.NewRecorder()
		err := handler.Like(w, req)
		assert.Equal(t, http.StatusBadRequest, err.(common_handler.StatusError).Code)
	})

	t.Run("Like GetCookieError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/track/like", nil)
		req = mux.SetURLVars(req, map[string]string{"id": strconv.FormatUint(trackId, 10)})
		w := httptest.NewRecorder()

		err := handler.Like(w, req)
		assert.Equal(t, http.StatusUnauthorized, err.(common_handler.StatusError).Code)
	})

	t.Run("Like Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/track/like", nil)
		req = mux.SetURLVars(req, map[string]string{"id": strconv.FormatUint(trackId, 10)})
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockSessionUseCase.EXPECT().GetUserId(sessionId).Return(userId, nil)
		mockTrackUseCase.EXPECT().Like(userId, trackId).Return(nil)

		err := handler.Like(w, req)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})

	t.Run("IsLike GetUserIdByCookieError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/track/is_like", nil)
		req = mux.SetURLVars(req, map[string]string{"id": strconv.FormatUint(trackId, 10)})
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockSessionUseCase.EXPECT().GetUserId(sessionId).Return("", errors.New("error while getting user"))

		err := handler.IsLike(w, req)
		assert.Equal(t, http.StatusUnauthorized, err.(common_handler.StatusError).Code)
	})

	t.Run("IsLike LikeCheckError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/track/is_like", nil)
		req = mux.SetURLVars(req, map[string]string{"id": strconv.FormatUint(trackId, 10)})
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockSessionUseCase.EXPECT().GetUserId(sessionId).Return(userId, nil)
		mockTrackUseCase.EXPECT().IsLike(userId, trackId).Return(false, errors.New("error while checking like"))

		err := handler.IsLike(w, req)
		assert.Equal(t, http.StatusInternalServerError, err.(common_handler.StatusError).Code)
	})

	t.Run("IsLike Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/track/is_like", nil)
		req = mux.SetURLVars(req, map[string]string{"id": strconv.FormatUint(trackId, 10)})
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockSessionUseCase.EXPECT().GetUserId(sessionId).Return(userId, nil)
		mockTrackUseCase.EXPECT().IsLike(userId, trackId).Return(isLiked, nil)

		err := handler.IsLike(w, req)
		assert.Nil(t, err)

		var received response.IsLiked
		err = json.NewDecoder(w.Body).Decode(&received)
		assert.Nil(t, err)
		assert.Equal(t, response.IsLiked{IsLiked: isLiked}, received)
	})

	t.Run("Unlike Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/track/unlike", nil)
		req = mux.SetURLVars(req, map[string]string{"id": strconv.FormatUint(trackId, 10)})
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockSessionUseCase.EXPECT().GetUserId(sessionId).Return(userId, nil)
		mockTrackUseCase.EXPECT().Unlike(userId, trackId).Return(nil)

		err := handler.Unlike(w, req)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})
}

func TestGetUserTracks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTrackUseCase := track_mock.NewMockUseCase(ctrl)
	mockSessionUseCase := session_mock.NewMockUseCase(ctrl)

	handler := &TrackHandler{
		trackUseCase:   mockTrackUseCase,
		sessionUseCase: mockSessionUseCase,
		logger:         logrus.New(),
	}

	const (
		sessionId = "sessionID"
		userId    = "qwer-qwer-qwer"
	)

	cookie := http.Cookie{
		Name:     session.CookieName,
		Value:    sessionId,
		Expires:  time.Now().Add(session.TimeToLiveCookie),
		Secure:   true,
		HttpOnly: true,
	}

	tracks := []track.Response{
		{
			Id:         1,
			Name:       "Track1",
			Preview:    "Preview1",
			Content:    "Content1",
			ArtistId:   1,
			ArtistName: "Artist1",
			Duration:   "1:20",
			IsLiked:    true,
		},
		{
			Id:         2,
			Name:       "Track2",
			Preview:    "Preview2",
			Content:    "Content2",
			ArtistId:   2,
			ArtistName: "Artist2",
			Duration:   "1:20",
			IsLiked:    true,
		},
	}

	expectedTracks := track.LikedTracks{
		Tracks: tracks,
	}

	t.Run("GetUserTracks Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/collection/tracks", nil)
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockSessionUseCase.EXPECT().GetUserId(sessionId).Return(userId, nil)
		mockTrackUseCase.EXPECT().GetUserLikedTracks(userId).Return(tracks, nil)

		err := handler.GetUserTracks(w, req)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, w.Code)

		var receivedTracks track.LikedTracks
		err = json.NewDecoder(w.Body).Decode(&receivedTracks)
		assert.Nil(t, err)
		assert.Equal(t, expectedTracks, receivedTracks)
	})
}
