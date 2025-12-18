package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	args := os.Args[1:]
    
	if len(args) == 0 {
		return
	}
	
	for _, input := range args {
		for i, char := range input {
			
			// Lowercase letter - make uppercase if last
			if char >= 'a' && char <= 'z' {
				if i == len(input)-1 {
					char -= 32
				}
				if i < len(input)-1 {
					if input[i+1] == ' ' {
						char -= 32
					}
				}
			}
			
			// Uppercase letter - make lowercase if NOT last
			if char >= 'A' && char <= 'Z' {
				if i != len(input)-1 {
					if input[i+1] != ' ' {
						char += 32
					}
				}
			}
			
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
