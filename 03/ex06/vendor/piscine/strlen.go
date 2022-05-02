package piscine

func ToLower(s string) string {
	var ret string
	for _, r := range s {
		if 'A' <= r && r <= 'Z' {
			ret += string(r + 'a' - 'A')
		} else {
			ret += string(r)
		}
	}
	return (ret)
}
