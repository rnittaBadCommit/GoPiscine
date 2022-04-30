package piscine

func UltimatePointOne(n ***int) {
	if n != nil && *n != nil && **n != nil {
		***n = 1
	}
}
