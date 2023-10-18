package user_delivery

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	common_handler "main/internal/pkg/common/handler"
	"main/internal/pkg/session"
	user_domain "main/internal/pkg/user"
	user_mock "main/test/mocks/user"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := user_mock.NewMockUseCase(ctrl)

	handler := NewHandler(mockUseCase, logrus.New())

	t.Run("DecodeRequestBodyError", func(t *testing.T) {
		requestBody := bytes.NewBufferString(`{"name": "John"`) // некорректный JSON

		req := httptest.NewRequest(http.MethodPost, "/signup", requestBody)
		w := httptest.NewRecorder()

		err := handler.SignUp(w, req)

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusBadRequest, err.(common_handler.StatusError).Code)
	})

	t.Run("RegisterError", func(t *testing.T) {
		user := user_domain.User{Username: "John", Email: "john@example.com"}

		requestBody, err := json.Marshal(user)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(requestBody))
		w := httptest.NewRecorder()

		err = handler.SignUp(w, req)

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusBadRequest, err.(common_handler.StatusError).Code)
	})

	t.Run("Invalid password", func(t *testing.T) {
		user := user_domain.User{Username: "John", Email: "john@example.com", Password: "gshdfghlsdhflgkhslkdfhglkhskdfhglksdfhgshkfhgksdhghslsdfgkjdgf", BirthDate: "12-12-2003"}

		requestBody, err := json.Marshal(user)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(requestBody))
		w := httptest.NewRecorder()

		err = handler.SignUp(w, req)

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusBadRequest, err.(common_handler.StatusError).Code)
	})

	t.Run("Success", func(t *testing.T) {
		user := user_domain.User{Id: "qwer-qwer-qwer", Username: "John", Email: "john@example.com", BirthDate: "12-12-2003", Password: "password"}
		const sessionId = "sesId"
		requestBody, err := json.Marshal(user)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(requestBody))
		w := httptest.NewRecorder()

		mockUseCase.EXPECT().Register(user).Return(nil)
		mockUseCase.EXPECT().Login(user.Email, user.Password).Return(sessionId, nil)
		err = handler.SignUp(w, req)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, w.Code)

		assert.NoError(t, err)
	})
}

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := user_mock.NewMockUseCase(ctrl)
	handler := NewHandler(mockUseCase, logrus.New())

	t.Run("DecodeRequestBodyError", func(t *testing.T) {
		requestBody := bytes.NewBufferString(`{"email": "john@example.`) // некорректный JSON

		req := httptest.NewRequest(http.MethodPost, "/login", requestBody)
		w := httptest.NewRecorder()

		err := handler.Login(w, req)

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusBadRequest, err.(common_handler.StatusError).Code)
	})

	t.Run("LoginError", func(t *testing.T) {
		user := user_domain.User{Email: "john@example.com", Password: "password"}

		requestBody, err := json.Marshal(user)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestBody))
		w := httptest.NewRecorder()

		mockUseCase.EXPECT().Login(user.Email, user.Password).Return("", errors.New("login failed"))

		err = handler.Login(w, req)

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusForbidden, err.(common_handler.StatusError).Code)
	})

	t.Run("Success", func(t *testing.T) {
		user := user_domain.User{Email: "john@example.com", Password: "password"}
		const sessionIdExpected = "sessionIdExpected"
		const firstCookie = 0

		requestBody, err := json.Marshal(user)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestBody))
		w := httptest.NewRecorder()

		mockUseCase.EXPECT().Login(user.Email, user.Password).Return(sessionIdExpected, nil)

		err = handler.Login(w, req)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, w.Code)

		cookie := w.Result().Cookies()[firstCookie]
		sessionIdReceived := cookie.Value
		assert.Equal(t, sessionIdExpected, sessionIdReceived)
		assert.Equal(t, time.Now().Add(session.TimeToLiveCookie).Unix(), cookie.Expires.Unix())
	})
}

func TestAuth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := user_mock.NewMockUseCase(ctrl)

	handler := NewHandler(mockUseCase, logrus.New())

	t.Run("GetCookieError", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/auth", nil)
		w := httptest.NewRecorder()

		err := handler.Auth(w, req)

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusUnauthorized, err.(common_handler.StatusError).Code)
	})

	t.Run("AuthError", func(t *testing.T) {
		sessionId := "sessionID"
		cookie := http.Cookie{
			Name:     session.CookieName,
			Value:    sessionId,
			Expires:  time.Now().Add(session.TimeToLiveCookie),
			Secure:   true,
			HttpOnly: true,
		}
		req := httptest.NewRequest(http.MethodGet, "/auth", nil)
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockUseCase.EXPECT().Auth(sessionId).Return(false, errors.New("unauthorized"))

		err := handler.Auth(w, req)

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusUnauthorized, err.(common_handler.StatusError).Code)
	})

	t.Run("Success", func(t *testing.T) {
		const sessionId = "sessionID"
		cookie := http.Cookie{
			Name:     session.CookieName,
			Value:    sessionId,
			Expires:  time.Now().Add(session.TimeToLiveCookie),
			Secure:   true,
			HttpOnly: true,
		}

		req := httptest.NewRequest(http.MethodGet, "/auth", nil)
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockUseCase.EXPECT().Auth(sessionId).Return(true, nil)

		err := handler.Auth(w, req)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestLogOut(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := user_mock.NewMockUseCase(ctrl)
	handler := UserHandler{userUseCase: mockUseCase, logger: logrus.New()}

	t.Run("Success", func(t *testing.T) {
		const sessionId = "sessionID"
		cookie := http.Cookie{
			Name:     session.CookieName,
			Value:    sessionId,
			Expires:  time.Now().Add(session.TimeToLiveCookie),
			Secure:   true,
			HttpOnly: true,
		}

		mockUseCase.EXPECT().Logout(sessionId).Return(nil)

		req := httptest.NewRequest(http.MethodDelete, "/logout", nil)
		req.AddCookie(&cookie)

		w := httptest.NewRecorder()

		err := handler.LogOut(w, req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)

		cookies := w.Result().Cookies()

		if len(cookies) != 0 {
			assert.True(t, cookies[0].Expires.Before(time.Now()))
		}
	})

	t.Run("Unauthorized", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/logout", nil)
		w := httptest.NewRecorder()

		err := handler.LogOut(w, req)
		assert.Error(t, err)
	})
}
