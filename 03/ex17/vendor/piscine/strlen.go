package piscine

func is_BaseValid(base []rune) bool {
	var count int
	for i, r := range base {
		if r == '+' || r == '-' {
			return (false)
		}
		for j, r2 := range base {
			if i != j && r == r2 {
				return (false)
			}
		}
		count++
	}
	return (count > 1)
}

func strLen(s []rune) int {
	var ret int

	for range s {
		ret++
	}
	return ret
}

func AtoiBase(s string, base string) int {
	rbase := []rune(base)

	if is_BaseValid(rbase) == false {
		return (0)
	}

	var ret int
	baseLen := strLen(rbase)
	for _, r := range s {
		is_valid := false
		var count int
		for _, r2 := range rbase {
			if r == r2 {
				is_valid = true
				break
			}
			count++
		}
		if is_valid == false {
			return (ret)
		} else {
			ret = ret * baseLen + count
		}
	}
	return (ret)
}
