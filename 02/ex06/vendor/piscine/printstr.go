package piscine

func Swap(a, b *int) {
	if a != nil && b != nil {
		*a, *b = *b, *a
	}
}
