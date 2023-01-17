package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		// array
		numbers := []int{1, 2, 3, 4, 5}

		result := Sum(numbers)
		expected := 15

		if expected != result {
			t.Errorf("Expected '%d' but received '%d' when numbers = '%v'", expected, result, numbers)
		}
	})

	t.Run("collection of 3 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		result := Sum(numbers)
		expected := 6

		if expected != result {
			t.Errorf("Expected '%d' but received '%d' when numbers = '%v'", expected, result, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	firstRow := []int{1, 2}
	secondRow := []int{3, 4}

	sum := SumAll(firstRow, secondRow)
	expected := []int{3, 7}

	if !reflect.DeepEqual(expected, sum) { // basically "expected != sum" but for arrays/slices
		t.Errorf("Expected '%d' but received '%d' when firstRow = %v, secondRow = %v", expected, sum, firstRow, secondRow)
	}
}

func TestSumAllTails(t *testing.T) {
	t.Run("make a sum of slices", func(t *testing.T) {
		firstRow := []int{1, 2}
		secondRow := []int{3, 4}

		sum := SumAllTails(firstRow, secondRow)
		expected := []int{2, 4}

		if !reflect.DeepEqual(expected, sum) {
			t.Errorf("Expected '%d' but received '%d' when firstRow = %v, secondRow = %v", expected, sum, firstRow, secondRow)
		}
	})

	t.Run("make a sum of slices with empty items", func(t *testing.T) {
		firstRow := []int{1, 2, 3}
		secondRow := []int{3, 4, 5}

		sum := SumAllTails(firstRow, secondRow)
		expected := []int{5, 9}

		if !reflect.DeepEqual(expected, sum) {
			t.Errorf("Expected '%d' but received '%d' when firstRow = %v, secondRow = %v", expected, sum, firstRow, secondRow)
		}
	})
}
