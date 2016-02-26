package models

import (
	"strconv"
	"time"
)

type Repo struct {
	ID           int
	StudentID    int
	AssignmentID int
	RunResultID  int

	CreatedAt   time.Time
	LastUpdated time.Time
}

func (repo *Repo) GetID() string {
	return strconv.Itoa(repo.ID)
}
