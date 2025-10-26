package piscine

func SaveAndMiss(s string, num int) string {
    if num <= 0 {
        return s
    }
    
    result := ""
    save := true
    count := 0
    
    for _, char := range s {
        if save == true {
            result += string(char)
        }
        
        count++
        
        if count == num {
            if save == true {
                save = false
            } else {
                save = true
            }
            count = 0
        }
    }
    
    return result
}
