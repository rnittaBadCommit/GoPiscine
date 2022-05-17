package piscine

func ActiveBits(n int) int {
	var ret int
	var un uint

	un = uint(n)
	for ; un != 0; un = un >> 1 {
		if un & 1 == 1 {
			ret++
		}
	}
	return ret
}
