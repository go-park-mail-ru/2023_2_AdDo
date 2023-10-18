package user_delivery

import (
	"encoding/json"
	common_handler "main/internal/pkg/common/handler"
	"main/internal/pkg/common/response"
	"main/internal/pkg/session"
	user_domain "main/internal/pkg/user"
	"net/http"
	"time"

	"github.com/gorilla/csrf"
)

type UserHandler struct {
	userUseCase user_domain.UseCase
}

func NewHandler(userUseCase user_domain.UseCase) UserHandler {
	handler := UserHandler{
		userUseCase: userUseCase,
	}
	return handler
}

// SignUp
//
//	@Description	register user
//	@Tags			user
//	@Security		csrfToken
//	@Security		cookieCsrfToken
//	@Accept			json
//	@Param			userData	body	user_domain.User	true	"User data"
//	@Success		200
//	@Failure		400	{string}	errMsg
//	@Failure		403	{string}	errMsg
//	@Failure		409	{string}	errMsg
//	@Failure		500	{string}	errMsg
//	@Header			200	{string}	Set-Cookie	"Set JSESSIONID in Cookie"
//	@Router			/sign_up [post]
func (handler *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) error {
	var u user_domain.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}

	err := u.Validate()
	if err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}

	err = handler.userUseCase.Register(u)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusConflict, Err: err}
	}

	sessionId, err := handler.userUseCase.Login(u.Email, u.Password)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusForbidden, Err: err}
	}

	response.SetCookie(w, sessionId)
	return nil
}

// Login
//
//	@Description	login user
//	@Tags			user
//	@Security		csrfToken
//	@Security		cookieCsrfToken
//	@Accept			json
//	@Param			userCrds	body	user_domain.UserCredentials	true	"User email and password"
//	@Success		200
//	@Failure		400	{string}	errMsg
//	@Failure		403	{string}	errMsg
//	@Failure		500	{string}	errMsg
//	@Header			200	{string}	Set-Cookie	"Set JSESSIONID in Cookie"
//	@Router			/login [post]
func (handler *UserHandler) Login(w http.ResponseWriter, r *http.Request) error {
	var credentials user_domain.UserCredentials
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}

	err := credentials.Validate()
	if err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}

	sessionId, err := handler.userUseCase.Login(credentials.Email, credentials.Password)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusForbidden, Err: err}
	}

	response.SetCookie(w, sessionId)
	return nil
}

// Auth
//
//	@Description	check user's authentication by cookie and set csrf-token
//	@Tags			user
//	@Security		cookieAuth
//	@Success		200
//	@Failure		401		{string}	errMsg
//	@Failure		500		{string}	errMsg
//	@Header			200,401	{string}	Set-Cookie		"Set X-Csrf-Token in Cookie"
//	@Header			200,401	{string}	X-Csrf-Token	"Set X-Csrf-Token in header"
//	@Router			/auth [get]
func (handler *UserHandler) Auth(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set(session.XCsrfToken, csrf.Token(r))
	sessionId, err := response.GetCookie(r)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}

	isAuth, err := handler.userUseCase.Auth(sessionId)

	if err != nil || !isAuth {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}

	return nil
}

// LogOut
//
//	@Description	logout user
//	@Tags			user
//	@Security		cookieAuth
//	@Security		csrfToken
//	@Security		cookieCsrfToken
//	@Success		200
//	@Failure		401	{string}	errMsg
//	@Failure		403	{string}	errMsg
//	@Failure		500	{string}	errMsg
//	@Router			/logout [post]
func (handler *UserHandler) LogOut(w http.ResponseWriter, r *http.Request) error {
	sessionId, err := response.GetCookie(r)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}

	if err := handler.userUseCase.Logout(sessionId); err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}

	http.SetCookie(w, &http.Cookie{
		Expires: time.Now().Add(-1 * time.Second),
	})
	return nil
}

// Me
//
//	@Description	get user info
//	@Tags			user
//	@Produce		json
//	@Security		cookieAuth
//	@Security		csrfToken
//	@Security		cookieCsrfToken
//	@Success		200	{object}	user_domain.User
//	@Failure		401	{string}	errMsg
//	@Failure		403	{string}	errMsg
//	@Failure		500	{string}	errMsg
//	@Router			/me [get]
func (handler *UserHandler) Me(w http.ResponseWriter, r *http.Request) error {
	sessionId, err := response.GetCookie(r)

	if err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}

	user, err := handler.userUseCase.GetUserInfo(sessionId)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}

	err = response.RenderJSON(w, user)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	return nil
}
