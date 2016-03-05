package spec

import (
	"testing"

	. "github.com/gradeshaman/gradebook-backend/db"
	"github.com/gradeshaman/gradebook-backend/models"
	. "github.com/gradeshaman/gradebook-backend/util"
)

func TestCreateAssignment(t *testing.T) {

	t.Skip("Dependent on course impl")
	WithCleanDB(func() {
		config := GetDBConfigFromEnv()
		db := config.ConnectToDB()

		assignment := &models.Assignment{}
		var assignmentStore AssignmentStore = &AssignmentMaker{db}
		err := assignmentStore.CreateAssignment(assignment)
		if err != nil {
			t.Fatal(err)
		}

		if assignment.ID == 0 {
			t.Fatal("Failed to set a new ID for a created assignment")
		}

		observedAssignment, err := assignmentStore.GetAssignmentByID(assignment.ID)
		if err != nil {
			t.Fatal(err)
		}

		if !assignment.Equals(observedAssignment) {
			t.Fatal("Observed assignment did not match the original assignment.")
		}
	})
}

func TestUpdateAssignment(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Create")
	}
}

func TestDeleteAssignment(t *testing.T) {
}

func TestSelectAssignment(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Create")
	}
}
