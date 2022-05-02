package piscine

func SortIntegerTable(table []int) {
	var i, j int

	for i, _ = range table {
		for j, _ = range table {
			if i < j && table[i] > table[j] {
				table[i], table[j] = table[j], table[i]
			}
		}
	}
}
