package spec

import (
	"testing"

	. "github.com/alligrader/gradebook-backend/db"
	. "github.com/alligrader/gradebook-backend/models"
	"github.com/alligrader/gradebook-backend/util"
)

func TestCreateStudent(t *testing.T) {
	util.WithCleanDB(func() {
		var (
			student *Student = &Student{
				Person: Person{
					FirstName: "Robbie",
					LastName:  "McKinstry",
					Username:  "thesnowmancometh",
					Password:  []byte("0000"),
				},
			}
		)

		err := StudentStore.Create(student)
		if err != nil {
			t.Fatal(err)
		}

		if student.ID == 0 {
			t.Fatal("Failed to set a new ID for a created student")
		}
		if student.Person.ID == 0 {
			t.Fatal("Failed to set a new ID for a created person")
		}

		observedStudent, err := StudentStore.GetByID(student.ID)
		if err != nil {
			t.Fatal(err)
		}

		if !student.Equals(observedStudent) {
			t.Fatal("Observed student did not match the original person.")
		}
	})

}

func TestUpdateStudent(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Create")
	}
}

func TestDeleteStudent(t *testing.T) {
}

func TestSelectStudent(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Create")
	}
}
