package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gradeshaman/gradebook-backend/util"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

func main() {

	util.Configure()
	util.ConfigureLogger()

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, GradeShaman!")
}
