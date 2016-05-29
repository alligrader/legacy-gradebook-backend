package spec

import (
	"fmt"
	"testing"

	. "github.com/alligrader/gradebook-backend/db"
	. "github.com/alligrader/gradebook-backend/models"
	"github.com/alligrader/gradebook-backend/util"
)

func TestCreateCourse(t *testing.T) {
	util.WithCleanDB(func() {

		var (
			teachers = []*Teacher{
				{
					User: User{
						FirstName: "Jon",
						LastName:  "Misurda",
						Username:  "Dr. Vape",
						Password:  []byte("0000"),
					},
				}, {
					User: User{
						FirstName: "Sebastion",
						LastName:  "Awful",
						Username:  "Bjergsen",
						Password:  []byte("0000"),
					},
				},
			}

			students = []*Student{
				{
					User: User{
						FirstName: "Neel",
						LastName:  "Kowdley",
						Username:  "3Legs",
						Password:  []byte("0000"),
					},
				}, {
					User: User{
						FirstName: "Dave",
						LastName:  "Sweeney",
						Username:  "Swim2Win",
						Password:  []byte("0000"),
					},
				},
			}
		)

		for _, teacher := range teachers {
			if err := TeacherStore.Create(teacher); err != nil {
				t.Fatal(err)
			}
			if teacher.ID == 0 {
				t.Fatal("No ID on teacher.")
			}
			if teacher.User.ID == 0 {
				t.Fatal("No ID on teacher.User")
			}
		}

		for _, student := range students {
			if err := StudentStore.Create(student); err != nil {
				t.Fatal(err)
			}
			if student.ID == 0 {
				t.Fatal("No ID on student")
			}
			if student.User.ID == 0 {
				t.Fatal("No ID on student.User")
			}
		}

		var course *Course = &Course{
			Name:     "Dr. Misurda's Wild Ride",
			Teachers: teachers,
			Students: students,
		}

		if err := CourseStore.Create(course); err != nil {
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
			fmt.Printf("Observed Course: %v \n", observedCourse)
			fmt.Printf("Expected Course: %v \n", course)
			t.Fatal("Observed course did not match the original course.")
		}
	})
}

func TestUpdateCourse(t *testing.T) {

	/*
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
	*/
}

func TestDeleteCourse(t *testing.T) {

	/*
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
	*/
}

func TestSelectCourse(t *testing.T) {
}
