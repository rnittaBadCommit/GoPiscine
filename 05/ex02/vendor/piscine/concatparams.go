package piscine

func ConcatParams(args []string) string {
	var ret string

	for i, s := range args {
		if i == 0 {
			ret = s
		} else {
			ret += string('\n') + s
		}
	}
	return (ret)
}
