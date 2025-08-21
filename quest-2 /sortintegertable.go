package piscine 

func SortIntegerTable(table []int) {
    // Outer loop: controls the number of passes
    for i := 0; i < len(table)-1; i++ {
        // Inner loop: compares adjacent elements and swaps them if needed
        for j := 0; j < len(table)-i-1; j++ {
            // Check if the current element is greater than the next one
            if table[j] > table[j+1] {
                // If so, swap them
                // A classic three-line swap
                temp := table[j]
                table[j] = table[j+1]
                table[j+1] = temp
            }
        }
    }
}
