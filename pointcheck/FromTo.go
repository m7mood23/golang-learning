package piscine 

import "strconv"

func FromTo(from int, to int) string {

   // Check if numbers are outside valid range (0-99)
   if from < 0 || from > 99 || to < 0 || to > 99 {
       return "Invalid\n"  // Return error message
   }

   result := ""  // Empty string to build our answer

   // Check if counting up or down
   if from <= to {
       // Count up: from 1 to 5 = 1,2,3,4,5
       for i:= from;i <= to; i++{
           // Add comma before each number except the first
           if result != ""{
               result += ", "
           }
           // Add leading zero for single digits (5 becomes 05)
         //"Prepend a 0 to any number that is less than 10."
           if i < 10 {
               result = result + "0" + strconv.Itoa(i)
           }else{
               result = result + strconv.Itoa(i)  // Convert number to string
           }
       }
   }else {
       // Count down: from 5 to 1 = 5,4,3,2,1
       for i := from; i >= to; i--{
           // Add comma before each number except the first
           if result != ""{
               result = result + ", "
           }
           // Add leading zero for single digits (5 becomes 05)
         //"Prepend a 0 to any number that is less than 10."
           if i < 10 {
               result = result  + "0" + strconv.Itoa(i)
           }else {
               result = result + strconv.Itoa(i)  // Convert number to string
           }
       }
   }

   return result + "\n"  // Add newline at end and return
}
