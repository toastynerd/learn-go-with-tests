package array_and_slice

import (
	"testing"
	"learn-go-with-tests/test_lib"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		expected := 15

		test_lib.AssertEqualInt(got, expected, t)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		expected := 6

		test_lib.AssertEqualInt(got, expected, t)
	})
}
