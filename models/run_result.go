package models

import (
	"database/sql"
	"strconv"
	"time"
)

type RunResult struct {
	ID                 int
	TestResult         int
	CompilationFailure string

	CreatedAt   time.Time
	LastUpdated time.Time
}

func (run *RunResult) GetID() string {
	return strconv.Itoa(run.ID)
}
