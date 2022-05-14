package piscine

func IsPrime(nb int) bool {
	if nb == 2 {
		return (true)
	} else if nb < 2 || nb % 2 == 0 {
		return (false)
	}
	for i := 3; i <= nb / i; i += 2 {
		if nb % i == 0 {
			return (false)
		}
	}
	return (true)
}
