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

func ReadCookie(r *http.Request) string {
	cookie, err := r.Cookie(CookieName)
	if err != nil {
		panic(err)
	}
	return cookie.Value
}

func RenderJSON(w http.ResponseWriter, v interface{}) {
	jsonResponse, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func ParseJSON(r *http.Request, dataStruct interface{}) {
	err := json.NewDecoder(r.Body).Decode(&dataStruct)
	if err != nil {
		panic(err)
	}
}

func (api *StorageHandler) Root(w http.ResponseWriter, r *http.Request) {
	var user storage.User
	ParseJSON(r, user)

	sessionId := ReadCookie(r)
	if api.database.CheckSession(user.Id, sessionId) {
		fmt.Fprintf(w, "authorized\n")
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "unauthorized\n")
	}
}

func (api *StorageHandler) Home(w http.ResponseWriter, r *http.Request) {
	var user storage.User
	ParseJSON(r, user)

	sessionId := ReadCookie(r)
	if api.database.CheckSession(user.Id, sessionId) {
		fmt.Fprintf(w, "authorized\n")
	} else {
		fmt.Fprintf(w, "unauthorized\n")
	}
}

func (api *StorageHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var user storage.User
	ParseJSON(r, user)

	id, err := api.database.CreateUser(user)
	if err != nil {
		panic(err)
	}

	if id == 0 {
		w.WriteHeader(http.StatusConflict)
		io.WriteString(w, `{"status": 409, "err": "user with such username has already created"}`)
		return
	}

	RenderJSON(w, storage.ResponseId{Id: id})
}

func (api *StorageHandler) Auth(w http.ResponseWriter, r *http.Request) {
	var user storage.User

	ParseJSON(r, user)

	isUser := api.database.CheckUserCredentials(user)

	if !isUser {
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, `{"status": 401, "err": "wrong login or password"}`)
		return
	}

	sessionId, err := api.database.CreateNewSession(user.Id)
	if err != nil {
		panic(err)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     CookieName,
		Value:    sessionId,
		Expires:  time.Now().Add(1 * time.Minute),
		Secure:   true,
		HttpOnly: true,
	})
	w.Header().Set("csrf", "my token")
}

func (api *StorageHandler) LogOut(w http.ResponseWriter, r *http.Request) {
	var user storage.User
	ParseJSON(r, user)

	err := api.database.DeleteSession(user.Id)
	if err != nil {
		panic(err)
	}

	http.SetCookie(w, &http.Cookie{
		Expires: time.Now().Add(-1 * time.Second),
	})
}
