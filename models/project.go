package models

import (
	"strconv"
	"time"
)

type Assignment struct {
	ID        int64
	StudentID int
	TeacherID int

	CreatedAt   time.Time
	LastUpdated time.Time
}

func (assig *Assignment) GetID() string {
	return strconv.FormatInt(assig.ID, 10)
}

// TODO impl
func (assig *Assignment) Equals(other *Assignment) bool {
	return false
}
