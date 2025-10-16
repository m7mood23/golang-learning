package main

import (
    "os"
    "github.com/01-edu/z01"
)

func main() {
    if len(os.Args) != 2 {
        return
    }
    
    s := os.Args[1]
    inSpace := true
    first := true
    
    for _, char := range s {
        if char == ' ' {
            inSpace = true
        } else {
            if inSpace && !first {
                z01.PrintRune(' ')
                z01.PrintRune(' ')
                z01.PrintRune(' ')
            }
            z01.PrintRune(char)
            inSpace = false
            first = false
        }
    }
    z01.PrintRune('\n')
}
