package models

import (
	"strconv"
	"time"
)

type Course struct {
	ID          int
	CreatedAt   time.Time
	LastUpdated time.Time
}

func (course *Course) GetID() string {
	return strconv.Itoa(course.ID)
}
