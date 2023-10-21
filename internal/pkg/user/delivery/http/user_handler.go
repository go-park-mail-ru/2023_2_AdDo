package user_delivery

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	common_handler "main/internal/pkg/common/handler"
	"main/internal/pkg/common/response"
	"main/internal/pkg/common/utils"
	"main/internal/pkg/session"
	user_domain "main/internal/pkg/user"
	"net/http"
	"time"

	"github.com/gorilla/csrf"
)

type UserHandler struct {
	userUseCase    user_domain.UseCase
	sessionUseCase session.UseCase
	logger      *logrus.Logger
}

func NewHandler(userUseCase user_domain.UseCase, sessionUseCase session.UseCase, logger *logrus.Logger) UserHandler {
	handler := UserHandler{
		userUseCase:    userUseCase,
		sessionUseCase: sessionUseCase,
		logger:      logger,
	}
	logger.Infoln("New User Handler successfully created")
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
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("SignUp Handler entered")

	var u user_domain.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("User model decoded from request body")

	err := u.Validate()
	if err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("User model is valid")

	err = handler.userUseCase.Register(u)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusConflict, Err: err}
	}
	handler.logger.Infoln("User signed up")

	sessionId, err := handler.userUseCase.Login(u.Email, u.Password)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusForbidden, Err: err}
	}
	handler.logger.Infoln("Session for user created")

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
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("Login Handler entered")

	var credentials user_domain.UserCredentials
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("User credentials parsed from body")

	err := credentials.Validate()
	if err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("User credentials are valid")

	sessionId, err := handler.userUseCase.Login(credentials.Email, credentials.Password)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusForbidden, Err: err}
	}
	handler.logger.Infoln("User got a new session id")

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
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("Auth Handler entered")

	w.Header().Set(session.XCsrfToken, csrf.Token(r))
	handler.logger.Infoln("csrf token added to response headers")

	sessionId, err := response.GetCookie(r)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}
	handler.logger.Infoln("session id got")

	isAuth, err := handler.userUseCase.Auth(sessionId)
	if err != nil || !isAuth {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}
	handler.logger.Infoln("auth success")

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
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("LogOut Handler entered")

	sessionId, err := response.GetCookie(r)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}
	handler.logger.Infoln("Got session id from cookie")

	if err := handler.userUseCase.Logout(sessionId); err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}
	handler.logger.Infoln("Session deleted from database")

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
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("Me Handler entered")

	sessionId, err := response.GetCookie(r)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}
	handler.logger.Infoln("Got session id from cookie")

	user, err := handler.userUseCase.GetUserInfo(sessionId)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}
	handler.logger.Infoln("Got user info from database successfully")

	err = response.RenderJSON(w, user)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	handler.logger.Infoln("Response body rendered")

	return nil
}

// UploadAvatar
//
// @Description	Upload user avatar
// @Tags		user
// @Param		Avatar	formData	file		true	"User avatar"
// @Security	cookieAuth
// @Security	csrfToken
// @Security	cookieCsrfToken
// @Success		200	{object}	user_domain.UploadAvatarResponse
// @Failure		400	{string}	errMsg
// @Failure		401	{string}	errMsg
// @Failure		500	{string}	errMsg
// @Router		/upload_avatar [post]
func (handler *UserHandler) UploadAvatar(w http.ResponseWriter, r *http.Request) error {
	sessionId, err := response.GetCookie(r)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}

	userId, err := handler.sessionUseCase.GetUserId(sessionId)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}
	
	src, hdr, err := r.FormFile("Avatar")
	if err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	defer src.Close()

	url, err := handler.userUseCase.UploadAvatar(userId, src, hdr.Size)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}

	response, err := json.Marshal(user_domain.UploadAvatarResponse{Url: url})
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	w.Write(response)

	return nil
}

// RemoveAvatar
//
// @Description	Remove user avatar
// @Tags		user
// @Security	cookieAuth
// @Security	csrfToken
// @Security	cookieCsrfToken
// @Success		200
// @Failure		401	{string}	errMsg
// @Failure		409	{string}	errMsg
// @Failure		500	{string}	errMsg
// @Router		/remove_avatar [post]
func (handler *UserHandler) RemoveAvatar(w http.ResponseWriter, r *http.Request) error {
	sessionId, err := response.GetCookie(r)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}

	userId, err := handler.sessionUseCase.GetUserId(sessionId)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}

	if err = handler.userUseCase.RemoveAvatar(userId); err != nil {
		return common_handler.StatusError{Code: http.StatusConflict, Err: err}
	}

	return nil
}
