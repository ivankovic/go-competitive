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
