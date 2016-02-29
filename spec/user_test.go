package spec

import (
	"testing"
)

func TestCreateUser(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Select")
	}
}

func TestUpdateUser(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Create")
	}
}

func TestDeleteUser(t *testing.T) {
}

func TestSelectUser(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Create")
	}
}
