package piscine 

import "strconv"

func ZipString(s string) string {
    
    if s == "" {// Empty string check
        return ""
    }
    
    result := "" // Final result
    count := 1   // Character counter
    

    for i := 1; i < len(s); i++ { // Compare each char with previous char
        if s[i] == s[i-1] {
            count++ // Same char - keep counting
        } else {
            result += strconv.Itoa(count) + string(s[i-1]) // Different char - add count+char to result
            count = 1 // Reset counter
        }
    }
    
   
    result += strconv.Itoa(count) + string(s[len(s)-1])  // Add the last group (loop misses it)
    
    return result
}
