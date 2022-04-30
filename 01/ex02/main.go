package main

import (
	"fmt"
	"piscine"
)

func main() {
	a := 13
	b := 2
	var div, mod int
	piscine.DivMod(a, b, &div, &mod)
	fmt.Println(div)
	fmt.Println(mod)
	a = 13
	b = 0
	piscine.DivMod(a, b, &div, &mod)
	fmt.Println(div)
	fmt.Println(mod)
}
