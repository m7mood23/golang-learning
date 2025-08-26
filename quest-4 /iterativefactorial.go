package piscine 

func IterativeFactorial(nb int) int {
	//if nb is less than 0 return 0 
	if nb < 0 {
		return 0
	}

	// this is how we calualte the iterativefactoiral 
	result := 1 
	for i := 1 ; i <= nb ; i++ {
		result = result * i 
	}
	return result 
}

// "5! is the notation."

// "5 factorial" is how you say it.

// 5 x 4 x 3 x 2 x 1 is the calculation that the notation represents. 


// An iterative function uses a loop (like a for loop) to repeatedly do an action until a condition is met.
