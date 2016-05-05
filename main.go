package main

import (
	"crypto/subtle"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/alligrader/gradebook-backend/tasks"
	"github.com/alligrader/gradebook-backend/util"
	"github.com/alligrader/gradebook-backend/ziphandler"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

var bits []byte

func init() {
	util.Configure()
	util.ConfigureLogger()
	githubKey := viper.GetString("AUTH_CLIENT_ID")
	githubSecret := viper.GetString("AUTH_CLIENT_SECRET")
	goth.UseProviders(github.New(githubKey, githubSecret, "", "user:email", "repo", "admin:repo_hook", "admin:org_hook", "admin:org"))

	t := time.Now().UnixNano()
	hash := strconv.Itoa(int(t))
	var err error

	bits, err = bcrypt.GenerateFromPassword([]byte(hash), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	gothic.SetState = func(req *http.Request) string {
		return string(bits)
	}

	gothic.Store = sessions.NewFilesystemStore(os.TempDir(), bits)
	gothic.GetProviderName = func(r *http.Request) (string, error) {
		vars := mux.Vars(r)
		provider := vars["provider"]

		if provider == "" {
			provider = vars[":provider"]
		}
		if provider == "" {
			return provider, errors.New("you must select a provider")
		}
		return provider, nil
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/auth/{provider}", gothic.BeginAuthHandler)
	r.HandleFunc("/auth/{provider}/callback", AuthCallback)
	r.HandleFunc("/zip", ziphandler.MockHandler)
	r.HandleFunc("/zip/upload", ziphandler.HandleZipUpload) // TODO remove the stutter
	http.Handle("/", r)
	log.Println("Running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, GradeShaman!")
}

func AuthCallback(w http.ResponseWriter, r *http.Request) {

	observedState := []byte(gothic.GetState(r))
	expectedState := bits

	if subtle.ConstantTimeCompare(observedState, expectedState) != 1 {
		http.Error(w, "State sent did not match state received.", 400)
		log.Error(string(observedState))
		log.Error(string(expectedState))
		return
	}

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	t, _ := template.New("foo").Parse(userTemplate)
	t.Execute(w, user)
}

var userTemplate = `
<p>Name: {{.Name}}</p>
<p>Email: {{.Email}}</p>
<p>NickName: {{.NickName}}</p>
<p>Location: {{.Location}}</p>
<p>AvatarURL: {{.AvatarURL}} <img src="{{.AvatarURL}}"></img></p>
<p>Description: {{.Description}}</p>
<p>UserID: {{.UserID}}</p>
<p>AccessToken: {{.AccessToken}}</p>
`
