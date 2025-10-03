package piscine

func ConcatSlice(slice1, slice2 []int) []int {


    if len(slice1) == 0 && len(slice2) == 0 {  // check if both are emty 
        return nil 
    }


    result := []int{}  // create an emapty slice intger 

    for i := range slice1{     //create for range only index with (slice1)         
        result = append(result,slice1[i])
    }

    for i := range slice2{     //same thing to (slice1) but now it's (slice1)
        result = append(result,slice2[i])
    }
    return result 
}
