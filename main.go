package main

import (
	"fmt"
	"net/http"

	"github.com/alligrader/gradebook-backend/oauth"
	_ "github.com/alligrader/gradebook-backend/tasks"
	_ "github.com/alligrader/gradebook-backend/util"
	"github.com/alligrader/gradebook-backend/ziphandler"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/markbates/goth/gothic"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", HelloHandler)
	r.HandleFunc("/zip", ziphandler.MockHandler)
	r.HandleFunc("/zip/upload", ziphandler.HandleZipUpload) // TODO remove the stutter

	r.HandleFunc("/users", CreateUser).Methods("POST")

	s := r.PathPrefix("/auth").Subrouter()

	s.HandleFunc("/{provider}", gothic.BeginAuthHandler)
	s.HandleFunc("/{provider}/callback", oauth.AuthCallback)

	http.Handle("/", r)
	log.Println("Running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Alligrader!\n")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	fmt.Fprintf(w, `{ "message": "Hello New User!" }`+"\n")
}
