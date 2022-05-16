package main
import (
"fmt"
"piscine"
)
func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", piscine.Split(s, "HA"))
	s = "HAHelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", piscine.Split(s, "HA"))
	s = "HelloHAhowHAareHAyou?HA"
	fmt.Printf("%#v\n", piscine.Split(s, "HA"))
	s = "HelloHAhowHAareHAyouHA?"
	fmt.Printf("%#v\n", piscine.Split(s, "HA"))
	s = "HAHAHelloHAhowHAareHAHAyou?HAHA"
	fmt.Printf("%#v\n", piscine.Split(s, "HA"))
	s = ""
	fmt.Printf("%#v\n", piscine.Split(s, "HA"))
	s = "HAHAHelloHAhowHAareHAHAyou?HAHA"
	fmt.Printf("%#v\n", piscine.Split(s, ""))
	s = ""
	fmt.Printf("%#v\n", piscine.Split(s, ""))
}

