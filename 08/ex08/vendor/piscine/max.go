package piscine

func Max(a []int) int {
	var ret int
	for i, n := range a {
		if i == 0 {
			ret = n
		} else if ret < n {
			ret = n
		}
	}
	return ret
}
