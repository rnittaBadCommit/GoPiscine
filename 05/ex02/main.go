package main

import (
	"piscine"
	"fmt"
)

func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(piscine.ConcatParams(test))

	test = []string{}
	fmt.Println(piscine.ConcatParams(test))
}
