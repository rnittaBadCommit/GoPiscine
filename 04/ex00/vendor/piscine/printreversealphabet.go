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

func PrintProgramName() {
	printStr(os.Args[0])
	ft.PrintRune('\n')
}
