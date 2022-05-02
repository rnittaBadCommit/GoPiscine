package piscine

func intMax() int {
	var ui uint

	return (int(^ui >> 1))
}

func BasicAtoi2(s string) int {
	var ret int

	for _, r := range s {
		if r < '0' || r > '9' {
			return (0)
		}
		if ret <= ret*10+int(r)-int('0') {
			ret = ret*10 + int(r) - int('0')
		} else {
			return (intMax())
		}
	}
	return (ret)
}
