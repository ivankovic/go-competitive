package pow

import "testing"

func TestPowGolden(t *testing.T) {
	for ti, tt := range []struct {
		x        int
		p        int
		n        int
		expected int
	}{
		{1, 1, 1, 0},
		{2, 0, 10, 1},
		{2, 1, 10, 2},
		{2, 2, 10, 4},
		{2, 4, 100, 16},
		{2, 4, 10, 6},
		{14, 5, 9845, 6194},
	} {
		result := Mod(tt.x, tt.p, tt.n)
		if result != tt.expected {
			t.Errorf("(%d) pow.Mod(%d, %d, %d) = %d expected %d", ti, tt.x, tt.p, tt.n, result, tt.expected)
		}
	}
}
