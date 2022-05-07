package piscine

func intMax() int {
	var max uint

	return (int(^max >> 1))
}

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return (0)
	} else if nb <= 1 {
		return (1)
	}
	tmp := RecursiveFactorial(nb - 1)
	if tmp > intMax()/nb {
		return (0)
	} else {
		return (tmp * nb)
	}
}
