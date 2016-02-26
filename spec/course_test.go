package spec

import (
	"testing"

	. "github.com/gradeshaman/gradebook-backend/db"
	. "github.com/gradeshaman/gradebook-backend/models"
	"github.com/gradeshaman/gradebook-backend/util"
)

func TestPing() {
	conn := util.ConnectToDB()
	defer conn.Close()
	if !conn.Ping() {
		t.Error("Failed to ping the database")
	}
}

func TestCreateCourse(t *testing.T) {
	conn := util.ConnectToDB()
	defer conn.Close()
	util.CreateEmptyDB(conn)
	defer util.CleanDB(conn)

	course := &Course{}
	var courseStore db.CourseStore = &db.CourseMaker{db}

	if !courseStore.Ping() {
		t.Error("Did not connect to database.")
	}

	if err := courseStore.CreateCourse(course); err != nil {
		t.Error(err)
	}
	if course.ID < 1 {
		t.Error("Failed to generate a new ID for the course.")
	}
	if course.CreatedAt.IsZero() {
		t.Error("Created At was not initialized.")
	}
	if course.LastUpdated.IsZero() {
		t.Error("Last Updated was not initialized.")
	}
}

func TestUpdateCourse(t *testing.T) {
}

func TestDeleteCourse(t *testing.T) {
}

func TestSelectCourse(t *testing.T) {
}
