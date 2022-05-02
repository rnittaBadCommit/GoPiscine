package piscine

func IsPrintable(s string) bool {
	for _, r := range s {
		if r < ' ' || r == 127 {
			return (false)
		}
	}
	return (true)
}
