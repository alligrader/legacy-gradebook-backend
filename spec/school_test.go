package spec

import (
	"testing"
)

func TestCreateInstitution(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Select")
	}
}

func TestUpdateInstitution(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Create")
	}
}

func TestDeleteInstitution(t *testing.T) {
}

func TestSelectInstitution(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Create")
	}
}
