package main

import (
	"fmt"
	"piscine"
	"os"
)

func f(a, b int) int {
	return (a - b)
}

func main() {
	var args[3] string

	if piscine.IsValidArgs(args) == false {
		os.Exit(1)
	}

	i1 := piscine.Atoi(args[0])
	i2 := piscine.Atoi(args[2])

	if piscine.IsOverFlow(i1, args[1], i2) {
		os.Exit(1)
	}

	if args[1] == "+" {
		fmt.Println(i1 + i2)
	} else if args[1] == "-" {
		fmt.Println(i1 - i2)
	} else if args[1] == "/" {
		fmt.Println(i1 / i2)
	} else if args[1] == "*" {
		fmt.Println(i1 * i2)
	} else if args[1] == "%" {
		fmt.Println(i1 % i2)
	}
}
