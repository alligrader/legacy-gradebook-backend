package oauth

import (
	"crypto/subtle"
	"errors"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/alligrader/gradebook-backend/util"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

var state_hash []byte

func init() {
	util.Configure()
	util.ConfigureLogger()

	var (
		githubKey         = viper.GetString("AUTH_CLIENT_ID")
		githubSecret      = viper.GetString("AUTH_CLIENT_SECRET")
		githubPermissions = []string{"user:email", "repo", "admin:repo_hook", "admin:org_hook", "admin:org"}
		stateTime         = time.Now().UnixNano()
		hash              = strconv.Itoa(int(stateTime))
		err               error
	)

	goth.UseProviders(github.New(githubKey, githubSecret, "", githubPermissions[0], githubPermissions[1], githubPermissions[2], githubPermissions[3], githubPermissions[4]))

	generateFromPassword(err, hash)
	setState()

	gothic.Store = sessions.NewFilesystemStore(os.TempDir(), state_hash)
	getProviderName()
}

func generateFromPassword(err error, hash string) {
	state_hash, err = bcrypt.GenerateFromPassword([]byte(hash), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
}

func setState() {
	gothic.SetState = func(req *http.Request) string {
		return string(state_hash)
	}
}

func getProviderName() {
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

func AuthCallback(w http.ResponseWriter, r *http.Request) {

	observedState := []byte(gothic.GetState(r))
	expectedState := state_hash

	if subtle.ConstantTimeCompare(observedState, expectedState) != 1 {
		http.Error(w, "State sent did not match state received.", http.StatusBadRequest)
		log.Info("Observed and expected states do not match.")
		return
	}

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		log.Warn(w, err)
		return
	}

	t, err := template.ParseFiles("oauth/templates/user.html.tmpl")
	if err != nil {
		log.Warn(w, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, user)
}
