package main

import (
	"piscine"
	"fmt"
)

func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(piscine.ConcatParams(test))

	fmt.Printf("=============\n")

	test = []string{}
	fmt.Println(piscine.ConcatParams(test))
}
