package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		return
	}

	str1 := args[0]
	str2 := args[1]
	
    
    i := 0
    
	for _, char := range str2 {
		if i < len(str1) {
			if char == rune(str1[i]) {
				i++
			}
		}
	}

	if i == len(str1) {
		for _, char := range str1 {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')  // ← MOVED INSIDE!
	}
}
