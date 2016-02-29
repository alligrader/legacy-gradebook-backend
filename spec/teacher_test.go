package spec

import (
	"testing"

	_ "github.com/gradeshaman/gradebook-backend/db"
	_ "github.com/gradeshaman/gradebook-backend/models"
	_ "github.com/gradeshaman/gradebook-backend/util"
)

func TestCreateTeacher(t *testing.T) {
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
