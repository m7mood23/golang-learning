package piscine 

func FirstWord(s string) string {
    start := 0

    // 1) skip leading spaces
    for i := 0; i < len(s); i++ {
        if s[i] != ' ' {
            start = i
            break
        }
    }

    end := len(s)

    // 2) go until a space or end of string
    for i := start; i < len(s); i++ {
        if s[i] == ' ' {
            end = i
            break
        }
    }

    return s[start:end] + "\n"
}
