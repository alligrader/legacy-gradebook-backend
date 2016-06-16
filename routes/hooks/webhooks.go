package hooks

import (
	"fmt"
	"net/http"
)

func HandlePushEvent(w http.ResponseWriter, r *http.Request) {

}

func HandleDeploymentEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	fmt.Fprintf(w, `{ "message": "Deployment Event Detected" }`+"\n")
}

func HandlePullRequestEvent(w http.ResponseWriter, r *http.Request) {

}
