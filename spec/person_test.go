package spec

import (
	"testing"

	. "github.com/alligrader/gradebook-backend/db"
	"github.com/alligrader/gradebook-backend/models"
	"github.com/alligrader/gradebook-backend/util"
)

func TestCreatePerson(t *testing.T) {

	util.WithCleanDB(func() {
		var (
			person *models.Person = &models.Person{
				FirstName: "Robbie",
				LastName:  "McKinstry",
				Username:  "thesnowmancometh",
				Password:  []byte("0000"),
			}
		)

		err := PersonStore.Create(person)
		if err != nil {
			t.Fatal(err)
		}

		if person.ID == 0 {
			t.Fatal("Failed to set a new ID for a created person")
		}

		observedPerson, err := PersonStore.GetByID(person.ID)
		if err != nil {
			t.Fatal(err)
		}

		if !person.Equals(observedPerson) {
			t.Fatal("Observed person did not match the original person.")
		}
	})

}

func TestUpdatePerson(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Create")
	}
}

func TestDeletePerson(t *testing.T) {
}

func TestSelectPerson(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Create")
	}
}
