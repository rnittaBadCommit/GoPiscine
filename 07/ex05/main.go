package main

import (
	"fmt"
	"piscine"
)

func main() {
	var args[3] string

	if piscine.IsValidArgs(args) == false {
		return
	}

	i1, err := piscine.Atoi(args[0])
	i2, err2 := piscine.Atoi(args[2])
	if err + err2 != 0 {
		return
	}

	if piscine.IsOverFlow(i1, args[1], i2) {
		return
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
