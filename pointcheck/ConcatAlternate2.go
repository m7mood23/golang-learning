package piscine

func ConcatAlternate(slice1, slice2 []int) []int {
     if len(slice1) == 0 && len(slice2) == 0 {
        return nil
    }
    
    if len(slice2) > len(slice1) {
        slice1, slice2 = slice2, slice1
    }
    
    result := []int{}
    
    for i := range slice1 {
        result = append(result, slice1[i])
        if i < len(slice2) {
            result = append(result, slice2[i])
        }
    }
    return result
}
