/*
   Copyright 2015 Marko Ivankovic

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
package gcd

import "testing"

func TestGCD(t *testing.T) {
	for ti, tt := range []struct {
		a        int
		b        int
		expected int
	}{
		{10, 5, 5},
		{10, 2, 2},
		{54, 24, 6},
		{-54, 24, 6},
		{54, -24, 6},
		{-54, -24, 6},
		{180, 48, 12},
		// Special cases
		{10, 0, 1},
		{0, 10, 1},
		{0, 0, 1},
	} {
		result := GCD(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("(%d) GCD(%d, %d) = %d - expected %d", ti, tt.a, tt.b, result, tt.expected)
		}
	}
}
