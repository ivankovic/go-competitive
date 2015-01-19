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
package permutation

import "github.com/ivankovic/go-competitive/swap"

func NextPermutation(slice *[]int) bool {
	k, l := -1, -1
	for k = len(*slice) - 2; k >= 0; k-- {
		if (*slice)[k] < (*slice)[k+1] {
			break
		}
	}
	if k == -1 {
		for i := 0; i < len(*slice)/2; i++ {
			swap.Swap(&(*slice)[i], &(*slice)[len(*slice)-i-1])
		}
		return false
	}

	for l = len(*slice) - 1; l > k; l-- {
		if (*slice)[l] > (*slice)[k] {
			break
		}
	}

	swap.Swap(&(*slice)[k], &(*slice)[l])

	for i := 1; i < (len(*slice)-k+1)/2; i++ {
		swap.Swap(&(*slice)[k+i], &(*slice)[len(*slice)-i])
	}

	return true
}
