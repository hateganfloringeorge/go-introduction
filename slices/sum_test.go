package slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	checkSums := func (t *testing.T, got, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v expected %v", got, expected)
		}
	}

	t.Run("sum of a slice", func(t *testing.T) {
		numbers := []int {4, 3, 2, 1}

		result := Sum(numbers)
		expected := 10

		if result != expected {
			t.Errorf("expected %d but got %d, %v", expected, result, numbers)
		}
	})

	t.Run("sumAll for a collection of slices", func(t *testing.T) {
		result := SumAll([]int{1,2}, []int{0,9})
		expected := []int{3, 9}

		checkSums(t, result, expected)
	})

	t.Run("sum all tails", func(t *testing.T) {
		result := SumAllTails([]int{1,2}, []int{0,9})
		expected := []int {2, 9}

		checkSums(t, result, expected)
	})

	t.Run("sum all tails with empty slice", func(t *testing.T) {
		result := SumAllTails([]int{}, []int{0,9}, []int{1, 2, 3}, []int {4})
		expected := []int {0, 9, 5, 0}

		checkSums(t, result, expected)
	})
}