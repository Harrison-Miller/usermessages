package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Env struct {
	Database      *Database
	Authenticator *Authenticator
}

func envHandler(e Env, handler func(Env, http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	h := func(w http.ResponseWriter, r *http.Request) {
		handler(e, w, r)
	}
	return h
}

// /api/message GET
func GetMessage(e Env, w http.ResponseWriter, r *http.Request) {
	email, password, ok := r.BasicAuth()
	var user User
	var err error
	if user, err = e.Authenticator.Login(email, password); err != nil || !ok {
		http.Error(w, "Unathorized", http.StatusUnauthorized)
		return
	}

	m, err := e.Database.GetMessage(user.UserID)
	if err != nil {
		http.Error(w, "No message has been set", http.StatusOK)
		return
	}

	JSONResponse(w, m)
}

// /api/message POST
func SetMessage(e Env, w http.ResponseWriter, r *http.Request) {
	email, password, ok := r.BasicAuth()
	var user User
	var err error
	if user, err = e.Authenticator.Login(email, password); err != nil || !ok {
		http.Error(w, "Unathorized", http.StatusUnauthorized)
		return
	}

	var m Message
	err = json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, "Malformed request", http.StatusBadRequest)
		return
	}

	err = e.Database.SetMessage(user.UserID, m.Message)
	if err != nil {
		http.Error(w, "Unable to write message", http.StatusInternalServerError)
	}
}

func main() {
	var dataDir = "data"
	if v, ok := os.LookupEnv("DATA_DIR"); ok {
		dataDir = v
	}

	// create data dir if it doesn't exist
	err := os.MkdirAll(dataDir, 0755)
	if err != nil {
		panic(err)
	}

	database, err := CreateDatabase(dataDir)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	auth, err := CreateAuthenticator(client, "https://reqres.in/api/users")
	if err != nil {
		panic(err)
	}

	env := Env{
		Database:      database,
		Authenticator: auth,
	}

	r := mux.NewRouter()
	r.HandleFunc("/api/message", envHandler(env, GetMessage)).Methods("GET")
	r.HandleFunc("/api/message", envHandler(env, SetMessage)).Methods("POST")

	fmt.Println("Server running on :8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}

}
