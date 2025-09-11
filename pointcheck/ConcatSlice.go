package piscine 

func ConcatSlice(slice1, slice2 []int) []int {

    if len(slice1) == 0 && len(slice2) == 0 {
        return nil 
    }

    return append(slice1,slice2...)

}
