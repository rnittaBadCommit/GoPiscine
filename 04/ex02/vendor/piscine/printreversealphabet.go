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

func RevParams() {
	var output string

	for i, s := range os.Args {
		if i > 0 {
			output = s + "\n" +  output
		}
	}
	printStr(output)
}
