package hooks

import (
	"fmt"
	"net/http"
)

func HandlePushEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	fmt.Fprintf(w, `{ "message": "Push Event Detected" }`+"\n")
}

func HandleDeploymentEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	fmt.Fprintf(w, `{ "message": "Deployment Event Detected" }`+"\n")
}

func HandlePullRequestEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	fmt.Fprintf(w, `{ "message": "Pull Request Event Detected" }`+"\n")
}
