package piscine

import "ft"

func PrintNbr(nbr int) {
	var ui uint

	if nbr < 0 {
		ft.PrintRune('-')
		ui = uint(-1 * nbr)
	} else {
		ui = uint(nbr)
	}
	var s string
	for ; ui > 0; ui /= 10 {
		s = string('0' + ui % 10) + s
	}
	for _, r := range s {
		ft.PrintRune(r)
	}
}
