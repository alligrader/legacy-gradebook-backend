package main

import (
	"fmt"
	"net/http"

	"github.com/gradeshaman/gradebook-backend/oauth"
	. "github.com/gradeshaman/gradebook-backend/tasks"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/markbates/goth/gothic"
)

func main() {

	SendTask()
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)

	s := r.PathPrefix("/auth").Subrouter()

	s.HandleFunc("/{provider}", gothic.BeginAuthHandler)
	s.HandleFunc("/{provider}/callback", oauth.AuthCallback)

	http.Handle("/", r)
	log.Println("Running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, GradeShaman!")
}
