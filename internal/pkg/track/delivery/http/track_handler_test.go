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
	"main/internal/pkg/session"
	"main/internal/pkg/track"
	session_mock "main/test/mocks/session"
	track_mock "main/test/mocks/track"
	"net/http"
	"net/http/httptest"
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

	trackId := track.Id{Id: 999}
	requestBody, err := json.Marshal(trackId)
	assert.NoError(t, err)

	t.Run("DecodeRequestBodyError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/listen", nil)
		w := httptest.NewRecorder()

		err := handler.Listen(w, req)

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusBadRequest, err.(common_handler.StatusError).Code)
	})

	t.Run("InternalError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/listen", bytes.NewBuffer(requestBody))
		w := httptest.NewRecorder()

		mockTrackUseCase.EXPECT().Listen(trackId.Id).Return(errors.New("add listen failed"))
		err = handler.Listen(w, req)

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusInternalServerError, err.(common_handler.StatusError).Code)
	})

	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/listen", bytes.NewBuffer(requestBody))
		w := httptest.NewRecorder()

		mockTrackUseCase.EXPECT().Listen(trackId.Id).Return(nil)
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

	t.Run("GetPathParamError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/track/like", nil)
		w := httptest.NewRecorder()
		err := handler.Like(w, req)
		assert.Equal(t, http.StatusBadRequest, err.(common_handler.StatusError).Code)
	})

	t.Run("GetCookieError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/track/like", nil)
		const trackId = "1"
		req = mux.SetURLVars(req, map[string]string{"id": trackId})
		w := httptest.NewRecorder()

		err := handler.Like(w, req)
		assert.Equal(t, http.StatusUnauthorized, err.(common_handler.StatusError).Code)
	})

	t.Run("Success", func(t *testing.T) {
		const sessionId = "sessionID"
		const userId = "qwer-qwer-qwer"
		cookie := http.Cookie{
			Name:     session.CookieName,
			Value:    sessionId,
			Expires:  time.Now().Add(session.TimeToLiveCookie),
			Secure:   true,
			HttpOnly: true,
		}

		req := httptest.NewRequest(http.MethodPost, "/track/like", nil)
		const trackId = "1"
		const trackIdInt uint64 = 1
		req = mux.SetURLVars(req, map[string]string{"id": trackId})
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockSessionUseCase.EXPECT().GetUserId(sessionId).Return(userId, nil)
		mockTrackUseCase.EXPECT().Like(userId, trackIdInt).Return(nil)

		err := handler.Like(w, req)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})
}
