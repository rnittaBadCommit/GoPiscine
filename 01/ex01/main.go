package main

import (
	"fmt"
	"piscine"
)

func main() {
	a := 0
	b := &a
	n := &b
	piscine.UltimatePointOne(&n)
	fmt.Println(a)
	piscine.UltimatePointOne(nil)
	b = nil
	piscine.UltimatePointOne(&n)
}
