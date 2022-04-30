package main

import (
	"fmt"
	"piscine"
)

func main() {
	nb := 0
	piscine.PointOne(&nb)
	fmt.Println(nb)
	piscine.PointOne(nil)
}
