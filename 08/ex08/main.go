package main

import (
	"fmt"
	"piscine"
)

func main() {
	a := []int{23, 123, 1, 11, 55, 93}
	max := piscine.Max(a)
	fmt.Println(max)

	a = []int{}
	max = piscine.Max(a)
	fmt.Println(max)

	a = []int{-100}
	max = piscine.Max(a)
	fmt.Println(max)
}
