package main

import (
	"fmt"
	"net/http"

	"github.com/gradeshaman/gradebook-backend/util"

	log "github.com/Sirupsen/logrus"
	_ "github.com/gorilla/mux"
)

func main() {

	util.Configure()
	util.ConfigureLogger()

	pingDatabase()
	/*
		r := mux.NewRouter()
		r.HandleFunc("/", HomeHandler)
		http.Handle("/", r)
		log.Println("Running on port 8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	*/
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, GradeShaman!")
}

func pingDatabase() {
	db := util.ConnectToDB()
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected!")
	}
}
