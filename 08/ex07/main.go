package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.ActiveBits(7))
	fmt.Println(piscine.ActiveBits(0))
	fmt.Println(piscine.ActiveBits(128))
	fmt.Println(piscine.ActiveBits(-1))
}
