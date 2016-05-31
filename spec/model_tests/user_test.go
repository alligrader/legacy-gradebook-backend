package spec

import (
	"testing"

	. "github.com/alligrader/gradebook-backend/db"
	"github.com/alligrader/gradebook-backend/models"
	"github.com/alligrader/gradebook-backend/util"
)

func TestCreateUser(t *testing.T) {

	util.WithCleanDB(func() {
		var (
			user *models.User = &models.User{
				FirstName: "Robbie",
				LastName:  "McKinstry",
				Username:  "thesnowmancometh",
				Password:  []byte("0000"),
			}
		)

		err := UserStore.Create(user)
		if err != nil {
			t.Fatal(err)
		}

		if user.ID == 0 {
			t.Fatal("Failed to set a new ID for a created user")
		}

		observedUser, err := UserStore.GetByID(user.ID)
		if err != nil {
			t.Fatal(err)
		}

		if !user.Equals(observedUser) {
			t.Fatal("Observed user did not match the original user.")
		}
	})

}

func TestUpdateUser(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Create")
	}
}

func TestDeleteUser(t *testing.T) {
}

func TestSelectUser(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Create")
	}
}
