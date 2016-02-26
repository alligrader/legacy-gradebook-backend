package models

import (
	"database/sql"
	"strconv"
	"time"
)

type User struct {
	ID int

	GithubID    int
	Email       string
	StudentID   sql.NullInt
	TeacherID   sql.NullInt
	CreatedAt   time.Time
	LastUpdated time.Time
}

func (user *User) GetID() string {
	return strconv.Itoa(user.ID)
}
