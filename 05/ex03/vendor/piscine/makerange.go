package piscine

func SplitWhiteSpaces(s string) []string {
	var ret []string
	var start, end int

	for _, r := range s {
		if r == ' ' || r == '\t' ||  r == '\n' {
			if start != end {
				ret = append(ret, s[start:end])
			}
			start = end + 1
		}
		end++
	}
	if start < end {
		ret = append(ret, s[start:])
	}
	return (ret)
}
