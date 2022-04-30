package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Atoi("12345"))
	fmt.Println(piscine.Atoi("0000000012345"))
	fmt.Println(piscine.Atoi("012 345"))
	fmt.Println(piscine.Atoi("Hello World!"))
	fmt.Println(piscine.Atoi("+1234"))
	fmt.Println(piscine.Atoi("-1234"))
	fmt.Println(piscine.Atoi("++1234"))
	fmt.Println(piscine.Atoi("--1234"))
	fmt.Println(piscine.Atoi("9223372036854775807"))
	fmt.Println(piscine.Atoi("9223372036854775808"))
	fmt.Println(piscine.Atoi("-9223372036854775807"))
	fmt.Println(piscine.Atoi("-9223372036854775808"))
	fmt.Println(piscine.Atoi("-9223372036854775809"))
	fmt.Printf("\n")
	fmt.Println(piscine.Atoi("92233720368547758070"))
	fmt.Println(piscine.Atoi("-92233720368547758070"))
}
