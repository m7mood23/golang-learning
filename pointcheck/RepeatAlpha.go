package piscine

// Formula reminder:
// For lowercase: position = int(char - 'a' + 1)
// For uppercase: position = int(char - 'A' + 1)

func RepeatAlpha(s string) string {
    result := ""

    for _, char := range s {
        if char >= 'a' && char <= 'z' {
            // Repeat lowercase letters (a=1, b=2,...)
            repeat := int(char - 'a' + 1)
            for i := 0; i < repeat; i++ {
                result += string(char)
            }
        } else if char >= 'A' && char <= 'Z' {
            // Repeat uppercase letters (A=1, B=2,...)
            repeat := int(char - 'A' + 1)
            for i := 0; i < repeat; i++ {
                result += string(char)
            }
        } else {
            // Keep non-letters as-is
            result += string(char)
        }
    }

    return result
}
