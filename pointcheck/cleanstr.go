package main

import (
    "os"
    "github.com/01-edu/z01"
)

func main() {
    if len(os.Args) != 2 {
        z01.PrintRune('\n')
        return
    }

    s := os.Args[1]

    if s == "" {
        z01.PrintRune('\n')
        return
    }

    words := []string{} 
    word := ""

    for _, char := range s {
        if char == ' ' || char == '\t' {
            if word != "" {
                words = append(words, word)
                word = ""
            }
        } else {
            word += string(char)
        }
    }

    if word != "" {
        words = append(words, word)
    }

    for i, w := range words {
        for _, char := range w {
            z01.PrintRune(char)
        }
        if i < len(words)-1 {
            z01.PrintRune(' ')
        }
    }
    z01.PrintRune('\n')
}
