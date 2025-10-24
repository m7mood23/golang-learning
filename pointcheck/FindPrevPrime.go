package piscine

func FindPrevPrime(nb int) int {
	for nb > 1 {
		isPrime := true
		for i := 2; i < nb; i++ {
			if nb%i == 0 {
				isPrime = false
			}
		}
		if isPrime == true {
			return nb
		} else {
			nb--
		}
	}
	return 0
}
