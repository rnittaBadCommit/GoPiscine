package main

import (
	"fmt"
	"piscine"
)

func main() {
	for i := -1; i < 100; i++ {
		fmt.Printf("%d: %t\n", i, piscine.FindNextPrime(i));
	}
}
