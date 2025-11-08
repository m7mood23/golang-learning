package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	str1 := os.Args[1]
	str2 := os.Args[2]
	printed := ""

	for _, char := range str1 {
		if seen(str2, char) && !seen(printed, char) {
			z01.PrintRune(char)
			printed += string(char)
		}
	}
	z01.PrintRune('\n')
}

//----------func seen ------------------------------
func seen(s string, c rune) bool {
	for _, char := range s {
		if char == c {
			return true
		}
	}
	return false
}
