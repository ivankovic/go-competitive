package prime

import "math"

func Sieve(n int) ([]int, []bool) {
	sieve := make([]bool, n)
	primes := make([]int, 0, int(float64(n)/math.Log(float64(n))))

	// 0 and 1 are not prime
	sieve[0] = true
	sieve[1] = true

	for i := 2; i < n; i++ {
		if !sieve[i] {
			primes = append(primes, i)
			for j := i + i; j < n; j += i {
				sieve[j] = true
			}
		}
	}

	return primes, sieve
}
