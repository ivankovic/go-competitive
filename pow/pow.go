package pow

// TODO: Unwrap recursion
func Mod(x, p, n int) int {
	switch {
	case p == 0:
		return 1
	case p == 1:
		return x % n
	case p%2 == 0:
		return Mod((x*x)%n, p/2, n) % n
	}
	return (x * Mod((x*x)%n, (p-1)/2, n)) % n
}
