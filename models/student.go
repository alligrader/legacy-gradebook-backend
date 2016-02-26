package models

import (
	"strconv"
	"time"
)

type Student struct {
	ID int

	CourseID    int
	CreatedAt   time.Time
	LastUpdated time.Time
}

func (student *Student) GetID() string {
	return strconv.Itoa(student.ID)
}
