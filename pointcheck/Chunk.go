package piscine

import "fmt"

func Chunk(slice []int, size int) {
    if size == 0 {         //If size is 0, print empty line and stop.
        fmt.Println()
        return
    }
    
    result := [][]int{}  //(a slice of slices). [[1,2,3],[4,5,6],[7,8]]
    
    for i := 0; i < len(slice); i += size {
        end := i + size
        
        if end > len(slice) {
            end = len(slice)
        }
        
        chunk := slice[i:end]
        result = append(result, chunk)
    }
    
    fmt.Println(result)
}
