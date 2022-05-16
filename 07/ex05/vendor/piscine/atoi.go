package piscine

func intMax() uint {
	var tmp uint

	return (^tmp>>1)
}

func Atoi(s string) (int, int) {
	var ret uint
	var isNegative bool
	max := intMax()

	if s[0] == '-' {
		s = s[1:]
		isNegative = true
	}
	for _, r := range s {
		if isNegative && ret > (max + 1 - uint(r - '0')) / 10 {
			return 0, -1
		} else if isNegative == false && ret > (max - uint(r - '0')) / 10 {
			return 0, -1
		}
		ret = ret * 10 + uint(r - '0')
	}
	if isNegative {
		return -1 * int(ret), 0
	} else {
		return int(ret), 0
	}
}
