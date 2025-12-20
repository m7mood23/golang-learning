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
	
	s1 := args[0]
	s2 := args[1]
	i := 0
	
	for _, char := range s2 {
		if i < len(s1) {
			if char == rune(s1[i]) {
				i++
			}
		}
	}
	
	if i == len(s1) {
		for _, char := range s1 {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
