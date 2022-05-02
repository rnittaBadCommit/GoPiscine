package piscine

func Capitalize(s string) string {
	var ret string
	var flag bool = true

	for _, r := range s {
		if 'a' <= r && r <= 'z' {
			if flag {
				ret += string(r + 'A' - 'a')
			} else {
				ret += string(r)
			}
			flag= false
		} else if 'A' <= r && r <= 'Z' {
			if flag {
				ret += string(r)
			} else {
				ret += string(r + - 'A' + 'a')
			}
			flag= false
		} else if '0' <= r && r <= '9' {
			ret += string(r)
			flag = false
		} else {
			ret += string(r)
			flag = true
		}
	}
	return (ret)
}

