package main

import (
	"piscine"
	"ft"
)

func main() {
	result := piscine.Rot14("Hello! How are You?")

	for _, r := range result {
		ft.PrintRune(r)
	}
	tmp := ("\n===================\n")
	for _, r := range tmp {
		ft.PrintRune(r)
	}
	result = piscine.Rot14("abcdefghijklmnopqrstuvwxyz ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	for _, r := range result {
		ft.PrintRune(r)
	}
}
