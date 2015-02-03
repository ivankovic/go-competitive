package prime

import (
	"github.com/ivankovic/go-competitive/pow"
	"math"
)

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

func IsPrime(n int) bool {
	switch {
	case n < 2:
		return false
	case n == 2:
		return true
	case n%2 == 0:
		return false
	case n < 2047:
		return primalityTest(n, []int{2})
	case n < 1373653:
		return primalityTest(n, []int{2, 3})
	case n < 9080191:
		return primalityTest(n, []int{31, 73})
	case n < 25326001:
		return primalityTest(n, []int{2, 3, 5})
	case n < 4759123141:
		return primalityTest(n, []int{2, 7, 61})
	}
	// TODO: Panic instead of return.
	return false
}

func primalityTest(n int, witnesses []int) bool {
	d, s := n-1, uint(0)
	for i := uint(31); i > 0; i-- {
		if (d % (1 << i)) == 0 {
			d /= 1 << i
			s += i
		}
	}

WitnessLoop:
	for _, a := range witnesses {
		x := pow.Mod(a, d, n)
		if x == 1 || x == n-1 {
			continue WitnessLoop
		}
		for i := uint(0); i < s-1; i++ {
			x = pow.Mod(x, 2, n)
			if x == 1 {
				return false
			} else if x == n-1 {
				continue WitnessLoop
			}
		}
		return false
	}

	return true
}
