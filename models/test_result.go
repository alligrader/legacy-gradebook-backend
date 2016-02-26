package models

import (
	"database/sql"
	"strconv"
	"time"
)

type TestResult struct {
	ID           int
	Passed       bool
	ErrorMessage string
	TestID       int

	CreatedAt   time.Time
	LastUpdated time.Time
}

func (test *TestResult) GetID() string {
	return strconv.Itoa(test.ID)
}
