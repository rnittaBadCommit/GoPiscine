package piscine

func BasicJoin(elems [] string) string {
	var ret string

	for _, s := range elems {
		ret += s
	}
	return (ret)
}
