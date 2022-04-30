package piscine

func StrRev(s string) string {
	var ret string
	for _, r := range s {
		ret = string(r) + ret
	}
	return (ret)
}
