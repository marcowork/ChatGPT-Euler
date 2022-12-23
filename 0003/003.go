package main

import "fmt"

func main() {
	// The number we want to find the largest prime factor of
	n := 600851475143

	// Start with the smallest prime number, 2
	largestFactor := 2

	// Keep dividing n by 2 until it is no longer evenly divisible
	for n%2 == 0 {
		n /= 2
	}

	// Now, start with the next prime number (3) and keep incrementing
	// by 2 (to skip even numbers) until we reach the square root of n
	for i := 3; i*i <= n; i += 2 {
		// If n is evenly divisible by i, divide it by i and update
		// the largest factor
		for n%i == 0 {
			largestFactor = i
			n /= i
		}
	}

	// If there is anything left of n, it must be a prime number
	// and therefore the largest prime factor
	if n > 2 {
		largestFactor = n
	}

	// Print the result
	fmt.Println(largestFactor)
}
