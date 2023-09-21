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

		test_lib.AssertEqualInt(expected, got, t)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		expected := 6

		test_lib.AssertEqualInt(expected, got, t)
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{3, 4})
	expected := []int{3, 7}

	test_lib.AssertEqualIntSlice(expected, got, t)
}

func TestSumAllTails(t *testing.T) {
	t.Run("sum the tails of all slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{3, 4}, []int{5, 6, 2})
		expected := []int{2, 4, 8}

		test_lib.AssertEqualIntSlice(expected, got, t)

	})

	t.Run("slices with only head", func(t* testing.T) {
		got := SumAllTails([]int{1}, []int{2}, []int{3})
		expected := []int{0, 0, 0}

		test_lib.AssertEqualIntSlice(expected, got, t)
	})

	t.Run("empty slices", func(t* testing.T) {
		got := SumAllTails([]int{}, []int{}, []int{3})
		expected := []int{0, 0, 0}

		test_lib.AssertEqualIntSlice(expected, got, t)
	})
}

func TestSumAllHeads(t *testing.T) {
	t.Run("sum the heads of all slices into a single value", func(t *testing.T) {
		got := SumAllHeads([]int{1, 2, 3}, []int{2, 3, 4}, []int{3, 4, 5})
		expected := 6

		test_lib.AssertEqualInt(expected, got, t)
	})

	t.Run("empty slices", func(t *testing.T) {
		got := SumAllHeads([]int{1}, []int{2, 3, 4}, []int{})
		expected := 3

		test_lib.AssertEqualInt(expected, got, t)
	})
}
