package models

import (
	"strconv"
	"time"
)

type Teacher struct {
	ID int

	CourseID    int
	CreatedAt   time.Time
	LastUpdated time.Time
}

func (teacher *Teacher) GetID() string {
	return strconv.Itoa(teacher.ID)
}
