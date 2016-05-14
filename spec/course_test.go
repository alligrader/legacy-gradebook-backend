package spec

import (
	"testing"

	. "github.com/alligrader/gradebook-backend/db"
	"github.com/alligrader/gradebook-backend/models"
	"github.com/alligrader/gradebook-backend/util"
)

func TestCreateCourse(t *testing.T) {
	util.WithCleanDB(func() {

		var (
			course *models.Course = &models.Course{Name: "Dr. Misurda's Wild Ride"}
		)

		err := CourseStore.Create(course)
		if err != nil {
			t.Fatal(err)
		}

		if course.ID == 0 {
			t.Fatal("Failed to set a new ID for a created assignment")
		}

		observedCourse, err := CourseStore.GetByID(course.ID)
		if err != nil {
			t.Fatal(err)
		}

		if !course.Equals(observedCourse) {
			t.Fatal("Observed course did not match the original course.")
		}
	})
}

func TestUpdateCourse(t *testing.T) {
	util.WithCleanDB(func() {

		var (
			course *models.Course = &models.Course{}
		)

		_ = CourseStore.Create(course)
		course.Name = "Advanced Topics in Static Analysis"
		err := CourseStore.UpdateCourse(course)

		if err != nil {
			t.Fatal(err)
		}

		observedCourse, err := CourseStore.GetByID(course.ID)
		if err != nil {
			t.Fatal(err)
		}

		if !course.Equals(observedCourse) {
			t.Fatal("The observed course does not match what was passed into the database.")
		}
		if observedCourse.Name != "Advanced Topics in Static Analysis" {
			t.Fatal("The object in the database was not updated with a new name.")
		}
	})
}

func TestDeleteCourse(t *testing.T) {
	util.WithCleanDB(func() {

		var (
			course *models.Course = &models.Course{}
		)

		_ = CourseStore.Create(course)
		err := CourseStore.Destroy(course)
		if err != nil {
			t.Fatal(err)
		}

		_, err = CourseStore.GetByID(course.ID)
		if err == nil {
			t.Fatal("No error, when an error is expected.")
		}
	})
}

func TestSelectCourse(t *testing.T) {
}
