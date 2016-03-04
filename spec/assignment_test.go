package spec

import (
	"testing"

	. "github.com/gradeshaman/gradebook-backend/db"
	"github.com/gradeshaman/gradebook-backend/models"
	. "github.com/gradeshaman/gradebook-backend/util"
)

func TestCanReachDB(t *testing.T) {
	config := GetDBConfigFromEnv()
	db := config.ConnectToDB()
	if err := db.Ping(); err != nil {
		t.Error(err)
	}
}

func TestCreateAssignment(t *testing.T) {

	WithCleanDB(func() {
		config := GetDBConfigFromEnv()
		db := config.ConnectToDB()

		assignment := &models.Assignment{}
		var assignmentStore AssignmentStore = &AssignmentMaker{db}
		assignmentStore.CreateAssignment(assignment)

		if assignment.ID == 0 {
			t.Error("Failed to set a new ID for a created assignment")
		}
		observedAssignment, err := assignmentStore.GetAssignmentByID(assignment.ID)
		if err != nil {
			t.Error(err)
		}
		if !assignment.Equals(observedAssignment) {
			t.Error("Observed assignment did not match the original assignment.")
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
