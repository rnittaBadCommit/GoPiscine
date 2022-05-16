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
		if len > 1 {
			data, err := os.ReadFile(fileName)
			if err != nil {
				printStr("something wrong\n")
			} else {
				printStr(string(data))
			}
		}
	}
	if len == 1 {
		for ;true; {
			data, err := os.ReadFile("/dev/stdin")
			if err != nil {
				printStr("something wrong\n")
			} else {
				printStr(string(data))
			}
		}
	}
}