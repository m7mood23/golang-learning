package main

import (
    "os"
    "github.com/01-edu/z01"
)

func main() {
    if len(os.Args) != 3 { // Check if exactly 2 arguments provided
        return // Exit if wrong number of arguments
    }
    
    first := os.Args[1]  // Get first string to match
    second := os.Args[2] // Get second string to search in
    
    lookingFor := 0 // Track which letter we need to find next
    
    for i := 0; i < len(second); i++ { // Go through each letter in second string
        if lookingFor < len(first) && second[i] == first[lookingFor] { // Check if current letter matches what we need
            lookingFor++ // Move to next letter we need to find
        }
    }

    if lookingFor == len(first) { // Check if we found all letters in order
        for _, char := range first { // Loop through each character in first string
            z01.PrintRune(char) // Print the character 
        }
        z01.PrintRune('\n') // Print newline to end the output
    }
}
