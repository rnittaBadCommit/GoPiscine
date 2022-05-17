package piscine

func Compact(ptr *[]string) int {
	var num int
	for _, s := range *ptr {
		if s != "" {
			num++
		}
	}
	new_s := make([]string, num)
	num = 0
	for _, s := range *ptr {
		if s != "" {
			new_s[num] = s
			num++
		}
	}
	*ptr = new_s
	return num
}
