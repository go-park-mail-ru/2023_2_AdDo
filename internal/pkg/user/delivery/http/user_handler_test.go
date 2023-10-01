package user_delivery

import (
	"bytes"
	"encoding/json"
	"errors"
	common_handler "main/internal/pkg/common/handler"
	"main/internal/pkg/session"
	user_domain "main/internal/pkg/user"
	user_mock "main/test/mocks/user"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создаем mock для UseCase
	mockUseCase := user_mock.NewMockUseCase(ctrl)

	handler := NewHandler(mockUseCase)

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

		mockUseCase.EXPECT().Register(user).Return(uint64(0), errors.New("registration failed"))

		err = handler.SignUp(w, req)

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusConflict, err.(common_handler.StatusError).Code)
	})

	t.Run("Success", func(t *testing.T) {
		user := user_domain.User{Username: "John", Email: "john@example.com"}

		requestBody, err := json.Marshal(user)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(requestBody))
		w := httptest.NewRecorder()

		mockUseCase.EXPECT().Register(user).Return(uint64(123), nil)

		err = handler.SignUp(w, req)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, w.Code)

		expectedResponse := user_domain.ResponseId{Id: 123}
		expectedResponseBody, err := json.Marshal(expectedResponse)
		assert.NoError(t, err)
		assert.Equal(t, expectedResponseBody, w.Body.Bytes())
	})
}

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создаем mock для UseCase
	mockUseCase := user_mock.NewMockUseCase(ctrl)

	handler := NewHandler(mockUseCase)

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

		mockUseCase.EXPECT().Login(user.Email, user.Password).Return(uint64(0), "", errors.New("login failed"))

		err = handler.Login(w, req)

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusForbidden, err.(common_handler.StatusError).Code)
	})

	t.Run("Success", func(t *testing.T) {
		user := user_domain.User{Email: "john@example.com", Password: "password"}

		requestBody, err := json.Marshal(user)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestBody))
		w := httptest.NewRecorder()

		mockUseCase.EXPECT().Login(user.Email, user.Password).Return(uint64(123), "sessionID", nil)

		err = handler.Login(w, req)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
		cookie := w.Result().Cookies()[0]
		assert.Equal(t, "sessionID", cookie.Value)
		assert.Equal(t, time.Now().Add(session.TimeToLive).Unix(), cookie.Expires.Unix())
	})
}

func TestAuth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создаем mock для UseCase
	mockUseCase := user_mock.NewMockUseCase(ctrl)

	handler := NewHandler(mockUseCase)

	t.Run("InvalidID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/auth?id=invalid", nil)
		w := httptest.NewRecorder()

		err := handler.Auth(w, req)

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusBadRequest, err.(common_handler.StatusError).Code)
	})

	t.Run("GetCookieError", func(t *testing.T) {
		id := 123
		req := httptest.NewRequest(http.MethodGet, "/auth?id="+strconv.Itoa(id), nil)
		w := httptest.NewRecorder()

		err := handler.Auth(w, req)

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusUnauthorized, err.(common_handler.StatusError).Code)
	})

	t.Run("AuthError", func(t *testing.T) {
		id := 123
		sessionId := "sessionID"
		cookie := http.Cookie{
			Name:     session.CookieName,
			Value:    sessionId,
			Expires:  time.Now().Add(session.TimeToLive),
			Secure:   true,
			HttpOnly: true,
		}
		req := httptest.NewRequest(http.MethodGet, "/auth?id="+strconv.Itoa(id), nil)
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockUseCase.EXPECT().Auth(uint64(id), sessionId).Return(false, errors.New("unauthorized"))

		err := handler.Auth(w, req)

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusUnauthorized, err.(common_handler.StatusError).Code)
	})

	t.Run("Success", func(t *testing.T) {
		id := 123
		sessionId := "sessionID"
		cookie := http.Cookie{
			Name:     session.CookieName,
			Value:    sessionId,
			Expires:  time.Now().Add(session.TimeToLive),
			Secure:   true,
			HttpOnly: true,
		}
		user := user_domain.User{Id: uint64(id), Username: "John"}

		req := httptest.NewRequest(http.MethodGet, "/auth?id="+strconv.Itoa(id), nil)
		req.AddCookie(&cookie)
		w := httptest.NewRecorder()

		mockUseCase.EXPECT().Auth(uint64(id), sessionId).Return(true, nil)
		mockUseCase.EXPECT().GetUserInfo(uint64(id)).Return(user, nil)

		err := handler.Auth(w, req)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestLogOut(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := user_mock.NewMockUseCase(ctrl)
	handler := UserHandler{userUseCase: mockUseCase}

	t.Run("Success", func(t *testing.T) {
		userID := 1
		u := user_domain.User{Id: uint64(userID)}

		mockUseCase.EXPECT().Logout(uint64(userID)).Return(nil)

		reqBody, _ := json.Marshal(u)
		req := httptest.NewRequest(http.MethodPost, "/logout", bytes.NewBuffer(reqBody))
		w := httptest.NewRecorder()

		err := handler.LogOut(w, req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)

		// Проверяем, что cookie установлен с временем завершения
		cookies := w.Result().Cookies()

		if len(cookies) != 0 {
			assert.True(t, cookies[0].Expires.Before(time.Now()))
		}
	})

	t.Run("Unauthorized", func(t *testing.T) {
		userID := 1
		u := user_domain.User{Id: uint64(userID)}

		mockUseCase.EXPECT().Logout(uint64(userID)).Return(errors.New("unauthorized"))

		reqBody, _ := json.Marshal(u)
		req := httptest.NewRequest(http.MethodPost, "/logout", bytes.NewBuffer(reqBody))
		w := httptest.NewRecorder()

		err := handler.LogOut(w, req)

		assert.Error(t, err)
	})
}
