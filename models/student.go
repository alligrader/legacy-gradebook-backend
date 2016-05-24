package models

import (
	"fmt"
	"strconv"
	"time"
)

type Student struct {
	ID int64
	Person

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
	return student.Person.Equals(&other.Person)
}

func (student *Student) String() string {
	return fmt.Sprintf(" { ID: %v, FirstName: %v, LastName: %v, Username: %v }",
		student.ID, student.Person.FirstName, student.Person.LastName, student.Person.Username,
	)
}
