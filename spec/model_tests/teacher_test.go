package spec

import (
	"testing"

	. "github.com/alligrader/gradebook-backend/db"
	. "github.com/alligrader/gradebook-backend/models"
	"github.com/alligrader/gradebook-backend/util"

	_ "github.com/go-sql-driver/mysql"
)

func TestCreateTeacher(t *testing.T) {
	util.WithCleanDB(func() {
		var (
			teacher *Teacher = &Teacher{
				User: User{
					FirstName: "Mark",
					LastName:  "Krotec",
					Username:  "mckrotec@yahoo.com",
					Password:  []byte("0000"),
				},
			}
		)

		err := TeacherStore.Create(teacher)
		if err != nil {
			t.Fatal(err)
		}

		if teacher.ID == 0 {
			t.Fatal("Failed to set a new ID for a created teacher")
		}
		if teacher.User.ID == 0 {
			t.Fatal("Failed to set a new ID for a created user")
		}

		observedTeacher, err := TeacherStore.GetByID(teacher.ID)
		if err != nil {
			t.Fatal(err)
		}

		if !teacher.Equals(observedTeacher) {
			t.Fatal("Observed student did not match the original user.")
		}
	})
}

func TestUpdateTeacher(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Create")
	}
}

func TestDeleteTeacher(t *testing.T) {
}

func TestSelectTeacher(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Create")
	}
}
