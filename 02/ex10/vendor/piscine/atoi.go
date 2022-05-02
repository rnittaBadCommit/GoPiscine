package piscine

func intMax() uint {
	var ui uint

	return (^ui >> 1)
}

func Atoi(s string) int {
	var ret uint
	var is_overflow bool
	var max uint

	max = intMax()
	s2 := s
	if s[0] == '+' || s[0] == '-' {
		s = s[1:]
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return (0)
		}
		if ret <= ret*10+uint(r)-uint('0') {
			ret = ret*10 + uint(r) - uint('0')
		} else {
			is_overflow = true
		}
	}
	if is_overflow || ret > max {
		if s2[0] == '-' {
			return (int(max) + 1)
		} else {
			return (int(max))
		}
	}
	if s2[0] == '-' {
		return (int(ret) * -1)
	}
	return (int(ret))
}
