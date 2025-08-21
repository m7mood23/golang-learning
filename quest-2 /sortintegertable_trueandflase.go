package piscine 

func SortIntegerTable(table []int) {
    swapped := true // Start with 'true' to run the loop at least once

    for swapped {
        swapped = false // Assume no swaps will happen in this pass
        
        for i := 0; i < len(table)-1; i++ {
            if table[i] > table[i+1] {
                // Swap the numbers
                temp := table[i]
                table[i] = table[i+1]
                table[i+1] = temp
                
                swapped = true // A swap happened, so we need to do another pass
            }
        }
    }
}
