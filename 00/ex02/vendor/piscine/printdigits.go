package piscine

import "ft"

func PrintDigits() {
	for i := 0; i < 10; i++ {
		ft.PrintRune(rune(i) + '0')
	}
	ft.PrintRune('\n')
}
