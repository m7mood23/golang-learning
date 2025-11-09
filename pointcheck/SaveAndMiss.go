package piscine

func SaveAndMiss(arg  string, num int) string {
    
    // Q: if num <= 0 -→ return original string    if num <= 0 {
        return arg
    }

    
    result := ""
    count := 0
    save := true
    
   
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
