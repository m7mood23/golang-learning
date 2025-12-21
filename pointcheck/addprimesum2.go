package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	
	if len(args) != 1 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	// Convert string to number
	num := 0
	for _, char := range args[0] {
		if char >= '0' && char <= '9' {
			num = num*10 + int(char-'0')
		} else {
			z01.PrintRune('0')
			z01.PrintRune('\n')
			return
		}
	}

	// Find sum of all primes up to num
	sum := 0
	for i := 2; i <= num; i++ {
		isPrime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			sum += i
		}
	}

	// Convert number to string
	result := ""
	if sum == 0 {
		result = "0"
	} else {
		temp := sum
		for temp > 0 {
			digit := temp % 10
			result = string(rune(digit+'0')) + result
			temp = temp / 10
		}
	}

	// Print result
	for _, char := range result {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

