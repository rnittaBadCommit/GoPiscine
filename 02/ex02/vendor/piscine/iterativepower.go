package piscine

func intMax() int {
	var max uint

	return (int(^max >> 1))
}

func IterativePower(nb int, power int) int {
	if power < 0 {
		return (0)
	}
	ret := 1
	for ; power > 0; power-- {
		ret *= nb
	}
	return (ret)
}
