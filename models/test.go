package models

import (
	"strconv"
	"time"
)

type Test struct {
	ID     int
	Weight int

	CreatedAt   time.Time
	LastUpdated time.Time
}

func (test *Test) GetID() string {
	return strconv.Itoa(test.ID)
}
