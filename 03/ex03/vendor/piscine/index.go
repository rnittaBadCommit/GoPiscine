package piscine

func strLen(s string) int {
	var ret int

	for range s {
		ret++
	}
	return (ret)
}

func Index(s string, toFind string) int {
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
