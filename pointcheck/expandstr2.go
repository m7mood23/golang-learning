package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {

	args := os.Args[1:]

	if len(args) != 1 {
		return
	}

	input := args[0]

	inspace := false
	firstword := false

	for _, char := range input {
		if char == ' ' {
			inspace = true
		} else {

			if inspace == true && firstword {
				z01.PrintRune(' ')
				z01.PrintRune(' ')
				z01.PrintRune(' ')
			}
			z01.PrintRune(char)
			inspace = false
			firstword = true
		}
	}
	z01.PrintRune('\n')
}
