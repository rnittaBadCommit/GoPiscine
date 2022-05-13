package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return (make([]int, 0, 0))
	}
	ret := make([]int, max - min, max - min)
	for i := min; i < max; i++ {
		ret[i - min] = i
	}
	return (ret)
}
