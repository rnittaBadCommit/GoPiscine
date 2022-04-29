package piscine

import "ft"

func PrintAlphabet() {
	for _, r := range []rune("abcdefghijklmnopqrstuvwxyz\n") {
		ft.PrintRune(r)
	}
}
