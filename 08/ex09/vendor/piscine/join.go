package piscine

func Join(elems [] string, sep string) string {
	var ret string

	for i, s := range elems {
		if i == 0 {
			ret = s
		} else {
			ret += sep + s
		}
	}
	return (ret)
}
