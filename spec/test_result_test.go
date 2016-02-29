package spec

import (
	"testing"
)

func TestCreateTestResult(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Select")
	}
}

func TestUpdateTestResult(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Create")
	}
}

func TestDeleteTestResult(t *testing.T) {
}

func TestSelectTestResult(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing dependent on Create")
	}
}
