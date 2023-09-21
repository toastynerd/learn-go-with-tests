package test_lib

import (
	"testing"
)

func AssertEqualInt(expected int, got int, t *testing.T) {
	if expected != got {
		t.Errorf("expected %d but got %d", expected, got);
	}
}

func IntSlicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, number := range a {
		if number != b[i] {
			return false
		}
	}
	return true
}

func AssertEqualIntSlice(expected []int, got []int, t *testing.T) {
	if !IntSlicesEqual(expected, got) {
		t.Errorf("expected %v but got %v", expected, got)
	}
}
