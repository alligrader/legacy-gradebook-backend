package routes

import (
	"github.com/gorilla/mux"

	"github.com/alligrader/gradebook-backend/routes/hooks"
)

var R *mux.Router

func init() {
	R = mux.NewRouter()
	s := R.PathPrefix("/api").Subrouter()

	s.HandleFunc("/classes/{id:[0-9]+}/users", HandleAddUserToClass).Methods("POST")
	s.HandleFunc("/classes", HandlePostClasses).Methods("POST")
	s.HandleFunc("/classes/{id:[0-9]+}", HandleGetClassByID).Methods("GET")

	s.HandleFunc("/organizations", HandlePostOrganizations).Methods("POST")
	s.HandleFunc("/organizations/{id:[0-9]+}", HandleGetOrganizationByID).Methods("GET")
	s.HandleFunc("/organizations/{id:[0-9]+}/teachers", HandleAddNewTeachers).Methods("POST")
	s.HandleFunc("/organizations/{id:[0-9]+}/billing", HandleAddBilling).Methods("POST")
	s.HandleFunc("/organizations/{id:[0-9]+}/billing", HandleGetBilling).Methods("GET")
	s.HandleFunc("/organizations/{id:[0-9]+}/billing", HandleDelBilling).Methods("DELETE")

	s.HandleFunc("/organizations/{id:[0-9]+}", HandleDelOrganization).Methods("DELETE")
	s.HandleFunc("/organizations/{id:[0-9]+}/classes/{id2:[0-9]+}", HandleDelClass).Methods("DELETE")

	s.HandleFunc("/hooks", hooks.HandlePushEvent).Headers("X-GitHub-Event", "PushEvent").Methods("POST")
	s.HandleFunc("/hooks", hooks.HandleDeploymentEvent).Headers("X-GitHub-Event", "DeploymentEvent").Methods("POST")
	s.HandleFunc("/hooks", hooks.HandlePullRequestEvent).Headers("X-GitHub-Event", "PullRequestEvent").Methods("POST")
}
