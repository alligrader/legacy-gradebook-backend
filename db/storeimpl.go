package db

import (
	. "github.com/gradeshaman/gradebook-backend/models"
	"github.com/jmoiron/sqlx"
)

type CourseMaker struct {
	*sqlx.DB
}

func (maker *CourseMaker) CreateCourse(class *Course) error {
	return nil
}

func (maker *CourseMaker) UpdateCourse(class *Course) error {

	return nil
}

func (maker *CourseMaker) GetCourseByID(id int) (*Course, error) {
	return nil, nil

}
func (maker *CourseMaker) DestroyCourse(class *Course) error {

	return nil
}

type TeacherMaker struct {
	*sqlx.DB
}

func (maker *TeacherMaker) CreateTeacher(teacher *Teacher) error {
	return nil

}
func (maker *TeacherMaker) UpdateTeacher(teacher *Teacher) error {
	return nil

}
func (maker *TeacherMaker) GetTeacherByID(id int) (*Teacher, error) {
	return nil, nil

}

func (maker *TeacherMaker) DestroyTeacher(t *Teacher) error {
	return nil
}

type StudentMaker struct {
	*sqlx.DB
}

func (maker *StudentMaker) CreateStudent(student *Student) error {
	return nil

}
func (maker *StudentMaker) UpdateStudent(student *Student) error {
	return nil

}
func (maker *StudentMaker) GetStudentByID(id int) (*Student, error) {
	return nil, nil

}
func (maker *StudentMaker) DestroyStudent(student *Student) error {
	return nil
}

type UserMaker struct {
	*sqlx.DB
}

func (maker *UserMaker) CreateUser(user *User) error {
	return nil

}
func (maker *UserMaker) UpdateUser(user *User) error {
	return nil

}
func (maker *UserMaker) GetUserByID(id int) (*User, error) {
	return nil, nil

}
func (maker *UserMaker) DestroyUser(user *User) error {
	return nil

}

type TestScoreMaker struct {
	*sqlx.DB
}

func (maker *TestScoreMaker) CreateTestStore(score *Test) error {
	return nil

}
func (maker *TestScoreMaker) UpdateTestScore(score *Test) error {
	return nil

}
func (maker *TestScoreMaker) GetTestScoreByID(id int) (*Test, error) {
	return nil, nil

}
func (maker *TestScoreMaker) DestroyUser(score *Test) error {
	return nil
}

type TestResultMaker struct {
	*sqlx.DB
}

func (maker *TestResultMaker) CreateTestResult(result *TestResult) error {
	return nil

}
func (maker *TestResultMaker) UpdateTestResult(result *TestResult) error {
	return nil

}
func (maker *TestResultMaker) GetTestResultByID(id int) (*TestResult, error) {
	return nil, nil

}
func (maker *TestResultMaker) DestroyTestResult(result *TestResult) error {
	return nil
}

type RunResultMaker struct {
	*sqlx.DB
}

func (maker *RunResultMaker) CreateRunResult(result *RunResult) error {
	return nil

}
func (maker *RunResultMaker) UpdateRunResult(result *RunResult) error {
	return nil

}
func (maker *RunResultMaker) GetRunResultByID(id int) (*RunResult, error) {
	return nil, nil

}
func (maker *RunResultMaker) DestroyRunResult(result *RunResult) error {
	return nil
}

type AssignmentMaker struct {
	*sqlx.DB
}

func (maker *AssignmentMaker) CreateAssignment(assig *Assignment) error {
	return nil
}

func (maker *AssignmentMaker) UpdateAssignment(assig *Assignment) error {
	return nil

}

func (maker *AssignmentMaker) GetAssignmentByID(id int) (*Assignment, error) {
	return nil, nil
}

func (maker *AssignmentMaker) DestroyAssignment(assig *Assignment) error {
	return nil
}
