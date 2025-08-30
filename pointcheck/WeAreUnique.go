package piscine

import "strings"

func WeAreUnique(str1, str2 string) int {
    if str1 == "" && str2 == "" {
        return -1
    }
    
    count := 0
    counted := ""  // Track what we already counted
    
    // Check characters in str1
    for _, char := range str1 {
        charStr := string(char)
        // Check if char is NOT in str2 AND we haven't counted it yet
        if !strings.Contains(str2, charStr) && !strings.Contains(counted, charStr) {
            count++
            counted = counted +  charStr  // Remember we counted t his
        }
    }
    
    // Check characters in str2  
    for _, char := range str2 {
        charStr := string(char)
        // Check if char is NOT in str1 AND we haven't counted it yet
        if !strings.Contains(str1, charStr) && !strings.Contains(counted, charStr) {
            count++
            counted += charStr  // Remember we counted this
        }
    }
    
    return count
}
