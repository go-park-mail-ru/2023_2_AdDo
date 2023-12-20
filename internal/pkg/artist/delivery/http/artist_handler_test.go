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
	"strconv"
	"testing"
	"time"
)

func TestArtistInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockArtistUseCase := artist_mock.NewMockUseCase(ctrl)
	mockSessionUseCase := session_mock.NewMockUseCase(ctrl)

	handler := &ArtistHandler{
		ArtistUseCase:  mockArtistUseCase,
		SessionUseCase: mockSessionUseCase,
		logger:         logrus.New(),
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
		const artistId uint64 = 1

		req := httptest.NewRequest(http.MethodGet, "/artist", nil)
		req = mux.SetURLVars(req, map[string]string{"id": strconv.FormatUint(artistId, 10)})
		w := httptest.NewRecorder()

		expectedArtist := artist.Response{
			Id:     artistId,
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

		mockArtistUseCase.EXPECT().GetArtistInfo(artistId).Return(expectedArtist, nil)
		mockSessionUseCase.EXPECT().GetUserId("").Return("", nil)

		err := handler.ArtistInfo(w, req)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, w.Code)

		var receivedArtist artist.Response
		err = json.NewDecoder(w.Body).Decode(&receivedArtist)
		assert.Nil(t, err)
		assert.Equal(t, expectedArtist, receivedArtist)
	})
}

func TestCollectionArtist(t *testing.T) {
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
		sessionId = "sessionID"
		userId    = "aeefe1e3-0def-41f0-a3c7-91380379c58c"
	)

	cookie := http.Cookie{
		Name:     session.CookieName,
		Value:    sessionId,
		Expires:  time.Now().Add(session.TimeToLiveCookie),
		Secure:   true,
		HttpOnly: true,
	}

	artists := []artist.Base{
		{
			Id:     1,
			Name:   "Artist1",
			Avatar: "/path/to/avatar1.jpg",
		},
		{
			Id:     2,
			Name:   "Artist2",
			Avatar: "/path/to/avatar2.jpg",
		},
	}

	expectedArtists := artist.Artists{
		Artists: artists,
	}

	t.Run("CollectionArtist Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/collection/artists", nil)
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockSessionUseCase.EXPECT().GetUserId(sessionId).Return(userId, nil)
		mockArtistUseCase.EXPECT().GetUserArtists(userId).Return(expectedArtists, nil)

		err := handler.CollectionArtist(w, req)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, w.Code)

		var receivedArtists artist.Artists
		err = json.NewDecoder(w.Body).Decode(&receivedArtists)
		assert.Nil(t, err)
		assert.Equal(t, expectedArtists, receivedArtists)
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
		artistId  uint64 = 1
		sessionId        = "sessionID"
		userId           = "aeefe1e3-0def-41f0-a3c7-91380379c58c"
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
		req := httptest.NewRequest(http.MethodPost, "/artist/like", nil)
		w := httptest.NewRecorder()

		err := handler.Like(w, req)
		assert.Equal(t, http.StatusBadRequest, err.(common_handler.StatusError).Code)
	})

	t.Run("Like GetCookieError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/artist/like", nil)
		req = mux.SetURLVars(req, map[string]string{"id": strconv.FormatUint(artistId, 10)})
		w := httptest.NewRecorder()

		err := handler.Like(w, req)
		assert.Equal(t, http.StatusUnauthorized, err.(common_handler.StatusError).Code)
	})

	t.Run("Like Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/artist/like", nil)
		req = mux.SetURLVars(req, map[string]string{"id": strconv.FormatUint(artistId, 10)})
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockSessionUseCase.EXPECT().GetUserId(sessionId).Return(userId, nil)
		mockArtistUseCase.EXPECT().Like(userId, artistId).Return(nil)

		err := handler.Like(w, req)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})

	t.Run("IsLike GetUserIdByCookieError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/artist/is_like", nil)
		req = mux.SetURLVars(req, map[string]string{"id": strconv.FormatUint(artistId, 10)})
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockSessionUseCase.EXPECT().GetUserId(sessionId).Return("", errors.New("error while getting user"))

		err := handler.IsLike(w, req)
		assert.Equal(t, http.StatusUnauthorized, err.(common_handler.StatusError).Code)
	})

	t.Run("IsLike LikeCheckError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/artist/is_like", nil)
		req = mux.SetURLVars(req, map[string]string{"id": strconv.FormatUint(artistId, 10)})
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockSessionUseCase.EXPECT().GetUserId(sessionId).Return(userId, nil)
		mockArtistUseCase.EXPECT().IsLike(userId, artistId).Return(false, errors.New("error while checking like"))

		err := handler.IsLike(w, req)
		assert.Equal(t, http.StatusNotFound, err.(common_handler.StatusError).Code)
	})

	t.Run("IsLike Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/artist/is_like", nil)
		req = mux.SetURLVars(req, map[string]string{"id": strconv.FormatUint(artistId, 10)})
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockSessionUseCase.EXPECT().GetUserId(sessionId).Return(userId, nil)
		mockArtistUseCase.EXPECT().IsLike(userId, artistId).Return(isLiked, nil)

		err := handler.IsLike(w, req)
		assert.Nil(t, err)

		var received response.IsLiked
		err = json.NewDecoder(w.Body).Decode(&received)
		assert.Nil(t, err)
		assert.Equal(t, response.IsLiked{IsLiked: isLiked}, received)
	})

	t.Run("Unlike Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/artist/unlike", nil)
		req = mux.SetURLVars(req, map[string]string{"id": strconv.FormatUint(artistId, 10)})
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockSessionUseCase.EXPECT().GetUserId(sessionId).Return(userId, nil)
		mockArtistUseCase.EXPECT().Unlike(userId, artistId).Return(nil)

		err := handler.Unlike(w, req)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})
}
