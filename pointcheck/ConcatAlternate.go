package piscine 

func ConcatAlternate(slice1, slice2 []int) []int {
   
    if len(slice1) == 0 && len(slice2) == 0 {  // Check if both slices are empty, return nil if true
        return nil
    }
    
    result := []int{} // Create empty slice to store final result
    

    first := slice1 // Set slice1 as default first slice
    second := slice2  // Set slice2 as default second slice

    
    if len(slice2) > len(slice1) { // If slice2 is longer, swap them
        first = slice2
        second = slice1
    }

    
  
    minLen := len(second)   // Get length of shorter slice (how many pairs we can make)
    
   
    for i := 0; i < minLen; i++ {   // Loop to alternate elements from both slices

        result = append(result, first[i])    // Add element from longer slice
        result = append(result, second[i])  // Add element from shorter slice
    }
    
  
    for i := minLen; i < len(first); i++ {  // Add leftover elements from longer slice
        result = append(result, first[i])
    }
    
    return result     // Return the final alternated slice
}
