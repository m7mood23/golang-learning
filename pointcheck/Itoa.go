package piscine 

func Itoa(n int) string {

   // Handle zero case
   if n == 0{
       return "0"
   }

   // Check if number is negative
   negtivesign:= false  
   if n < 0 {
       negtivesign = true 
       n = -n  // Make it positive for processing
   }

   // Build the string digit by digit
   result := ""

   for n > 0{
       lastdigit:= n%10  // Extract last digit from n
       result =   string(rune( '0' + lastdigit)) + result  // Add digit to front of string
       n = n/10  // Remove last digit from n
   }

   // Add negative sign only if original number was negative
   if negtivesign == true{
       return  "-" + result 
   }else{
       return result 
   }
   return result
}
