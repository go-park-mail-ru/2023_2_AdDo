package common_handler

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Error interface {
	error
	Status() int
}

type StatusError struct {
	Code int
	Err  error
}

func (se StatusError) Error() string {
	return se.Err.Error()
}

func (se StatusError) Status() int {
	return se.Code
}

type Handler struct {
	H func(w http.ResponseWriter, r *http.Request) error
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.H(w, r)
	if err != nil {
		var e Error
		switch {
		case errors.As(err, &e):
			w.WriteHeader(e.Status())
			io.WriteString(w, fmt.Sprintf("{ \"status\": %d, \"err\": \"%s\" }\n", e.Status(), e.Error()))
		default:
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, fmt.Sprintf("{ \"status\": %d, \"err\": \"%s\" }\n", http.StatusInternalServerError, e.Error()))
		}
	}
}
