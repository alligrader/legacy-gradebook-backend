package hooks

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/alligrader/gradebook-backend/tasks"

	log "github.com/Sirupsen/logrus"
	"github.com/google/go-github/github"
)

func HandlePushEvent(w http.ResponseWriter, r *http.Request) {

	var (
		message github.PushEvent = github.PushEvent{}
		// ghEvent                  = &github.Event{}
		err = json.NewDecoder(r.Body).Decode(&message)
	)

	if err != nil {
		log.Error(err)
		return
	}

	// fmt.Printf("%v\n", ghEvent)
	// fmt.Println(ghEvent.Type)
	//i := ghEvent.Payload()

	// message = i.(github.PushEvent)

	// tasks.Bus.PushCheckstyle(*(message.PushID))
	// tasks.Bus.PushFindbugs(*(message.PushID))
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
