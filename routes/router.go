package routes

import (
	"github.com/gorilla/mux"

	"github.com/alligrader/gradebook-backend/routes/hooks"
)

var R *mux.Router

func init() {
	R = mux.NewRouter().PathPrefix("/api").Subrouter()

	R.HandleFunc("/classes/{id:[0-9]+}/users", HandleAddUserToClass).Methods("POST")
	R.HandleFunc("/classes", HandlePostClasses).Methods("POST")
	R.HandleFunc("/classes/{id:[0-9]+}", HandleGetClassByID).Methods("GET")

	R.HandleFunc("/organizations", HandlePostOrganizations).Methods("POST")
	R.HandleFunc("/organizations/{id:[0-9]+}", HandleGetOrganizationByID).Methods("GET")
	R.HandleFunc("/organizations/{id:[0-9]+}/teachers", HandleAddNewTeachers).Methods("POST")
	R.HandleFunc("/organizations/{id:[0-9]+}/billing", HandleAddBilling).Methods("POST")
	R.HandleFunc("/organizations/{id:[0-9]+}/billing", HandleGetBilling).Methods("GET")
	R.HandleFunc("/organizations/{id:[0-9]+}/billing", HandleDelBilling).Methods("DELETE")

	R.HandleFunc("/organizations/{id:[0-9]+}", HandleDelOrganization).Methods("DELETE")
	R.HandleFunc("/organizations/{id:[0-9]+}/classes/{id2:[0-9]+}", HandleDelClass).Methods("DELETE")

	R.HandleFunc("/hooks", hooks.HandlePushEvent).Headers("X-GitHub-Event", "PushEvent").Methods("POST")
	R.HandleFunc("/hooks", hooks.HandleDeploymentEvent).Headers("X-GitHub-Event", "DeploymentEvent").Methods("POST")
	R.HandleFunc("/hooks", hooks.HandlePullRequestEvent).Headers("X-GitHub-Event", "PullRequestEvent").Methods("POST")
}
