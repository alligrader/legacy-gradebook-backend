package spec

import (
	"testing"
)

func TestCreateStudent(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Select")
	}
}

func TestUpdateStudent(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Create")
	}
}

func TestDeleteStudent(t *testing.T) {
}

func TestSelectStudent(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Create")
	}
}
