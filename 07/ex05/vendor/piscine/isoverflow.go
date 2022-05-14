package piscine

func IsOverFlow(int i1, ope string, i2 int) bool {
	if ope == "+" {
		if i1 > 0 && i2 > 0 {
			return (i1 + i2 <= 0)
		} else if i1 < 0 && i2 < 0 {
			return (i1 + i2 >= 0)
		}
	} else if ope == "-" {
		if i1 > 0 && i2 < 0 {
			return (i1 - i2 <= 0)
		} else if i1 < 0 && i2 > 0 {
			return (i1 - i2 >= 0)
		}
	} else if ope == "/" {
		return (i2 == -1 && i1 / i2 == i1)
	} else if ope == "*" && i1 != 0 && i2 != 0 {
		if (i1 == -1 || i2 == -1) && (i1 * i2 == i1 || i1 * i2 == i2) {
			return (true)
		}
		return (i1 * i2 / i2 != i1)
	}
	return (false)
}
