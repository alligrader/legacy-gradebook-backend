package spec

import (
	"testing"

	. "github.com/gradeshaman/gradebook-backend/db"
	"github.com/gradeshaman/gradebook-backend/models"
	"github.com/gradeshaman/gradebook-backend/util"
	"github.com/jmoiron/sqlx"
)

func TestCreateCourse(t *testing.T) {
	util.WithCleanDB(func() {

		var (
			config      *util.DBConfig = util.GetDBConfigFromEnv()
			db          *sqlx.DB       = config.ConnectToDB()
			course      *models.Course = &models.Course{Name: "Dr. Misurda's Wild Ride"}
			courseStore CourseStore    = &CourseMaker{db}
		)

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

		if !course.Equals(observedCourse) {
			t.Fatal("Observed course did not match the original course.")
		}
	})
}

func TestUpdateCourse(t *testing.T) {
	util.WithCleanDB(func() {

		var (
			config      *util.DBConfig = util.GetDBConfigFromEnv()
			db          *sqlx.DB       = config.ConnectToDB()
			course      *models.Course = &models.Course{}
			courseStore CourseStore    = &CourseMaker{db}
		)

		_ = courseStore.CreateCourse(course)
		course.Name = "Advanced Topics in Static Analysis"
		err := courseStore.UpdateCourse(course)

		if err != nil {
			t.Fatal(err)
		}

		observedCourse, err := courseStore.GetCourseByID(course.ID)
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
			config      *util.DBConfig = util.GetDBConfigFromEnv()
			db          *sqlx.DB       = config.ConnectToDB()
			course      *models.Course = &models.Course{}
			courseStore CourseStore    = &CourseMaker{db}
		)

		_ = courseStore.CreateCourse(course)
		err := courseStore.DestroyCourse(course)
		if err != nil {
			t.Fatal(err)
		}

		_, err = courseStore.GetCourseByID(course.ID)
		if err == nil {
			t.Fatal("No error, when an error is expected.")
		}
	})
}

func TestSelectCourse(t *testing.T) {
}
