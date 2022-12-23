package main

import "fmt"

func main() {
	// Start with the smallest number that is evenly divisible by all of the numbers from 1 to 10
	n := 2520
	for {
		// Check if n is evenly divisible by all of the numbers from 11 to 20
		found := true
		for i := 11; i <= 20; i++ {
			if n%i != 0 {
				found = false
				break
			}
		}

		// If n is evenly divisible by all of the numbers from 11 to 20, then we have found the answer
		if found {
			fmt.Println(n)
			break
		}

		// Otherwise, try the next number
		n++
	}
}
