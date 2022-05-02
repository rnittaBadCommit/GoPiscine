package main

import (
	"fmt"
	"piscine"
)

func main() {
	for i := -1; i < 20; i++ {
		fmt.Println(i, piscine.Sqrt(i))
	}
	fmt.Println(piscine.Sqrt(9223372036854775807))
}
