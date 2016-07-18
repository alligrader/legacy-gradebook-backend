package spec

import (
	"fmt"
	"testing"

	_ "github.com/alligrader/gradebook-backend/db"
	"github.com/alligrader/gradebook-backend/db/access"
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

func teacherPrivileges(t *testing.T) {

	privileges := []access.Permission{
		access.InviteStudent,
		access.KickStudent,
		access.ReadStudents,
		access.CreateAssignment,
		access.ReadAssignment,
		access.UpdateAssignment,
		access.DeleteAssignment,
		access.GradeSubmission,
		access.UpdateSubmissionGrade,
		access.DeleteSubmissionGrade,
		access.CreateCourse,
		access.UpdateCourse,
		access.DeleteCourse,
	}

	roleToTest := []access.Role{access.Teacher}
	for _, priv := range privileges {
		if !access.AnyGranted(roleToTest, priv, nil) {
			t.Errorf("Missing privilege %v", priv)
		}
	}
}

func failOnFalse(failed bool, t *testing.T, message string) {
	if failed {
		t.Error(message)
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
