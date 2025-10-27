package piscine

func SaveAndMiss(arg  string, num int) string {
    
    // on the question(If the int is 0 or a negative number return the original string.) 
    if num <= 0 {
        return arg
    }

    //------------------------------//
    count := 0
    save := true
    result := ""
    
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
