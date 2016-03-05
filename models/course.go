package models

import (
	"strconv"
	"time"
)

type Course struct {
	ID          int       `db:"id"`
	Name        string    `db:"name"`
	CreatedAt   time.Time `db:"created_at"`
	LastUpdated time.Time `db:"last_updated"`
}

func (course *Course) GetID() string {
	return strconv.Itoa(course.ID)
}

func (course *Course) Equals(other *Course) bool {
	if other == nil {
		return false
	}
	return course.ID == other.ID && course.Name == other.Name
}
