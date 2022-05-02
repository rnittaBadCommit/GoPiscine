package piscine

func ToUpper(s string) string {
	var ret string
	for _, r := range s {
		if 'a' <= r && r <= 'z' {
			ret += string(r + 'A' - 'a')
		} else {
			ret += string(r)
		}
	}
	return (ret)
}
