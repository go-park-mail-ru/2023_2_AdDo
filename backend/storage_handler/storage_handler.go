package storage_handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
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

func (api *StorageHandler) Root(w http.ResponseWriter, r *http.Request) {
	var user storage.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if api.database.CheckUserCredentials(user) {
		fmt.Fprintf(w, "authorized\n")
	} else {
		fmt.Fprintf(w, "unauthorized\n")
	}

}

func (api *StorageHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var user storage.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := api.database.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	body := map[string]interface{}{
		"id": id,
	}

	err = json.NewEncoder(w).Encode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (api *StorageHandler) Auth(w http.ResponseWriter, r *http.Request) {
	var user storage.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	isUser := api.database.CheckUserCredentials(user)

	if !isUser {
		http.Error(w, err.Error(), 401)
		return
	}

	sessionId, err := api.database.CreateNewSession(user.Id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "JSESSIONID",
		Value:   sessionId,
		Expires: time.Now().Add(10 * time.Hour),
	})
}

func (api *StorageHandler) LogOut(w http.ResponseWriter, r *http.Request) {
	var user storage.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	api.database.DeleteSession(user.Id)
}
