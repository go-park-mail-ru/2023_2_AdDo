package user_delivery

import (
	"encoding/json"
	common_handler "main/internal/pkg/common/handler"
	"main/internal/pkg/common/response"
	user_domain "main/internal/pkg/user"
	"net/http"
	"strconv"
	"time"
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

func (handler *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) error {
	var u user_domain.User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	err := u.Validate()
	if err != nil {
		return err
	}

	err = handler.userUseCase.Register(u)
	if err != nil {
		return err
	}

	id, sessionId, err := handler.userUseCase.Login(u.Email, u.Password)
	if err != nil {
		return err
	}
	response.SetCookie(w, sessionId)

	if err != nil || id == 0 {
		return common_handler.StatusError{Code: http.StatusConflict, Err: err}
	}

	err = response.RenderJSON(w, user_domain.ResponseId{Id: id})
	if err != nil {
		return err
	}

	return nil
}

func (handler *UserHandler) Login(w http.ResponseWriter, r *http.Request) error {
	var u user_domain.User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}

	id, sessionId, err := handler.userUseCase.Login(u.Email, u.Password)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusForbidden, Err: err}
	}

	response.SetCookie(w, sessionId)
	err = response.RenderJSON(w, user_domain.ResponseId{Id: id})
	return err
}

func (handler *UserHandler) Auth(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	sessionId, err := response.GetCookie(r)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}

	isAuth, err := handler.userUseCase.Auth(uint64(id), sessionId)

	if err != nil || !isAuth {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}

	user, err := handler.userUseCase.GetUserInfo(uint64(id))

	err = response.RenderJSON(w, user)
	return err
}

func (handler *UserHandler) LogOut(w http.ResponseWriter, r *http.Request) error {
	var u user_domain.User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}

	if err := handler.userUseCase.Logout(u.Id); err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}

	http.SetCookie(w, &http.Cookie{
		Expires: time.Now().Add(-1 * time.Second),
	})
	return nil
}
