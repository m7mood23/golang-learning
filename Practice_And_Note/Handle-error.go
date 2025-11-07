package main

import (
    "errors"
    "fmt"
)

func ValidateUsername(username string) error {
    if username == "" {
        return errors.New("username cannot be empty")
    }

    if len(username) < 3 {
        return errors.New("username must be at least 3 characters")
    }

    if len(username) > 15 {
        return errors.New("username must be less than 15 characters")
    }

    return nil
}

func main() {
    // Test 1: Empty
    err := ValidateUsername("")
    if err != nil {
        fmt.Println("Error:", err)
    }

    // Test 2: Too short (2 chars)
    err = ValidateUsername("ab")
    if err != nil {
        fmt.Println("Error:", err)
    }

    // Test 3: Exactly 3 (should be valid!)
    err = ValidateUsername("abc")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("'abc' is valid!")
    }

    // Test 4: Too long
    err = ValidateUsername("verylongusername1234")
    if err != nil {
        fmt.Println("Error:", err)
    }

    // Test 5: Valid username
    err = ValidateUsername("mahmood")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("'mahmood' is valid!")
    }
}


//

// **Expected output:**

// Error: username cannot be empty
// Error: username must be at least 3 characters
// 'abc' is valid!
// Error: username must be less than 15 characters
// 'mahmood' is valid!
