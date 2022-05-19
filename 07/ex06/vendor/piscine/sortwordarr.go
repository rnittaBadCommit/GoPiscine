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

func SortWordArr(a []string) {
	for i, _ := range a {
		for j, _ := range a {
			if i < j && a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
