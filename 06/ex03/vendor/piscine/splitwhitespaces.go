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
	var buf []byte

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
		for ; true; {
			n, err := os.Stdin.Read(buf)
			if err != nil {
				os.Exit(1)
			} else if n != 0 {
				printStr(string(buf))
			}
		}
	}
/*
	if len == 1 {
		for ;true; {
			data, err := os.Read("/dev/stdin")
			if err != nil {
				printStr("something wrong\n")
			} else {
				printStr(string(data))
			}
		}
	}

	*/
}
