package piscine

func Gcd(a, b uint) uint {

	if b == 0 {
		return a
	}

	return Gcd(b, a%b)

}
