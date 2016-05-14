package models

import (
	"strconv"
	"time"
)

type Student struct {
	ID int
	Person

	CreatedAt   time.Time
	LastUpdated time.Time
}

func (student *Student) GetID() string {
	return strconv.Itoa(student.ID)
}

func (student *Student) Equals(other *Student) bool {
	if other == nil {
		return false
	}
	return student.Person.Equals(&other.Person)
}
