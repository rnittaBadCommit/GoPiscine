package main

import (
	"piscine"
	"fmt"
)

func main() {
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("Hello how are you?"))
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("  Hello how are you? "))
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("  Hello how are you ?"))
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("Hellohowareyou?"))
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces(" Hellohowareyou? "))
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("    "))
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces(""))
}
