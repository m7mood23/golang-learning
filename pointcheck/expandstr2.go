package main

import (
    "os"
    "github.com/01-edu/z01"
)

func main() {
    args:= os.Args[1:]

    if len(args) != 1{
        return 
    }
    
    s := args[0]
    inSpace := true
    isFirstWord  := false
    
    for _, char := range s {
        if char == ' ' {
            inSpace = true
        } else {
            if inSpace == true && isFirstWord  {
                z01.PrintRune(' ')
                z01.PrintRune(' ')
                z01.PrintRune(' ')
            }
            z01.PrintRune(char)
            inSpace = false
            isFirstWord  = true
        }
    }
    z01.PrintRune('\n')
}
