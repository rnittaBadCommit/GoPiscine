package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	aLen := SliceLen(a)

	for i := 0; i < aLen - 1; i++ {
		if f(a[i], a[i + 1]) > 0 {
			return (false)
		}
	}
	return (true)
}
