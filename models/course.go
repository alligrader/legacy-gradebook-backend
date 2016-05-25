package models

import (
	"strconv"
	"time"
)

type Course struct {
	ID          int64  `db:"id"`
	Name        string `db:"name"`
	Students    []*Student
	Teachers    []*Teacher
	CreatedAt   time.Time `db:"created_at"`
	LastUpdated time.Time `db:"last_updated"`
}

func NewCourse() *Course {
	return &Course{
		Students: []*Student{},
		Teachers: []*Teacher{},
	}
}

func (course *Course) GetID() string {
	return strconv.FormatInt(course.ID, 10)
}

func (course *Course) Equals(other *Course) bool {
	if other == nil {
		return false
	}
	for _, teacher := range course.Teachers {
		if !other.containsTeacher(teacher) {
			return false
		}
	}

	for _, teacher := range other.Teachers {
		if !course.containsTeacher(teacher) {
			return false
		}
	}

	for _, student := range course.Students {
		if !other.containsStudent(student) {
			return false
		}
	}

	for _, student := range other.Students {
		if !course.containsStudent(student) {
			return false
		}
	}

	return course.Name == other.Name
}

func (course *Course) containsStudent(student *Student) bool {
	result := false
	for _, otherStudent := range course.Students {
		if student.Equals(otherStudent) {
			result = true
			break
		}
	}
	return result
}

func (course *Course) containsTeacher(teacher *Teacher) bool {
	result := false
	for _, otherTeacher := range course.Teachers {
		if teacher.Equals(otherTeacher) {
			result = true
			break
		}
	}
	return result
}
