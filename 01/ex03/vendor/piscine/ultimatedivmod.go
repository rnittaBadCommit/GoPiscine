package piscine

func UltimateDivMod(a, b *int) {
	if a != nil && b != nil && *b != 0 {
		*a, *b = *a / *b, *a%*b
	}
}
