package user_delivery

import (
	"bytes"
	"encoding/json"
	"errors"
	"image"
	"image/png"
	common_handler "main/internal/common/handler"
	avatar_domain "main/internal/pkg/image"
	"main/internal/pkg/session"
	user_domain "main/internal/pkg/user"
	session_mock "main/test/mocks/session"
	user_mock "main/test/mocks/user"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserUseCase := user_mock.NewMockUseCase(ctrl)
	mockSessionUseCase := session_mock.NewMockUseCase(ctrl)

	handler := NewHandler(mockUserUseCase, mockSessionUseCase, logrus.New())

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

		mockUserUseCase.EXPECT().Register(user).Return(nil)
		mockUserUseCase.EXPECT().Login(user.Email, user.Password).Return(sessionId, nil)
		err = handler.SignUp(w, req)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusNoContent, w.Code)

		assert.NoError(t, err)
	})
}

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserUseCase := user_mock.NewMockUseCase(ctrl)
	mockSessionUseCase := session_mock.NewMockUseCase(ctrl)

	handler := NewHandler(mockUserUseCase, mockSessionUseCase, logrus.New())

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

		mockUserUseCase.EXPECT().Login(user.Email, user.Password).Return("", errors.New("login failed"))

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

		mockUserUseCase.EXPECT().Login(user.Email, user.Password).Return(sessionIdExpected, nil)

		err = handler.Login(w, req)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusNoContent, w.Code)

		cookie := w.Result().Cookies()[firstCookie]
		sessionIdReceived := cookie.Value
		assert.Equal(t, sessionIdExpected, sessionIdReceived)
		assert.Equal(t, time.Now().Add(session.TimeToLiveCookie).Unix(), cookie.Expires.Unix())
	})
}

func TestAuth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserUseCase := user_mock.NewMockUseCase(ctrl)
	mockSessionUseCase := session_mock.NewMockUseCase(ctrl)

	handler := NewHandler(mockUserUseCase, mockSessionUseCase, logrus.New())

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

		mockUserUseCase.EXPECT().Auth(sessionId).Return(false, errors.New("unauthorized"))

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

		mockUserUseCase.EXPECT().Auth(sessionId).Return(true, nil)

		err := handler.Auth(w, req)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})
}

