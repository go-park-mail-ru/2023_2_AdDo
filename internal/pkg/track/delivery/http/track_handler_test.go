package track_delivery

import (
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"main/internal/pkg/session"
	"main/internal/pkg/track"
	session_mock "main/test/mocks/session"
	track_mock "main/test/mocks/track"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

//func TestMusic(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	mockTrackUseCase := track_mock.NewMockUseCase(ctrl)
//	mockSessionUseCase := session_mock.NewMockUseCase(ctrl)
//
//	handler := &TrackHandler{
//		trackUseCase:   mockTrackUseCase,
//		sessionUseCase: mockSessionUseCase,
//	}
//
//	t.Run("Success", func(t *testing.T) {
//		expectedTracks := []track.Response{
//			{
//				Id:     1,
//				Name:   "Track 1",
//				Artist: []artist.Base{artist.Base{Name: "Artist 1"}},
//				Album:  []album.Response{album.Response{Name: "Album 1"}},
//			},
//			{
//				Id:     2,
//				Name:   "Track 2",
//				Artist: []artist.Base{artist.Base{Name: "Artist 2"}},
//				Album:  []album.Response{album.Response{Name: "Album 2"}},
//			},
//		}
//		sessionId := "sessionId"
//		mockTrackUseCase.EXPECT().GetAll().Return(expectedTracks, nil)
//		//mockSessionUseCase.EXPECT().CheckSession(sessionId, uint64(1)).Return(true, nil)
//
//		req, err := http.NewRequest(http.MethodGet, "/music?id=1", nil)
//
//		assert.NoError(t, err)
//		cookie := http.Cookie{
//			Name:    session.CookieName,
//			Value:   sessionId,
//			Path:    "",
//			Domain:  "",
//			Expires: time.Time{},
//		}
//		req.AddCookie(&cookie)
//
//		rec := httptest.NewRecorder()
//
//		err = handler.Music(rec, req)
//		assert.NoError(t, err)
//
//		assert.Equal(t, http.StatusOK, rec.Code)
//		assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))
//		receivedTracks := make([]track.Response, 0)
//		body, err := io.ReadAll(rec.Body)
//		if err != nil {
//			t.Errorf("error reading response")
//		}
//		json.Unmarshal(body, &receivedTracks)
//		assert.Equal(t, expectedTracks, receivedTracks)
//	})
//
//	t.Run("Error - TrackUseCase.GetAll", func(t *testing.T) {
//		mockTrackUseCase.EXPECT().GetAll().Return(nil, errors.New("error"))
//
//		req, err := http.NewRequest(http.MethodGet, "/music", nil)
//		assert.NoError(t, err)
//
//		rec := httptest.NewRecorder()
//
//		err = handler.Music(rec, req)
//		assert.Error(t, err)
//		assert.IsType(t, common_handler.StatusError{}, err)
//
//		statusErr, _ := err.(common_handler.StatusError)
//		assert.Equal(t, http.StatusInternalServerError, statusErr.Code)
//		assert.Equal(t, "error", statusErr.Err.Error())
//	})
//
//	// TODO: Add more test cases for different scenarios
//}

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

	t.Run("Success", func(t *testing.T) {
		trackId := track.Id{Id: 999}
		requestBody, err := json.Marshal(trackId)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/listen", bytes.NewBuffer(requestBody))
		w := httptest.NewRecorder()

		mockTrackUseCase.EXPECT().Listen(trackId.Id).Return(nil)
		err = handler.Listen(w, req)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
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

	t.Run("Success", func(t *testing.T) {
		trackId := track.Id{Id: 1}
		requestBody, err := json.Marshal(trackId)
		assert.NoError(t, err)

		const sessionId = "sessionID"
		const userId = "qwer-qwer-qwer"
		cookie := http.Cookie{
			Name:     session.CookieName,
			Value:    sessionId,
			Expires:  time.Now().Add(session.TimeToLiveCookie),
			Secure:   true,
			HttpOnly: true,
		}

		req := httptest.NewRequest(http.MethodPost, "/like", bytes.NewBuffer(requestBody))
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockSessionUseCase.EXPECT().GetUserId(sessionId).Return(userId, nil)
		mockTrackUseCase.EXPECT().Like(userId, trackId.Id).Return(nil)

		err = handler.Like(w, req)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
