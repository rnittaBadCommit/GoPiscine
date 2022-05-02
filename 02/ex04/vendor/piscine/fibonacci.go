package piscine

func intMax() int {
	var max ui

	return (^max>>1)
}

func Fibonacci(index int) int {
	if index < 0 {
		return (-1)
	} else if index == 0 {
		return (0)
	} else if index <= 2 {
		return (1)
	} else {
		tmp1 := Fibonacci(index - 1)
		tmp2 := Fibonacci(index - 2)
		max := intMax()
		if tmp1 < max() - tmp2 {
			return (Fibonacci(index - 1) + Fibonacci(index - 2))
		} else {
			return (max)
		}
	}
}
