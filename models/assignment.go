package models

import (
	"strconv"
	"time"
)

type Assignment struct {
	ID        int
	StudentID int
	TeacherID int

	CreatedAt   time.Time
	LastUpdated time.Time
}

func (assig *Assignment) GetID() string {
	return strconv.Itoa(assig.ID)
}

// TODO impl
func (assig *Assignment) Equals(other *Assignment) bool {
	return false
}
