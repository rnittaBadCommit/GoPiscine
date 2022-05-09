package main

import (
	"fmt"
	"piscine"
)

func main() {
	for i := -1; i < 100; i++ {
		if piscine.IsPrime(i) {
			fmt.Printf("%d\n", i);
		}
	}
}
