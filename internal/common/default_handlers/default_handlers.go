package default_handlers

import (
	"errors"
	common_handler "main/internal/common/handler"
	"net/http"
)

func NoAuth(w http.ResponseWriter, r *http.Request) error {
	return common_handler.StatusError{Code: http.StatusUnauthorized, Err: errors.New("no jsessionid in cookie")}
}

func Forbidden(w http.ResponseWriter, r *http.Request) error {
	return common_handler.StatusError{Code: http.StatusForbidden, Err: errors.New("no access")}
}

func BadRequest(w http.ResponseWriter, r *http.Request) error {
	return common_handler.StatusError{Code: http.StatusBadRequest, Err: errors.New("have no useful params")}
}
