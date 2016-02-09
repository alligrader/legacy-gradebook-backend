package spec

import (
	"testing"
)

func TestCreateSchool(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Select")
	}
}

func TestUpdateSchool(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Create")
	}
}

func TestDeleteSchool(t *testing.T) {
}

func TestSelectSchool(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Create")
	}
}
