package swap

func Swap(a, b *int) {
	t := *a
	(*a) = (*b)
	(*b) = t
}
