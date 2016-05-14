package models

import (
	"strconv"
	"time"
)

type Teacher struct {
	ID int
	Person

	CreatedAt   time.Time
	LastUpdated time.Time
}

func (teacher *Teacher) GetID() string {
	return strconv.Itoa(teacher.ID)
}

func (teacher *Teacher) Equals(other *Teacher) bool {
	if other == nil {
		return false
	}
	return teacher.Person.Equals(&other.Person)
}
