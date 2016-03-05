package spec

import (
	"testing"

	. "github.com/gradeshaman/gradebook-backend/db"
	"github.com/gradeshaman/gradebook-backend/models"
	. "github.com/gradeshaman/gradebook-backend/util"
)

func TestCreateCourse(t *testing.T) {
	WithCleanDB(func() {
		config := GetDBConfigFromEnv()
		db := config.ConnectToDB()

		course := &models.Course{}
		var courseStore CourseStore = &CourseMaker{db}
		err := courseStore.CreateCourse(course)
		if err != nil {
			t.Fatal(err)
		}

		if course.ID == 0 {
			t.Fatal("Failed to set a new ID for a created assignment")
		}

		observedCourse, err := courseStore.GetCourseByID(course.ID)
		if err != nil {
			t.Fatal(err)
		}

		if course == nil {
			t.Fatal("What the fuck?")
		}
		if !course.Equals(observedCourse) {
			t.Fatal("Observed course did not match the original course.")
		}
	})
}

func TestUpdateCourse(t *testing.T) {
}

func TestDeleteCourse(t *testing.T) {
}

func TestSelectCourse(t *testing.T) {
}
