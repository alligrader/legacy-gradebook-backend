package spec

import (
	"testing"
)

func TestCreateInstructor(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Select")
	}
}

func TestUpdateInstructor(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Create")
	}
}

func TestDeleteInstructor(t *testing.T) {
}

func TestSelectInstructor(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Create")
	}
}
