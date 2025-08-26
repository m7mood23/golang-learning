package piscine 


func RecursiveFactorial(n int) int {
    // base case
    if n < 0 {
        return 0
    }
    if n == 0 {
        return 1
    }

    // recursive case
    return n * RecursiveFactorial(n-1)
}
