package models

import (
	"strconv"
	"time"
)

type Person struct {
	ID        int
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Username  string `db:"username"`
	Password  []byte `db:"password"`

	CreatedAt   time.Time `db:"created_at"`
	LastUpdated time.Time `db:"last_updated"`
}

func (person *Person) GetID() string {
	return strconv.Itoa(person.ID)
}

func (person *Person) Equals(other *Person) bool {
	return person.FirstName == other.FirstName &&
		person.LastName == other.LastName &&
		person.Username == other.Username
}
