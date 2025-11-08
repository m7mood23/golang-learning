package main

import ("os"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	num := atoi(os.Args[1])
	if num <= 0 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	sum := 0
	for i := 2; i <= num; i++ {
		if isPrime(i) {
			sum += i
		}
	}

	printNum(sum)
	z01.PrintRune('\n')
}
//------------------------------------------------------------
//1-func isPrime
func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

//------------------------------------------------------------
//2-func atoi 
func atoi(s string) int {
	num := 0
	for _, char := range s {
		if char < '0' || char > '9' {
			return 0
		}
		num = num*10 + int(char-'0')
	}
	return num
}
//------------------------------------------------------------
//3-func printNum
func printNum(n int) {
	if n >= 10 {
		printNum(n / 10)
	}
	z01.PrintRune(rune(n%10 + '0'))
}
