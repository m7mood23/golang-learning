package main

import "github.com/01-edu/z01"

func CheckPassword(password string) {

	hasUpper := false
	hasLower := false
	hasDigit := false

	for _, char := range password {

		if char >= 'a' && char <= 'z' {

			hasLower = true
		}

		if char >= 'A' && char <= 'Z' {

			hasUpper = true
		}

		if char >= '0' && char <= '9' {
			hasDigit = true

		}
	}

	if hasUpper == true && hasLower == true && hasDigit == true {
		z01.PrintRune('s')
	} else {
		z01.PrintRune('w')
	}
	z01.PrintRune('\n')
}

func main() {
	CheckPassword("Hello123")   // Should print S (has upper, lower, digit)
	CheckPassword("hello123")   // Should print W (no uppercase)
	CheckPassword("HELLO123")   // Should print W (no lowercase)
	CheckPassword("HelloWorld") // Should print W (no digit)
	CheckPassword("aB3")        // Should print S (has all three)
}
