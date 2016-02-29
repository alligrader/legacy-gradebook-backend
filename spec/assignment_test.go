package spec

import (
	"testing"

	"github.com/gradeshaman/gradebook-backend/db"
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
	if testing.Short() {
		t.Skip("Testing dependent on Select")
	}
	config := GetDBConfigFromEnv()
	db := config.ConnectToDB()
	course := &models.Assignment{}
	var courseStore CourseStore = &db.CourseMaker{db}
	courseStore.CreateCourse(course)

	if course.ID == 0 {
		t.Error("Failed to set a new ID for a created course")
	}
	observedCourse := courseStore.GetCourseByID(course.ID)
	if !course.Equals(observedCourse) {
		t.Error("Observed course did not make the original course.")
	}
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
