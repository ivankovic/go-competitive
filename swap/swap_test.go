package swap

import (
	"reflect"
	"testing"
)

func TestSwapInt(t *testing.T) {

	for ti, tt := range []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2}, []int{2, 1}},
		{[]int{-1, 2}, []int{2, -1}},
	} {
		Swap(&tt.input[0], &tt.input[1])
		if !reflect.DeepEqual(tt.input, tt.expected) {
			t.Errorf("(%d) Swap => %v expected %v", ti, tt.input, tt.expected)
		}
	}
}
