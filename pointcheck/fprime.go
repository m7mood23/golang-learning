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
	
	num, err := strconv.Atoi(input)
	if err != nil || num <= 1 {
		return
	}

	result := ""
	for i := 2; i <= num; i++ {
		for num%i == 0 {
			if result != "" {
				result += "*"
			}
			result += strconv.Itoa(i)
			num = num / i
		}
	}
	fmt.Println(result)
}
