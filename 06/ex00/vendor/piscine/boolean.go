package piscine

import "os"

func LengthOfArg() int {
	var ret int
	for range os.Args {
		ret++;
	}
	return (ret - 3)
}
