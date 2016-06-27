package hooks

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alligrader/gradebook-backend/tasks"

	log "github.com/Sirupsen/logrus"
	"github.com/google/go-github/github"
)

func HandlePushEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	fmt.Fprintf(w, `{ "message": "Push Event Detected" }`+"\n")

	var requestBody github.Event = github.Event{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		log.Error(err)
	}

	var message github.PushEvent

	message = requestBody.Payload().(github.PushEvent)

	tasks.Bus.PushCheckstyle(*message.PushID)
	tasks.Bus.PushFindbugs(*message.PushID)
}

func HandleDeploymentEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	fmt.Fprintf(w, `{ "message": "Deployment Event Detected" }`+"\n")
}

func HandlePullRequestEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	fmt.Fprintf(w, `{ "message": "Pull Request Event Detected" }`+"\n")
}
