package piscine

import "ft"

func printStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
}

func makeStr(nb1, nb2, nb3 int) string {
	var ret string

	ret = string('0' + rune(nb1))
	ret = ret + string('0' + rune(nb2))
	ret = ret + string('0' + rune(nb3))
	return ret
}


func PrintComb() {
	var flag bool

	flag = false
	for i := 0; i <= 9; i++ {
		for j := i + 1; j <= 9; j++ {
			for k := j + 1; k <= 9; k++ {
				if flag {
					printStr(", ")
				}
				printStr(makeStr(i, j, k))
				flag = true
			}
		}
	}
}

