package models

import (
	"fmt"
	"strconv"
	"time"
)

type Student struct {
	ID int64
	User

	CreatedAt   time.Time
	LastUpdated time.Time
}

func (student *Student) GetID() string {
	return strconv.FormatInt(student.ID, 10)
}

func (student *Student) Equals(other *Student) bool {
	if other == nil {
		return false
	}
	return student.User.Equals(&other.User)
}

func (student *Student) String() string {
	return fmt.Sprintf(" { ID: %v, FirstName: %v, LastName: %v, Username: %v }",
		student.ID, student.User.FirstName, student.User.LastName, student.User.Username,
	)
}
