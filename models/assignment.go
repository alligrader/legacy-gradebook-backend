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
