package piscine

import "ft"

func printStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
}

func makeStr(nb1, nb2 int) string {
	var ret string

	ret = string('0' + rune(nb2 % 10))
	ret = string('0' + rune(nb2 / 10)) + ret
	ret = string('0' + rune(nb1 % 10)) + string(' ') + ret
	ret = string('0' + rune(nb1 / 10)) + ret
	return (ret)
}

func PrintComb2() {
	var flag bool

	flag = false
	for i := 0; i <= 99; i++ {
		for j := i + 1; j <= 99; j++ {
			if flag {
				printStr(", ")
			}
			printStr(makeStr(i, j))
			flag = true
		}
	}
}

