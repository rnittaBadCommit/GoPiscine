package piscine

import "ft"

func printstring(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
}

func recursive(left, n int, is_first bool, s string) {
	for ; n <= 9; n++ {
		tmp := s + string('0'+rune(n))
		if left == 0 {
			if is_first {
				printstring(", ")
			}
			is_first = true
			printstring(tmp)
		} else {
			recursive(left-1, n+1, is_first, tmp)
			is_first = true
		}
	}
}

func PrintCombN(n int) {
	recursive(n-1, 0, false, "")
	ft.PrintRune('\n')
}
