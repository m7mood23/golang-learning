package piscine

func SaveAndMiss(arg  string, num int) string {
    if num <= 0 {
        return arg
    }
    
    result := ""
    save := true
    count := 0
    
    for _, char := range arg  {
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
