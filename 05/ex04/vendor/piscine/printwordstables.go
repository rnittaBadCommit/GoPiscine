package piscine

import "ft"

func SplitWhiteSpaces(s string) []string {
	var ret []string
	var start, end int

	for _, r := range s {
		if r == ' ' || r == '\t' ||  r == '\n' {
			if start != end {
				ret = append(ret, s[start:end])
			}
			start = end + 1
		}
		end++
	}
	if start < end {
		ret = append(ret, s[start:])
	}
	return (ret)
}

func printString(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
}

func PrintWordsTables(a []string) {
	var new_s string

	for _, s := range a {
			new_s += s + string('\n')
	}
	printString(new_s)
}