func TestLogOut(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserUseCase := user_mock.NewMockUseCase(ctrl)
	mockSessionUseCase := session_mock.NewMockUseCase(ctrl)

	handler := NewHandler(mockUserUseCase, mockSessionUseCase, logrus.New())

	t.Run("Success", func(t *testing.T) {
		const sessionId = "sessionID"
		cookie := http.Cookie{
			Name:     session.CookieName,
			Value:    sessionId,
			Expires:  time.Now().Add(session.TimeToLiveCookie),
			Secure:   true,
			HttpOnly: true,
		}

		mockUserUseCase.EXPECT().Logout(sessionId).Return(nil)

		req := httptest.NewRequest(http.MethodDelete, "/logout", nil)
		req.AddCookie(&cookie)

		w := httptest.NewRecorder()

		err := handler.LogOut(w, req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, w.Code)

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

func TestUploadAvatar(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const url = "/upload_avatar"

	mockUserUseCase := user_mock.NewMockUseCase(ctrl)
	mockSessionUseCase := session_mock.NewMockUseCase(ctrl)

	handler := NewHandler(mockUserUseCase, mockSessionUseCase, logrus.New())

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("Avatar", "images.png")
	if err != nil {
		t.Error(err)
	}

	width := 32
	height := 24

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	err = png.Encode(part, img)
	if err != nil {
		t.Error(err)
	}

	writer.Close()

	t.Run("Success", func(t *testing.T) {
		const sessionID = "sessionID"
		cookie := http.Cookie{
			Name:     session.CookieName,
			Value:    sessionID,
			Expires:  time.Now().Add(session.TimeToLiveCookie),
			Secure:   true,
			HttpOnly: true,
		}
		const userId = "1"

		req := httptest.NewRequest(http.MethodPost, url, body)
		req.Header.Add("Content-Type", writer.FormDataContentType())
		req.AddCookie(&cookie)

		src, hdr, err := req.FormFile("Avatar")
		if err != nil {
			t.Error(err)
		}
		mockSessionUseCase.EXPECT().GetUserId(sessionID).Return(userId, nil)
		mockUserUseCase.EXPECT().UploadAvatar(userId, src, hdr.Size)

		w := httptest.NewRecorder()

		err = handler.UploadAvatar(w, req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("RequestWithoutFile", func(t *testing.T) {
		const sessionID = "sessionID"
		cookie := http.Cookie{
			Name:     session.CookieName,
			Value:    sessionID,
			Expires:  time.Now().Add(session.TimeToLiveCookie),
			Secure:   true,
			HttpOnly: true,
		}
		const userId = "1"

		req := httptest.NewRequest(http.MethodPost, url, nil)
		req.Header.Add("Content-Type", writer.FormDataContentType())
		req.AddCookie(&cookie)

		mockSessionUseCase.EXPECT().GetUserId(sessionID).Return(userId, nil)

		w := httptest.NewRecorder()

		err = handler.UploadAvatar(w, req)
		assert.Error(t, err)
		assert.Equal(t, http.StatusBadRequest, err.(common_handler.Error).Status())
	})

	t.Run("Unauthorized", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, url, body)

		w := httptest.NewRecorder()

		err = handler.UploadAvatar(w, req)
		assert.Error(t, err)
		assert.Equal(t, http.StatusUnauthorized, err.(common_handler.StatusError).Code)
	})
}

func TestRemoveAvatar(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const url = "/remove_avatar"

	mockUserUseCase := user_mock.NewMockUseCase(ctrl)
	mockSessionUseCase := session_mock.NewMockUseCase(ctrl)

	handler := NewHandler(mockUserUseCase, mockSessionUseCase, logrus.New())

	t.Run("Success", func(t *testing.T) {
		const sessionId = "sessionID"
		cookie := http.Cookie{
			Name:     session.CookieName,
			Value:    sessionId,
			Expires:  time.Now().Add(session.TimeToLiveCookie),
			Secure:   true,
			HttpOnly: true,
		}
		const userId = "1"

		mockSessionUseCase.EXPECT().GetUserId(sessionId).Return(userId, nil)
		mockUserUseCase.EXPECT().RemoveAvatar(userId).Return(nil)

		req := httptest.NewRequest(http.MethodPut, url, nil)
		req.AddCookie(&cookie)

		w := httptest.NewRecorder()

		err := handler.RemoveAvatar(w, req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})

	t.Run("Conflict", func(t *testing.T) {
		const sessionId = "sessionID"
		cookie := http.Cookie{
			Name:     session.CookieName,
			Value:    sessionId,
			Expires:  time.Now().Add(session.TimeToLiveCookie),
			Secure:   true,
			HttpOnly: true,
		}
		const userId = "1"

		mockSessionUseCase.EXPECT().GetUserId(sessionId).Return(userId, nil)
		mockUserUseCase.EXPECT().RemoveAvatar(userId).Return(avatar_domain.ErrAvatarDoesNotExist)

		req := httptest.NewRequest(http.MethodPut, url, nil)
		req.AddCookie(&cookie)

		w := httptest.NewRecorder()

		err := handler.RemoveAvatar(w, req)

		assert.Error(t, err)
		assert.Equal(t, http.StatusConflict, err.(common_handler.StatusError).Code)
	})

	t.Run("Unauthorized", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, url, nil)

		w := httptest.NewRecorder()

		err := handler.RemoveAvatar(w, req)
		assert.Error(t, err)
		assert.Equal(t, http.StatusUnauthorized, err.(common_handler.StatusError).Code)
	})
}

func TestForgotPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const url = "/auth/forgot_password"

	mockUserUseCase := user_mock.NewMockUseCase(ctrl)
	mockSessionUseCase := session_mock.NewMockUseCase(ctrl)

	handler := NewHandler(mockUserUseCase, mockSessionUseCase, logrus.New())

	t.Run("Successful", func(t *testing.T) {
		email := user_domain.ForgotPasswordInput{Email: "user@mail.ru"}
		body, err := json.Marshal(email)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))

		mockUserUseCase.EXPECT().CheckEmailExist(email.Email).Times(1).Return(nil)
		mockUserUseCase.EXPECT().SendResetToken(email.Email).Times(1).Return(nil)

		w := httptest.NewRecorder()

		err = handler.ForgotPassword(w, req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, w.Result().StatusCode)
	})

	t.Run("BadEmail", func(t *testing.T) {
		email := user_domain.ForgotPasswordInput{Email: "user"}
		body, err := json.Marshal(email)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))

		w := httptest.NewRecorder()

		err = handler.ForgotPassword(w, req)

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusBadRequest, err.(common_handler.StatusError).Code)
	})

	t.Run("UserWithEmailDoesNotExist", func(t *testing.T) {
		email := user_domain.ForgotPasswordInput{Email: "user@mail.ru"}
		body, err := json.Marshal(email)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))

		w := httptest.NewRecorder()

		mockUserUseCase.EXPECT().CheckEmailExist(email.Email).Times(1).Return(user_domain.ErrUserDoesNotExist)

		err = handler.ForgotPassword(w, req)

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusBadRequest, err.(common_handler.StatusError).Code)
	})

	t.Run("InternalError", func(t *testing.T) {
		email := user_domain.ForgotPasswordInput{Email: "user@mail.ru"}
		body, err := json.Marshal(email)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))

		w := httptest.NewRecorder()

		mockUserUseCase.EXPECT().CheckEmailExist(email.Email).Times(1).Return(nil)
		mockUserUseCase.EXPECT().SendResetToken(email.Email).Times(1).Return(errors.New("internal error"))

		err = handler.ForgotPassword(w, req)

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusInternalServerError, err.(common_handler.StatusError).Code)
	})
}

func TestResetPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const (
		resetToken = "anytoken"
		email      = "user@mail.ru"
		url        = "/auth/reset_password/"
	)

	mockUserUseCase := user_mock.NewMockUseCase(ctrl)
	mockSessionUseCase := session_mock.NewMockUseCase(ctrl)

	handler := NewHandler(mockUserUseCase, mockSessionUseCase, logrus.New())

	passwordsInput := user_domain.ResetPasswordInput{Password: "password", ConfirmPassword: "password"}

	t.Run("Successful", func(t *testing.T) {
		body, err := json.Marshal(passwordsInput)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
		req = mux.SetURLVars(req, map[string]string{"reset_token": resetToken})

		w := httptest.NewRecorder()

		mockUserUseCase.EXPECT().CheckTokenExist(resetToken).Times(1).Return(email, nil)
		mockUserUseCase.EXPECT().UpdatePassword(email, passwordsInput.Password).Times(1).Return(nil)

		err = handler.ResetPassword(w, req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, w.Result().StatusCode)
	})

	t.Run("PasswordDoNotMatch", func(t *testing.T) {
		diffPasswords := user_domain.ResetPasswordInput{Password: "password", ConfirmPassword: "passwordd"}
		body, err := json.Marshal(diffPasswords)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
		req = mux.SetURLVars(req, map[string]string{"reset_token": resetToken})

		w := httptest.NewRecorder()

		err = handler.ResetPassword(w, req)

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusBadRequest, err.(common_handler.StatusError).Code)
	})

	t.Run("TokenDoesNotExist", func(t *testing.T) {
		body, err := json.Marshal(passwordsInput)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
		req = mux.SetURLVars(req, map[string]string{"reset_token": resetToken})

		w := httptest.NewRecorder()

		mockUserUseCase.EXPECT().CheckTokenExist(resetToken).Times(1).Return("", errors.New("token does not exit"))

		err = handler.ResetPassword(w, req)

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusBadRequest, err.(common_handler.StatusError).Code)
	})

	t.Run("InernalError", func(t *testing.T) {
		body, err := json.Marshal(passwordsInput)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
		req = mux.SetURLVars(req, map[string]string{"reset_token": resetToken})

		w := httptest.NewRecorder()

		mockUserUseCase.EXPECT().CheckTokenExist(resetToken).Times(1).Return(email, nil)
		mockUserUseCase.EXPECT().UpdatePassword(email, passwordsInput.Password).Times(1).Return(errors.New("internal error"))

		err = handler.ResetPassword(w, req)

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusInternalServerError, err.(common_handler.StatusError).Code)
	})
}
