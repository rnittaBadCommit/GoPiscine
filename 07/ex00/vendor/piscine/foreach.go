package piscine

import "ft"

func ForEach(f func(int), a []int) {
	for _, i := range a {
		f(i)
	}
	ft.PrintRune('\n')
}
