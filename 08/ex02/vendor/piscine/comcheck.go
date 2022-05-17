package piscine

import "os"

func Comcheck() {
	for _, arg := range os.Args {
		if arg == "42" || arg == "piscine" || arg == "piscine 42" {
			printStr("Alert!!!\n")
			return
		}
	}
}
