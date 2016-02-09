package spec

import (
	"testing"
)

func TestCreateTeacher(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Select")
	}
}

func TestUpdateTeacher(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Create")
	}
}

func TestDeleteTeacher(t *testing.T) {
}

func TestSelectTeacher(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Create")
	}
}
