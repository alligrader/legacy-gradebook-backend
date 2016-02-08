package spec

import (
	"testing"
)

func TestCreateAssignment(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Select")
	}
}

func TestUpdateAssignment(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Create")
	}
}

func TestDeleteAssignment(t *testing.T) {
}

func TestSelectAssignment(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Create")
	}
}
