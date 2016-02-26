package spec

import (
	"testing"
)

func TestCreateClass(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Select")
	}
}

func TestUpdateClass(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Create")
	}
}

func TestDeleteClass(t *testing.T) {
}

func TestSelectClass(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Create")
	}
}
