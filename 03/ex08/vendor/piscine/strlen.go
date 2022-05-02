package piscine

func IsAlpha(s string) bool {
	for _, r := range s {
		if !('0' <= r && r <= '9') && !('a' <= r && r <= 'z') && !('A' <= r && r <= 'Z') {
			return (false)
		}
	}
	return (true)
}
