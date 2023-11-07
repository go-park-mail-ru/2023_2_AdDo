package middleware

import (
	"github.com/gorilla/csrf"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"main/internal/pkg/session"
)

func NewCors() mux.MiddlewareFunc {
	return handlers.CORS(
		handlers.AllowedOrigins([]string{"http://82.146.45.164:8081"}),
		handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "X-Csrf-Token"}),
		handlers.ExposedHeaders([]string{session.CookieName}),
		handlers.AllowCredentials(),
	)
}

func NewCSRF() mux.MiddlewareFunc {
	return csrf.Protect(
		session.CSRFKey,
		csrf.Secure(false),
		csrf.HttpOnly(false),
		csrf.MaxAge(session.TimeToLiveCSRF),
		csrf.RequestHeader("X-Csrf-Token"),
		csrf.CookieName("X-Csrf-Token"),
		csrf.FieldName("X-Csrf-Token"),
	)
}
