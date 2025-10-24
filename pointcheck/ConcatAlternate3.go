package piscine

func ConcatAlternate(slice1, slice2 []int) []int {

    if len(slice2) > len(slice1) {
        slice1, slice2 = slice2, slice1
    }
    
    var result []int  // Use var instead of := []int{}   
    
    for i := range slice1 {
        result = append(result, slice1[i])
        if i < len(slice2) {
            result = append(result, slice2[i])
        }
    }
    return result
}
