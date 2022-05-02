package piscine

import "ft"

func is_BaseValid(base []rune) bool {
	var count int
	for i, r := range base {
		if r == '+' || r == '-' {
			return (false)
		}
		for j, r2 := range base {
			if i != j && r == r2 {
				return (false)
			}
		}
		count++
	}
	return (count > 1)
}

func printStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
}

func strLen(s []rune) uint {
	var ret uint

	for range s {
		ret++
	}
	return ret
}

func PrintNbrBase(nbr int, base string) {
	var out string
	var ui uint
	rbase := []rune(base)

	if is_BaseValid(rbase) == false {
		printStr("NV")
		return
	}
	if nbr < 0 {
		ft.PrintRune('-')
		ui = uint(nbr * -1)
	} else {
		ui = uint(nbr)
	}
	baseLen := strLen(rbase)
	for ; ui > 0; ui /= baseLen {
		out = string(rbase[ui % baseLen]) + out
	}
	printStr(out)
}
