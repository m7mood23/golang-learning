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
			
			// Lowercase letter
			if char >= 'a' && char <= 'z' {
				if i == len(input)-1 {
					char -= 32
				} else {
					if input[i+1] == ' ' {
						char -= 32
					}
				}
			}
			
			// Uppercase letter
			if char >= 'A' && char <= 'Z' {
				if i == len(input)-1 {
					// Keep uppercase (do nothing)
				} else {
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
