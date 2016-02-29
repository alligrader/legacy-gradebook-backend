package spec

import (
	"testing"
)

func TestCreateTestModel(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Select")
	}
}

func TestUpdateTestModel(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Create")
	}
}

func TestDeleteTestModel(t *testing.T) {
}

func TestSelectTestModel(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Create")
	}
}
