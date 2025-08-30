package piscine 

func FindPrevPrime(nb int) int {
   for nb > 1 {  // Keep trying numbers until we reach 1
       isPrime := true  // Assume this number is prime
       
       // Check if nb can be divided by any number from 2 to nb-1
       for i := 2; i < nb; i++ {  
           if nb%i == 0 {  // If nb divides perfectly by i
               isPrime = false  // Then nb is not prime
           }
       }
       
       if isPrime == true {  // If no divisors were found
           return nb  // Return this prime number
       }else{
           nb--  // Try the next smaller number
       }
   }
   return 0  // No prime found (when nb becomes 1 or less)
}
