package piscine

func Compare(a, b string) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i, _ := range a {
		for j, _ := range a {
			if i < j && f(a[i], a[j]) > 0 {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
