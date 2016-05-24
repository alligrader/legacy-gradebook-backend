package spec

import (
	"testing"

	. "github.com/alligrader/gradebook-backend/db"
	. "github.com/alligrader/gradebook-backend/models"
	"github.com/alligrader/gradebook-backend/util"
)

func TestCreateCourse(t *testing.T) {
	t.Skip()
	util.WithCleanDB(func() {

		var (
			teachersVec []*Teacher = make([]*Teacher, 0, 2)
			studentsVec []*Student = make([]*Student, 0, 2)

			teachers = []Teacher{
				{
					Person: Person{
						FirstName: "Jon",
						LastName:  "Misurda",
						Username:  "Dr. Vape",
						Password:  []byte("0000"),
					},
				}, {
					Person: Person{
						FirstName: "Sebastion",
						LastName:  "Awful",
						Username:  "Bjergsen",
						Password:  []byte("0000"),
					},
				},
			}

			students = []Student{
				{
					Person: Person{
						FirstName: "Neel",
						LastName:  "Kowdley",
						Username:  "3Legs",
						Password:  []byte("0000"),
					},
				}, {
					Person: Person{
						FirstName: "Dave",
						LastName:  "Sweeney",
						Username:  "Swim2Win",
						Password:  []byte("0000"),
					},
				},
			}
		)

		for _, teacher := range teachers {
			if err := TeacherStore.Create(&teacher); err != nil {
				t.Fatal(err)
			}
			if teacher.ID == 0 {
				t.Fatal("No ID on teacher.")
			}
			if teacher.Person.ID == 0 {
				t.Fatal("No ID on teacher.Person")
			}
			teachersVec = append(teachersVec, &teacher)
		}

		for _, student := range students {
			if err := StudentStore.Create(&student); err != nil {
				t.Fatal(err)
			}
			if student.ID == 0 {
				t.Fatal("No ID on student")
			}
			if student.Person.ID == 0 {
				t.Fatal("No ID on student.Person")
			}
			studentsVec = append(studentsVec, &student)
		}

		var course *Course = &Course{
			Name:     "Dr. Misurda's Wild Ride",
			Teachers: teachersVec,
			Students: studentsVec,
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
