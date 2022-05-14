package piscine

import "os"

func isNumeric(s string) bool {
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


func IsValidArgs(args [3]string) int {
	var tmp int

	for i, arg := range os.Args() {
		if i > 4 {
			return (0)
		} else if i > 0 {
			args[i] = arg
		}
	}
	if !isNumeric(args[0]) || !isOperator([args[1]) || !isNumeric(args[2]) {
		return (0)
	} else if args[2] == "0" {
		if args[1] == "/" {
			fmt.Println("No division by 0")
			return (0)
		} else if args[1] == "%" {
			fmt.Println("No modulo by 0")
			return (0)
		}
	}
	return (1)
}


