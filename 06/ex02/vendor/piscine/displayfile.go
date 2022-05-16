package piscine

import (
	"os"
	"ft"
)

func printStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
}

func DisplayFile() {
	var len int
	var fileName string

	for _, fileName = range os.Args {
		len++;
	}
	if len == 2 {
		data, err := os.ReadFile(fileName)
		if err != nil {
			printStr("something wrong\n")
		} else {
			printStr(string(data))
		}
	} else if len == 1 {
		printStr("File name missing\n")
	} else {
		printStr("Too many arguments\n")
	}
}