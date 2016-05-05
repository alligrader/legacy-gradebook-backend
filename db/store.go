package db

import (
	. "github.com/alligrader/gradebook-backend/models"
)

type CourseStore interface {
	Ping() error
	CreateCourse(class *Course) error
	UpdateCourse(class *Course) error
	GetCourseByID(id int) (*Course, error)
	DestroyCourse(class *Course) error
}

type TeacherStore interface {
	Ping() error
	CreateTeacher(teacher *Teacher) error
	UpdateTeacher(teacher *Teacher) error
	GetTeacherByID(id int) (*Teacher, error)
	DestroyTeacher(t *Teacher) error
}

type StudentStore interface {
	Ping() error
	CreateStudent(student *Student) error
	UpdateStudent(student *Student) error
	GetStudentByID(id int) (*Student, error)
	DestroyStudent(student *Student) error
}
type UserStore interface {
	Ping() error
	CreateUser(user *User) error
	UpdateUser(user *User) error
	GetUserByID(id int) (*User, error)
	DestroyUser(user *User) error
}
type TestStore interface {
	Ping() error
	CreateTestStore(score *Test) error
	UpdateTestScore(score *Test) error
	GetTestScoreByID(id int) (*Test, error)
	DestroyUser(score *Test) error
}
type TestResultStore interface {
	Ping() error
	CreateTestResult(result *TestResult) error
	UpdateTestResult(result *TestResult) error
	GetTestResultByID(id int) (*TestResult, error)
	DestroyTestResult(result *TestResult) error
}
type RunResultStore interface {
	Ping() error
	CreateRunResult(result *RunResult) error
	UpdateRunResult(result *RunResult) error
	GetRunResultByID(id int) (*RunResult, error)
	DestroyRunResult(result *RunResult) error
}
type AssignmentStore interface {
	Ping() error
	CreateAssignment(assig *Assignment) error
	UpdateAssignment(assig *Assignment) error
	GetAssignmentByID(id int) (*Assignment, error)
	DestroyAssignment(assig *Assignment) error
}
