package main

import (
	"fmt"
	"net/http"

	. "github.com/gradeshaman/gradebook-backend/tasks"
	"github.com/gradeshaman/gradebook-backend/util"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

func main() {

	SendTask()

	util.Configure()
	util.ConfigureLogger()

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)
	log.Println("Running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, GradeShaman!")
}
