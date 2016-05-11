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

type AssignmentStore interface {
	Ping() error
	CreateAssignment(assig *Assignment) error
	UpdateAssignment(assig *Assignment) error
	GetAssignmentByID(id int) (*Assignment, error)
	DestroyAssignment(assig *Assignment) error
}
