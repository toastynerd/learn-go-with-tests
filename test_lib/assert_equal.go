package test_lib

import (
	"testing"
)

func AssertEqualInt(expected int, got int, t *testing.T) {
	if expected != got {
		t.Errorf("expected %d but got %d", expected, got);
	}
}
