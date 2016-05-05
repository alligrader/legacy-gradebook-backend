package spec

import (
	"testing"

	_ "github.com/alligrader/gradebook-backend/db"
	_ "github.com/alligrader/gradebook-backend/models"
	_ "github.com/alligrader/gradebook-backend/util"
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
