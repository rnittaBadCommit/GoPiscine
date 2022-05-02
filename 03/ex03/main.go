package main

import (
	"fmt"
	"piscine"
	"ft"
)

func main() {
	fmt.Println(piscine.Index("Heあいうllo!", "Heあいうl"))
	fmt.Println(piscine.Index("Heあいうllo!", "あいうl"))

	ft.PrintRune('\n')


	fmt.Println(piscine.Index("Hello!", "l"))
	fmt.Println(piscine.Index("Salut!", "alu"))
	fmt.Println(piscine.Index("Salut!", "!"))
	fmt.Println(piscine.Index("Salut!", "t!"))
	fmt.Println(piscine.Index("Salut!", "Salut!"))
	fmt.Println(piscine.Index("Salut!", "Salutz"))
	fmt.Println(piscine.Index("Salut!", ""))
	fmt.Println(piscine.Index("Ola!", "hOl"))
	fmt.Println(piscine.Index("", "hOl"))
}
