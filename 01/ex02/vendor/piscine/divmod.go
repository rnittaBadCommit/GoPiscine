package piscine

func DivMod(a, b int, div, mod *int) {
	if b != 0 {
		if div != nil {
			*div = a / b
		}
		if mod != nil {
			*mod = a % b
		}
	}
}
