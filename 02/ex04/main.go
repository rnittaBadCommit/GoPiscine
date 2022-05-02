package main

import (
	"fmt"
	"piscine"
)

func main() {
	for i := -1; i < 10; i++ {
		fmt.Println(piscine.Fibonacci(i))
	}
}
