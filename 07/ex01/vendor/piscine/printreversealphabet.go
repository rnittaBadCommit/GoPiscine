package piscine

func Map(f func(int) bool, a []int) []bool {
	var ret []bool

	for _, i := range a {
		ret = append(ret, f(i))
	}
	return (ret)
}
