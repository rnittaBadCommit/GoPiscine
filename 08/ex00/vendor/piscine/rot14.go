package piscine

func Rot14(s string) string {
	var ret string
	for _, r := range s {
		if 'z' - 13 <= r && r <= 'z' {
			ret += string('a' + (14 - ('z' - r) - 1))
		} else if 'Z' - 13 <= r && r <= 'Z' {
			ret += string('A' + (14 - ('Z' - r) - 1))
		} else if ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') {
			ret += string(r + 14)
		} else {
			ret += string(r)
		}
	}
	return (ret)
}
