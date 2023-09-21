package test_lib

import (
	"testing"
)

func TestIntSlicesEqual(t *testing.T) {
	t.Run("two equal slices", func(t *testing.T) {

		a := []int{1, 2}
		b := []int{1, 2}

		if !IntSlicesEqual(a, b) {
			t.Errorf("Slices %v and %v did not evaluate as equal", a, b)
		}
	})

	t.Run("Two different length slices", func(t *testing.T) {

		a := []int{1, 2}
		b := []int{1, 2, 3}

		if IntSlicesEqual(a, b) {
			t.Errorf("Slices %v and %v evaluated as equal", a, b)
		}
	})

	t.Run("Two different equal length slice", func(t *testing.T) {

		a := []int{1, 2}
		b := []int{3, 4}

		if IntSlicesEqual(a, b) {
			t.Errorf("Slices %v and %v evaluated as equal", a, b)
		}
	})

	t.Run("Two different equal lenght slice with one equal element", func(t *testing.T) {

		a := []int{1, 2}
		b := []int{1, 3}

		if IntSlicesEqual(a, b) {
			t.Errorf("Slices %v and %v evaluated as equal", a, b)
		}
	})
}
