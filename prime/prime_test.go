package prime

import "testing"

func TestSieveUnder20(t *testing.T) {
	primes, sieve := Sieve(21)

	for i, e := range []bool{
		true,                            // 0
		true, false, false, true, false, // 1-5
		true, false, true, true, true, // 6-10
		false, true, false, true, true, // 11-15
		true, false, true, false, true, // 15-20
	} {
		if sieve[i] != e {
			t.Errorf("sieve[%d] = %t expected %t", i, sieve[i], e)
		}
	}

	for i, e := range []int{
		2, 3, 5, 7, 11, 13, 17, 19,
	} {
		if primes[i] != e {
			t.Errorf("primes[%d] = %d expected %d", i, primes[i], e)
		}
	}
}

func TestSieveUnder1000000(t *testing.T) {
	for ti, tt := range []struct {
		n    int
		last int
		size int
	}{
		{21, 19, 8},
		{1000, 997, 168},
		{10000, 9973, 1229},
		{1000000, 999983, 78498},
		{999983, 999979, 78497},
	} {
		primes, _ := Sieve(tt.n)
		if len(primes) != tt.size || primes[len(primes)-1] != tt.last {
			t.Errorf("(%d) Sieve(%d) size = %d last = %d expected %d %d", ti, tt.n,
				len(primes), primes[len(primes)-1], tt.size, tt.last)
		}
	}
}

func TestIsPrimeGoldens(t *testing.T) {
	for ti, tt := range []struct {
		n     int
		prime bool
	}{
		{1, false},
		{2, true},
		{2113, true},
		{341, false},
		{180811, true},
		{39916801, true},
		{16, false},
		{39916800, false},
	} {
		result := IsPrime(tt.n)
		if result != tt.prime {
			t.Errorf("(%d) IsPrime(%d) = %t expected %t", ti, tt.n, result, tt.prime)
		}
	}
}

func TestCrossIsPrimeAndSieve(t *testing.T) {
	_, sieve := Sieve(10000)

	for i, v := range sieve {
		result := !IsPrime(i)
		if result != v {
			t.Errorf("Sieve[%d] = %t and NOT(IsPrime(%d)) = %t", i, v, i, result)
		}
	}
}
