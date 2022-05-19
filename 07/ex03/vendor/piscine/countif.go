package piscine

func CountIf(f func(string) bool, a []string) int {
	var ret int

	for _, s := range a {
		if f(s) {
			ret++
		}
	}
	return (ret)
}
