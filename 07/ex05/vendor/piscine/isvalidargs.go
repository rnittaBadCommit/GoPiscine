package piscine

import (
	"fmt"
	"os"
)

func isNumeric(s string) bool {
	if s == "" || s == "-" {
		return false
	}
	if s[0] == '-' {
		s = s[1:]
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return (false)
		}
	}
	return (true)
}

func isOperator(s string) bool {
	if s == "+" || s == "-" || s == "/" || s == "*" || s == "%" {
		return (true)
	} else {
		return (false)
	}
}


func IsValidArgs(args [3]string) bool {
	var argsLen int

	for _, arg := range os.Args {
		if argsLen > 4 {
			return (false)
		} else if argsLen > 0 {
			args[argsLen - 1] = arg
		}
		argsLen++
	}
	if argsLen != 4 {
		return (false)
	}

	if !isNumeric(args[0]) || !isOperator(args[1]) || !isNumeric(args[2]) {
		return (false)
	} else if args[2] == "0" {
		if args[1] == "/" {
			fmt.Println("No division by 0")
			return (false)
		} else if args[1] == "%" {
			fmt.Println("No modulo by 0")
			return (false)
		}
	}
	return (true)
}


