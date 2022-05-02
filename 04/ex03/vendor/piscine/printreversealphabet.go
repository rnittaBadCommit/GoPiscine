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

func returnArgs() []string {
	return (os.Args)
}

func SortParams() {
//	args := [][]rune(os.Args)
	args := returnArgs()

	for i, _ := range args {
		if i == 0 {
			continue
		}
		for j, _ := range args {
			if j == 0 {
				continue
			}
			if i < j && args[i] > args[j] {
				args[i], args[j] = args[j], args[i]
			}
		}
	}
	for i, s := range args {
		if i == 0 {
			continue
		}
		printStr(s)
		ft.PrintRune('\n')
	}
}
