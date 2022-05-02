package piscine

func Sqrt(nb int) int {
	var i int

	if nb <= 0 {
		return (0)
	}
	for i = 1; i < nb / i; i++ {
	}
	if i * i == nb {
		return (i)
	} else {
		return (0)
	}
}
