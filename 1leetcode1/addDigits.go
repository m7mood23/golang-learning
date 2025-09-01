func addDigits(num int) int {
    for num >= 10 {  // Keep going while more than 1 digit
        sum := 0
        for num > 0 {
            lastdigit := num % 10
            sum = sum + lastdigit
            num = num / 10
        }
        num = sum  // Set num to the sum and check again
    }
    return num
}
