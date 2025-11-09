package main

import ("os"
	"github.com/01-edu/z01"
)

func main() {
	
    args:= os.Args[1:]
    
    if len(args) != 2{
        return 
    }

	s1 := args[0]
	s2 := args[1]
	keep := ""

	for _, char := range s1 {
		if seen(s2, char) && !seen(keep, char) {
			z01.PrintRune(char)
			keep += string(char)
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
