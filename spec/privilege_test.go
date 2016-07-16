package spec

import (
	"fmt"
	"testing"

	_ "github.com/alligrader/gradebook-backend/db"
	. "github.com/alligrader/gradebook-backend/models"
	_ "github.com/alligrader/gradebook-backend/util"
)

func TestRolesHavePrivileges(t *testing.T) {
	// Fetch the privileges for each role
	// And make sure that they match.

	roles := map[string]func(*testing.T){
		"teacher":            teacherPrivileges,
		"admin":              adminPrivileges,
		"student":            studentPrivileges,
		"teaching_assistant": teachingAssistantPrivileges,
		"org_admin":          orgAdminPrivileges,
	}

	for key, val := range roles {
		name := fmt.Sprintf("name=%s", key)
		t.Run(name, val)
	}
}

func failOnFalse(failed bool, t *testing.T, message string) {
	if failed {
		t.Error(message)
	}
}

func teacherPrivileges(t *testing.T) {

	t.Skip()
	table := []struct {
		role    string
		action  string
		status  int
		message string
	}{
		{"teacher", "invite_student", 0, "Teacher cannot invite student"},
		{"teacher", "remove_student", 0, "Teacher cannot remove student"},
		{"teacher", "read_student", 0, "Teacher cannot read student"},
		{"teacher", "create_assignment", 0, "Teacher cannot create assignment"},
		{"teacher", "read_assignment", 0, "Teacher cannot read assignment"},
		{"teacher", "update_assignment", 0, "Teacher cannot update assignment"},
		{"teacher", "update_submission_grade", 0, "Teacher cannot update submission grade"},
		{"teacher", "read_submission", 0, "Teacher cannot read submission"},
		{"teacher", "delete_submission_grade", 0, "Teacher cannot delete submission grade"},
		{"teacher", "create_class", 0, "Teacher cannot create class"},
		{"teacher", "read_class", 0, "Teacher cannot read class"},
		{"teacher", "update_class", 0, "Teacher cannot update class"},
		{"teacher", "delete_class", 0, "Teacher cannot delete class"},
	}

	for _, elem := range table {
		r := Role{Name: elem.role}
		failOnFalse(!r.Can(Action(elem.action), Status(elem.status)), t, elem.message)
	}
}

func adminPrivileges(t *testing.T) {
	t.Skip()
}

func studentPrivileges(t *testing.T) {
	t.Skip()
}

func teachingAssistantPrivileges(t *testing.T) {
	t.Skip()
}

func orgAdminPrivileges(t *testing.T) {
	t.Skip()
}

func TestAggregateRoles(t *testing.T) {
	t.Skip()
	// Merge two roles together
	// Make sure they contain all of the right privileges
	// Make sure they do not contain the wrong privileges
}
