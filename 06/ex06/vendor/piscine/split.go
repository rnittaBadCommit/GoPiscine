package piscine


func strLen(s string) int {
	var ret int

	for range s {
		ret++
	}
	return (ret)
}

func index(s string, toFind string) int {
	length1 := strLen(s)
	length2 := strLen(toFind)
	if length2 > length1 {
		return (-1)
	}
	str := []rune(s)
	for ret, _ := range str {
		newstr := str[ret:ret + length2]
		if string(newstr) == toFind {
			return (ret)
		} else if ret == length1 - length2 {
			break
		}
	}
	return (-1)
}

func Split(s, sep string) []string {
	var ret []string
	var start int
	sp_len := strLen(sep)

	if s == "" {
		return (nil)
	}
	if sep == "" {
		ret = append(ret, s[0:])
		return (ret)
	}
	for ; true; {
		end := index(s, sep)
		if end == -1 {
			break
		} else if end != 0{
			ret = append(ret, s[:end])
		}
		s = s[end + sp_len:]
	}
	if s[start:] != "" {
		ret = append(ret, s[start:])
	}
	return (ret)
}


