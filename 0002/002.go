package main

import "fmt"

func main() {
	// Initialize the first two terms of the Fibonacci sequence
	a, b := 1, 2
	// Initialize the sum of the even-valued terms to 0
	sum := 0

	// Loop through the terms of the Fibonacci sequence
	for b <= 4000000 {
		// Check if the current term is even
		if b%2 == 0 {
			// If it is, add it to the sum
			sum += b
		}
		// Calculate the next term in the sequence
		a, b = b, a+b
	}

	// Print the sum of the even-valued terms
	fmt.Println(sum)
}
