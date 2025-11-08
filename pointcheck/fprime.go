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

	num, err := strconv.Atoi(args[0])
	if err != nil {
		return
	}

	result := ""

	// Divide by 2 as many times as possible
	for num%2 == 0 {
		result += "2*"
		num = num / 2
	}

	// Try dividing by 3, 5, 7, 9, 11, etc.
	for i := 3; i <= num; i++ {
    for num%i == 0 {
        result += strconv.Itoa(i) + "*"
        num = num / i
    }
    }

	// Remove the last "*" and print
	if len(result) > 0 {
		fmt.Println(result[:len(result)-1])
	}
}
