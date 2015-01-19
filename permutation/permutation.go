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
