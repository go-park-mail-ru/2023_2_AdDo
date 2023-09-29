package response

import (
	"encoding/json"
	"main/internal/pkg/session"
	"net/http"
)

func GetCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie(session.CookieName)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

func RenderJSON(w http.ResponseWriter, v interface{}) error {
	jsonResponse, err := json.Marshal(v)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonResponse)
	return err
}
