package main

import (
	"ft"
	"piscine"
)

type boolean int
const yes = 1
const no = 0

func even(nbr int) int {
	if nbr % 2 == 0 {
		return (1)
	} else {
		return (0)
	}
}

func printStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
}

func isEven(nbr int) boolean {
	if even(nbr) == 1 {
		return yes
	} else {
		return no
	}
}

func main() {
	if isEven(piscine.LengthOfArg()) == 1 {
		printStr("I have an even number of arguments\n")
	} else {
		printStr("I have an odd number of arguments\n")
	}
}
