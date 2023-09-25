package storage_handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"main/storage"
	"net/http"
	"time"
)

type StorageHandler struct {
	database *storage.Database
}

func NewStorageHandler(db *sql.DB) *StorageHandler {
	return &StorageHandler{database: storage.NewDatabasePostgres(db)}
}

const CookieName = "JSESSIONID"

func ReadCookie(r *http.Request) (value string, err error) {
	cookie, err := r.Cookie(CookieName)
	if err != nil {
		return value, err
	}
	return cookie.Value, err
}

func (api *StorageHandler) Root(w http.ResponseWriter, r *http.Request) {
	var user storage.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"status": 500, "err": "json decoding error: check openapi"}`)
		return
	}

	sessionId, err := ReadCookie(r)
	if api.database.CheckSession(user.Id, sessionId) {
		fmt.Fprintf(w, "authorized\n")
	} else {
		fmt.Fprintf(w, "unauthorized\n")
	}
}

func (api *StorageHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var user storage.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"status": 500, "err": "json decoding error: check openapi"}`)
		return
	}

	id, err := api.database.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"status": 500, "err": "db_error"}`)
		return
	}

	if id == 0 {
		w.WriteHeader(http.StatusConflict)
		io.WriteString(w, `{"status": 409, "err": "user with such username has already created"}`)
		return
	}

	body := map[string]interface{}{
		"id": id,
	}

	err = json.NewEncoder(w).Encode(&body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"status": 500, "err": "json encoding error: check openapi"}`)
		return
	}
}

func (api *StorageHandler) Auth(w http.ResponseWriter, r *http.Request) {
	var user storage.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"status": 500, "err": "json decoding error: check openapi"}`)
		return
	}

	isUser := api.database.CheckUserCredentials(user)

	if !isUser {
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, `{"status": 401, "err": "wrong login or password"}`)
		return
	}

	sessionId, err := api.database.CreateNewSession(user.Id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"status": 500, "err": "db_error"}`)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    CookieName,
		Value:   sessionId,
		Expires: time.Now().Add(1 * time.Minute),
	})
}

func (api *StorageHandler) LogOut(w http.ResponseWriter, r *http.Request) {
	var user storage.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"status": 500, "err": "json decoding error: check openapi"}`)
		return
	}
	err = api.database.DeleteSession(user.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"status": 500, "err": "db_error"}`)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Expires: time.Now().Add(-1 * time.Second),
	})
}
