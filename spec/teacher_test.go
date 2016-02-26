package spec

import (
	"testing"

	. "github.com/gradeshaman/gradebook-backend/db"
	. "github.com/gradeshaman/gradebook-backend/models"
	"github.com/gradeshaman/gradebook-backend/util"
)

func TestCreateTeacher(t *testing.T) {
	conn := util.ConnectToDB()
	defer conn.Close()
	util.CreateEmptyDB(conn)
	defer util.CleanDB(conn)

	teacher := &teacher{CourseID: 5}
	var teacherStore db.TeacherStore = &db.TeacherMaker{db}

	if !teacherStore.Ping() {
		t.Error("Did not connect to database.")
	}

	if err := teacherStore.CreateTeacher(teacher); err != nil {
		t.Error(err)
	}
	if teacher.ID < 1 {
		t.Error("Failed to generate a new ID for the course.")
	}
	if teacher.CreatedAt.IsZero() {
		t.Error("Created At was not initialized.")
	}
	if teacher.LastUpdated.IsZero() {
		t.Error("Last Updated was not initialized.")
	}
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
