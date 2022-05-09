package piscine

func intMax() int {
	var max uint

	return (int(^max >> 1))
}

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return (0)
	}
	ret := 1
	max := intMax()
	for i := nb; i > 1; i-- {
		if ret > max/i {
			return (0)
		}
		ret *= i
	}
	return (ret)

}
