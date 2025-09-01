func reverseString(s []byte) {
    result := []byte{}  // Use []byte instead of string
    
    for i := len(s) - 1; i >= 0; i-- {
        result = append(result, s[i])  // Collect bytes in reverse
    }
    
    // Copy result back to original slice
    for i := 0; i < len(s); i++ {
        s[i] = result[i]
    }
}
