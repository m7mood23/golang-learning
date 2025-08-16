package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {

	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {

		z01.PrintRune('-')
		n = -n
	}
	result := []rune{}
	for n > 0 {

		lastdigit := n % 10

		newchar := rune('0' + lastdigit)

		result = append(result, newchar)
		n = n / 10
	}

	for i := len(result) - 1; i >= 0; i-- {
		z01.PrintRune(result[i])
	}

}
