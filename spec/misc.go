package spec

import (
	"testing"

	. "github.com/alligrader/gradebook-backend/util"
)

func TestCanReachDB(t *testing.T) {
	config := GetDBConfigFromEnv()
	db := config.ConnectToDB()
	if err := db.Ping(); err != nil {
		t.Error(err)
	}
}
