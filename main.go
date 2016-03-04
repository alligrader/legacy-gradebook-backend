package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"

	. "github.com/gradeshaman/gradebook-backend/tasks"
	"github.com/gradeshaman/gradebook-backend/util"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
	"github.com/spf13/viper"
)

func init() {

	gothic.Store = sessions.NewFilesystemStore(os.TempDir(), []byte("goth-example"))
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

	SendTask()

	util.Configure()
	util.ConfigureLogger()

	githubKey := viper.GetString("AUTH_CLIENT_ID")
	githubSecret := viper.GetString("AUTH_CLIENT_SECRET")
	goth.UseProviders(github.New(githubKey, githubSecret, ""))

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/auth/{provider}", auth) // gothic.BeginAuthHandler)
	r.HandleFunc("/auth/{provider}/callback", AuthCallback)
	http.Handle("/", r)
	log.Println("Running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, GradeShaman!")
}

func auth(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	provider := vars["provider"]
	log.Warn(provider)
	gothic.BeginAuthHandler(w, r)
}

func AuthCallback(w http.ResponseWriter, r *http.Request) {

	fmt.Println("State: ", gothic.GetState(r))

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
