package main

import (
    "os"
    "github.com/01-edu/z01"
)

func main() {
    args := os.Args[1:]
    
    if len(args) != 2 {
        z01.PrintRune('\n')
        return
    }
    
    s1 := args[0]
    s2 := args[1]
    result := ""

    //for range s1
    for _, char := range s1 {
        found := false
        for _, check := range result {
            if char == check {
                found = true
            }
        }
        if found == false {
            result += string(char)
        }
    }

    //for range s2 
    for _, char := range s2 {
        found := false
        for _, check := range result {
            if char == check {
                found = true
            }
        }
        if found == false {
            result += string(char)
        }
    }
    
    for _, char := range result {
        z01.PrintRune(char)
    }
    z01.PrintRune('\n')
}
