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

//func RenderEasyJSON(w http.ResponseWriter, v easyjson.Marshaler) error {
//	jsonResponse, err := easyjson.Marshal(v)
//	if err != nil {
//		return err
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	_, err = w.Write(jsonResponse)
//	return err
//}

func SetCookie(w http.ResponseWriter, sessionId string) {
	http.SetCookie(w, &http.Cookie{
		Name:     session.CookieName,
		Value:    sessionId,
		Expires:  time.Now().Add(session.TimeToLiveCookie),
		HttpOnly: true,
	})
}
