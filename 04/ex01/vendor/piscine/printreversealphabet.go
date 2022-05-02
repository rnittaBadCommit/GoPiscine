package piscine

import (
	"ft"
	"os"
)

func printStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
}

func PrintParams() {
	for i, s := range os.Args {
		if i > 0 {
			printStr(s)
			ft.PrintRune('\n')
		}
	}
}
