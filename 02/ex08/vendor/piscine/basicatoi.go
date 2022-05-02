package piscine

func intMax() int {
	var ui uint

	return (int(^ui >> 1))
}

func BasicAtoi(s string) int {
	var ret int

	for _, r := range s {
		if ret <= ret*10+int(r)-int('0') {
			ret = ret*10 + int(r) - int('0')
		} else {
			return (intMax())
		}
	}
	return (ret)
}
