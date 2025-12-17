package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	input := args[0]

	// Convert to number
	num, err := strconv.Atoi(input)
	if err != nil || num <= 1 {
		return
	}

	factors := []int{}
	
	for i := 2; i <= num; i++ {
		// Keep dividing by i while it divides evenly
		for num%i == 0 {
			factors = append(factors, i)
			num = num / i
		}
	}

	// Print factors with * between them
	for i, factor := range factors {
		if i > 0 {
			fmt.Print("*")
		}
		fmt.Print(factor)
	}
	fmt.Println()
}

fprime 
