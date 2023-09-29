package common_handler

import (
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
		switch e := err.(type) {
		case Error:
			w.WriteHeader(e.Status())
			io.WriteString(w, fmt.Sprintf("{\"status\": %d, \"err\": \"%s\"}", e.Status(), e.Error()))
			http.Error(w, e.Error(), e.Status())
		default:
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, fmt.Sprintf("{\"status\": %d, \"err\": \"%s\"}", http.StatusInternalServerError, e.Error()))
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}
