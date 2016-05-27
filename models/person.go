package models

import (
	"strconv"
	"time"
)

type Person struct {
	ID        int64
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Username  string `db:"username"`
	Password  []byte `db:"password"`

	CreatedAt   time.Time `db:"created_at"`
	LastUpdated time.Time `db:"last_updated"`
}

func (person *Person) InsertColumns() string {
	return "first_name, last_name, username, password"
}

func (person *Person) GetColumns() string {
	return "first_name, last_name, username, created_at, last_updated"
}

func (person *Person) Fields() []interface{} {
	return []interface{}{
		&person.ID, &person.FirstName, &person.LastName,
		&person.Username, &person.CreatedAt, &person.LastUpdated,
	}
}

func (person *Person) GetID() string {
	return strconv.FormatInt(person.ID, 10)
}

func (person *Person) Equals(other *Person) bool {
	if other == nil {
		return false
	}
	return person.FirstName == other.FirstName &&
		person.LastName == other.LastName &&
		person.Username == other.Username
}
