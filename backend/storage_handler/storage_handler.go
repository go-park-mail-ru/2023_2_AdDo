package storage_handler

import (
	"encoding/json"
	"fmt"
	"main/storage"
	"net/http"
	"time"
)

type StorageHandler struct {
	database *storage.DummyDB
}

func NewStorageHandler() *StorageHandler {
	return &StorageHandler{database: storage.NewDummyDB()}
}

func (api *StorageHandler) Root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "index.html here\n")
}

func (api *StorageHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var user storage.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := api.database.CreateUser(user.Username, user.Password)

	body := map[string]interface{}{
		"id": id,
	}

	json.NewEncoder(w).Encode(&body)
}

func (api *StorageHandler) Auth(w http.ResponseWriter, r *http.Request) {
	var user storage.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userDb, isUser := api.database.GetUser(user.Username)

	if !isUser {
		http.Error(w, err.Error(), 401)
		return
	}

	sessionId := api.database.CreateNewSession(userDb.Id)

	http.SetCookie(w, &http.Cookie{
		Name:    "JSESSIONID",
		Value:   sessionId,
		Expires: time.Now().Add(10 * time.Hour),
	})
}

func (api *StorageHandler) LogOut(w http.ResponseWriter, r *http.Request) {

}
