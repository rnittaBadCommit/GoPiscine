package main

import (
	"fmt"
	"piscine"
)

func f(a, b int) int {
	return (a - b)
}

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}
	a3 := []int{0, 2, 2, 3}

	result1 := piscine.IsSorted(f, a1)
	result2 := piscine.IsSorted(f, a2)
	result3 := piscine.IsSorted(f, a3)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}
