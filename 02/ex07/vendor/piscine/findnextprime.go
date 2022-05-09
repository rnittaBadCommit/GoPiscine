package piscine

func isPrime(nb int) bool {
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
	return (true);
}

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return (2)
	}
	if nb % 2 == 0 {
		nb++
	}
	for ; nb > 0; nb += 2 {
		if isPrime(nb) {
			return (nb)
		}
	}
	return (0);
}
