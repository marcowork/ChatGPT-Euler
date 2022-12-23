package main

import "fmt"

func main() {
	sum := 0

	// Iterate through all numbers below 1000
	for i := 1; i < 1000; i++ {
		// If the number is a multiple of 3 or 5, add it to the sum
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	fmt.Println(sum)
}
