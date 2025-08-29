package piscine

func IsCapitalized(s string) bool {
    if s == "" {
        return false  // empty string = false
    }
    
    checkfirstletter := true  // start checking first letter
    
    for _, char := range s {  // loop through each character
        if checkfirstletter {  // if we need to check this char 
            if char >= 'a' && char <= 'z' {  // if lowercase .     * If character is lowercase, return false immediately
                return false  // fail immediately
            }
            checkfirstletter = false  // stop checking this word
        } else if char == ' ' {  // if space found
            checkfirstletter = true  // check next char
        }
    }
    
    return true  // all words passed
}
