package permutation

import (
	"reflect"
	"testing"
)

func TestIntSliceHandmade(t *testing.T) {
	value := []int{1, 2, 3}

	for ti, tt := range []struct {
		expected []int
		result   bool
	}{
		{[]int{1, 3, 2}, true},
		{[]int{2, 1, 3}, true},
		{[]int{2, 3, 1}, true},
		{[]int{3, 1, 2}, true},
		{[]int{3, 2, 1}, true},
		{[]int{1, 2, 3}, false},
	} {
		result := NextPermutation(&value)
		if !reflect.DeepEqual(value, tt.expected) || result != tt.result {
			t.Errorf("(%d) NextPermutation => %v = %t expected %v = %t", ti, value, result, tt.expected, tt.result)
		}
	}
}

func TestIntSliceNumbers(t *testing.T) {
	for ti, tt := range []struct {
		input    []int
		expected int
	}{
		{[]int{1}, 0},
		{[]int{1, 2}, 1},
		{[]int{1, 2, 3}, 5},
		{[]int{1, 2, 3, 4}, 23},
		{[]int{1, 2, 3, 4, 5}, 119},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3628799},
	} {
		count := 0
		for NextPermutation(&tt.input) {
			count++
		}
		if count != tt.expected {
			t.Errorf("(%d) NextPermutation count => %d expected %d", ti, count, tt.expected)
		}
	}
}

func BenchmarkIntSlice(b *testing.B) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for bi := 0; bi < b.N; bi++ {
		for NextPermutation(&input) {
		}
	}
}
