package piscine

func Swap(a *int, b *int) {
 // Step 1: Save the value that 'a' points to in a temporary variable
   // This prevents us from losing the original value of 'a'
  // and also same for b 
	temp1 := *a
	temp2 := *b

	*a = temp2
	*b = temp1

}
