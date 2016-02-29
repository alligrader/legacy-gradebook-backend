package spec

import (
	"testing"
)

func TestCreateRunResult(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Select")
	}
}

func TestUpdateRunResult(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Create")
	}
}

func TestDeleteRunResult(t *testing.T) {
}

func TestSelectRunResult(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Create")
	}
}
