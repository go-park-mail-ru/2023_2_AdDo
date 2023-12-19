package response

import (
	"main/internal/pkg/session"
	"net/http"
	"time"
)

type IsLiked struct {
	IsLiked bool `json:"IsLiked"`
}

func GetCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie(session.CookieName)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

func SetCookie(w http.ResponseWriter, sessionId string) {
	http.SetCookie(w, &http.Cookie{
		Name:     session.CookieName,
		Value:    sessionId,
		Expires:  time.Now().Add(session.TimeToLiveCookie),
		HttpOnly: true,
	})
}
