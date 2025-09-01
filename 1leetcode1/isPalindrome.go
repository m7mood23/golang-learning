func isPalindrome(x int) bool {
        String_int:= strconv.Itoa(x)
            //start 0 , end 1 becuase 3/2 = 1.5 to 1 that's how computer see 
        for i:= 0; i< len(String_int)/2 ;i++{
            if String_int[i] != String_int[len(String_int)-1 - i]{
                return false 
            }
        }

    return true 
}
