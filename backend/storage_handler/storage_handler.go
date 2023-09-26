package storage_handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"main/handler"
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

func RenderJSON(w http.ResponseWriter, v interface{}) error {
	jsonResponse, err := json.Marshal(v)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
	return nil
}

func (api *StorageHandler) Home(w http.ResponseWriter, r *http.Request) error {
	var user storage.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return handler.StatusError{Code: 409, Err: err}
	}

	sessionId, err := r.Cookie(CookieName)
	if err != nil {
		return handler.StatusError{Code: 409, Err: err}
	}

	isAuth, err := api.database.CheckSession(user.Id, sessionId.Value)
	if err != nil {
		return err
	}

	if isAuth {
		//auth
		fmt.Fprintf(w, "authorized\n")
	} else {
		// no auth
		fmt.Fprintf(w, "unauthorized\n")
	}
	return nil
}

func (api *StorageHandler) SignUp(w http.ResponseWriter, r *http.Request) error {
	var user storage.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return handler.StatusError{Code: 409, Err: err}
	}

	id, err := api.database.CreateUser(user)
	if err != nil {
		return handler.StatusError{Code: 409, Err: err}
	}

	if id == 0 {
		return handler.StatusError{Code: http.StatusConflict, Err: err}
	}

	err = RenderJSON(w, storage.ResponseId{Id: id})
	if err != nil {
		return err
	}
	return nil
}

func (api *StorageHandler) Auth(w http.ResponseWriter, r *http.Request) error {
	var user storage.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return handler.StatusError{Code: 409, Err: err}
	}

	userId, err := api.database.CheckUserCredentials(user)

	if err != nil {
		return handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}

	sessionId, err := api.database.CreateNewSession(userId)
	if err != nil {
		return handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	http.SetCookie(w, &http.Cookie{
		Name:     CookieName,
		Value:    sessionId,
		Expires:  time.Now().Add(1 * time.Minute),
		Secure:   true,
		HttpOnly: true,
	})
	err = RenderJSON(w, storage.ResponseId{Id: userId})
	if err != nil {
		return err
	}
	return nil
}

func (api *StorageHandler) LogOut(w http.ResponseWriter, r *http.Request) error {
	var user storage.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	err = api.database.DeleteSession(user.Id)
	if err != nil {
		return handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	http.SetCookie(w, &http.Cookie{
		Expires: time.Now().Add(-1 * time.Second),
	})
	return nil
}
